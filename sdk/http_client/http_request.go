// http_request.go
package http_client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

// DoRequest constructs and executes a standard HTTP request with support for retry logic.
// It is intended for operations that can be encoded in a single JSON or XML body such as
// creating or updating resources. This method includes token validation, concurrency control,
// performance metrics, dynamic header setting, and structured error handling.
//
// Parameters:
// - method: The HTTP method to use (e.g., GET, POST, PUT, DELETE, PATCH).
// - endpoint: The API endpoint to which the request will be sent.
// - body: The payload to send in the request, which will be marshaled based on the API handler rules.
// - out: A pointer to a variable where the unmarshaled response will be stored.
//
// Returns:
// - A pointer to the http.Response received from the server.
// - An error if the request could not be sent, the response could not be processed, or if retry attempts fail.
//
// The function starts by validating the client's authentication token and managing concurrency using
// a token system. It then determines the appropriate API handler for marshaling the request body and
// setting headers. The request is sent to the constructed URL with all necessary headers including
// authorization, content type, and user agent.
//
// If configured for debug logging, the function logs all request headers before sending. The function then
// enters a loop to handle retryable HTTP methods, implementing a retry mechanism for transient errors,
// rate limits, and other retryable conditions based on response status codes.
//
// The function also updates performance metrics to track total request count and cumulative response time.
// After processing the response, it handles any API errors and unmarshals the response body into the provided
// 'out' parameter if the response is successful.
//
// Note:
// The function assumes that retryable HTTP methods have been properly defined in the retryableHTTPMethods map.
// It is the caller's responsibility to close the response body when the request is successful to avoid resource leaks.
func (c *Client) DoRequest(method, endpoint string, body, out interface{}) (*http.Response, error) {
	// Auth Token validation check
	valid, err := c.ValidAuthTokenCheck()
	if err != nil || !valid {
		return nil, fmt.Errorf("validity of the authentication token failed with error: %w", err)
	}

	// Acquire a token for concurrency management with a timeout and measure its acquisition time
	tokenAcquisitionStart := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	requestID, err := c.ConcurrencyMgr.Acquire(ctx)
	if err != nil {
		return nil, err
	}
	defer c.ConcurrencyMgr.Release(requestID)

	tokenAcquisitionDuration := time.Since(tokenAcquisitionStart)
	c.PerfMetrics.lock.Lock()
	c.PerfMetrics.TokenWaitTime += tokenAcquisitionDuration
	c.PerfMetrics.lock.Unlock()

	// Add the request ID to the context
	ctx = context.WithValue(ctx, requestIDKey{}, requestID)

	// Determine which set of encoding and content-type request rules to use
	//handler := GetAPIHandler(endpoint, c.config.LogLevel)
	handler := GetAPIHandler(c.config)

	// Marshal Request with correct encoding
	requestData, err := handler.MarshalRequest(body, method, endpoint)
	if err != nil {
		return nil, err
	}

	// Construct URL using the ConstructAPIResourceEndpoint function
	url := c.ConstructAPIResourceEndpoint(endpoint)

	// Initialize total request counter
	c.PerfMetrics.lock.Lock()
	c.PerfMetrics.TotalRequests++
	c.PerfMetrics.lock.Unlock()

	// Perform Request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestData))
	if err != nil {
		return nil, err
	}

	// Define header content type based on url and http method
	contentType := handler.GetContentTypeHeader(endpoint)
	// Define Request Headers dynamically based on handler logic
	acceptHeader := handler.GetAcceptHeader()

	// Set Headers
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accept", acceptHeader)
	req.Header.Set("User-Agent", GetUserAgentHeader())

	// Debug: Print request headers if in debug mode
	c.logger.Debug("HTTP Request Headers:", req.Header)

	// Define if request is retryable
	retryableHTTPMethods := map[string]bool{
		http.MethodGet:    true, // GET
		http.MethodDelete: true, // DELETE
		http.MethodPut:    true, // PUT
		http.MethodPatch:  true, // PATCH
	}

	if retryableHTTPMethods[method] {
		// Define a deadline for total retries based on http client TotalRetryDuration config
		totalRetryDeadline := time.Now().Add(c.config.TotalRetryDuration)
		i := 0
		for {
			// Check if we've reached the maximum number of retries or if our total retry time has exceeded
			if i > c.config.MaxRetryAttempts || time.Now().After(totalRetryDeadline) {
				return nil, fmt.Errorf("max retry attempts reached or total retry duration exceeded")
			}

			// This context is used to propagate cancellations and timeouts for the request.
			// For example, if a request's context gets canceled or times out, the request will be terminated early.
			req = req.WithContext(ctx)

			// Start response time measurement
			responseTimeStart := time.Now()

			// Execute Request with context
			resp, err := c.httpClient.Do(req)
			if err != nil {
				c.logger.Error("Failed to send request", "method", method, "endpoint", endpoint, "error", err)
				return nil, err
			}

			// After each request, compute and update response time
			responseDuration := time.Since(responseTimeStart)
			c.PerfMetrics.lock.Lock()
			c.PerfMetrics.TotalResponseTime += responseDuration
			c.PerfMetrics.lock.Unlock()

			// Checks for the presence of a deprecation header in the HTTP response and logs if found.
			if i == 0 {
				CheckDeprecationHeader(resp, c.logger)
			}

			// Handle (unmarshall) response with API Handler
			if err := handler.UnmarshalResponse(resp, out); err != nil {
				switch e := err.(type) {
				case *APIError:
					c.logger.Error("Received an API error", "status_code", e.StatusCode, "message", e.Message)
					return resp, e
				default:
					// Existing error handling logic
					c.logger.Error("Failed to unmarshal HTTP response", "method", method, "endpoint", endpoint, "error", err)
					return resp, err
				}
			}

			if resp.StatusCode >= 200 && resp.StatusCode < 300 {
				c.logger.Info("HTTP request succeeded", "method", method, "endpoint", endpoint, "status_code", resp.StatusCode)
				return resp, nil
			} else if resp.StatusCode == http.StatusNotFound {
				c.logger.Warn("Resource not found", "method", method, "endpoint", endpoint)
				return resp, fmt.Errorf("resource not found: %s", endpoint)
			}

			// Retry Logic
			if isNonRetryableError(resp) {
				c.logger.Warn("Encountered a non-retryable error", "status", resp.StatusCode, "description", translateStatusCode(resp.StatusCode))
				return resp, c.handleAPIError(resp)
			} else if isRateLimitError(resp) {
				waitDuration := parseRateLimitHeaders(resp) // Checks for the Retry-After, X-RateLimit-Remaining and X-RateLimit-Reset headers
				c.logger.Warn("Encountered a rate limit error. Retrying after wait duration.", "wait_duration", waitDuration)
				time.Sleep(waitDuration)
				i++
				continue // This will restart the loop, effectively "retrying" the request
			} else if isTransientError(resp) {
				waitDuration := calculateBackoff(i) //uses exponential backoff (with jitter)
				c.logger.Warn("Encountered a transient error. Retrying after backoff.", "wait_duration", waitDuration)
				time.Sleep(waitDuration)
				i++
				continue // This will restart the loop, effectively "retrying" the request
			} else {
				c.logger.Error("Received unexpected error status from HTTP request", "method", method, "endpoint", endpoint, "status_code", resp.StatusCode, "description", translateStatusCode(resp.StatusCode))
				return resp, c.handleAPIError(resp)
			}
		}
	} else {
		// Start response time measurement
		responseTimeStart := time.Now()
		// For non-retryable HTTP Methods (POST - Create)
		req = req.WithContext(ctx)
		resp, err := c.httpClient.Do(req)

		if err != nil {
			c.logger.Error("Failed to send request", "method", method, "endpoint", endpoint, "error", err)
			return nil, err
		}

		// After the request, compute and update response time
		responseDuration := time.Since(responseTimeStart)
		c.PerfMetrics.lock.Lock()
		c.PerfMetrics.TotalResponseTime += responseDuration
		c.PerfMetrics.lock.Unlock()

		CheckDeprecationHeader(resp, c.logger)

		// Unmarshal the response with the determined API Handler
		if err := handler.UnmarshalResponse(resp, out); err != nil {
			switch e := err.(type) {
			case *APIError:
				c.logger.Error("Received an API error", "status_code", e.StatusCode, "message", e.Message)
				return resp, e
			default:
				// Existing error handling logic
				c.logger.Error("Failed to unmarshal HTTP response", "method", method, "endpoint", endpoint, "error", err)
				return resp, err
			}
		}

		// Check if the response status code is within the success range
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return resp, nil
		} else {
			statusDescription := translateStatusCode(resp.StatusCode)
			c.logger.Error("Received non-success status code from HTTP request", "method", method, "endpoint", endpoint, "status_code", resp.StatusCode, "description", statusDescription)
			return resp, fmt.Errorf("Error status code: %d - %s", resp.StatusCode, statusDescription)
		}
	}
	// TODO refactor to remove repition.
}

// DoMultipartRequest creates and executes a multipart HTTP request. It is used for sending files
// and form fields in a single request. This method handles the construction of the multipart
// message body, setting the appropriate headers, and sending the request to the given endpoint.
//
// Parameters:
// - method: The HTTP method to use (e.g., POST, PUT).
// - endpoint: The API endpoint to which the request will be sent.
// - fields: A map of form fields and their values to include in the multipart message.
// - files: A map of file field names to file paths that will be included as file attachments.
// - out: A pointer to a variable where the unmarshaled response will be stored.
//
// Returns:
// - A pointer to the http.Response received from the server.
// - An error if the request could not be sent or the response could not be processed.
//
// The function first validates the authentication token, then constructs the multipart
// request body based on the provided fields and files. It then constructs the full URL for
// the request, sets the required headers (including Authorization and Content-Type), and
// sends the request.
//
// If debug mode is enabled, the function logs all the request headers before sending the request.
// After the request is sent, the function checks the response status code. If the response is
// not within the success range (200-299), it logs an error and returns the response and an error.
// If the response is successful, it attempts to unmarshal the response body into the 'out' parameter.
//
// Note:
// The caller should handle closing the response body when successful.
func (c *Client) DoMultipartRequest(method, endpoint string, fields map[string]string, files map[string]string, out interface{}) (*http.Response, error) {
	// Auth Token validation check
	valid, err := c.ValidAuthTokenCheck()
	if err != nil || !valid {
		return nil, fmt.Errorf("validity of the authentication token failed with error: %w", err)
	}

	// Determine which set of encoding and content-type request rules to use
	//handler := GetAPIHandler(endpoint, c.config.LogLevel)
	handler := GetAPIHandler(c.config)

	// Marshal the multipart form data
	requestData, contentType, err := handler.MarshalMultipartRequest(fields, files)
	if err != nil {
		return nil, err
	}

	// Construct URL using the ConstructAPIResourceEndpoint function
	url := c.ConstructAPIResourceEndpoint(endpoint)

	// Create the request
	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestData))
	if err != nil {
		return nil, err
	}

	// Set Request Headers
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("User-Agent", GetUserAgentHeader())

	// Debug: Print request headers if in debug mode

	c.logger.Debug("HTTP Multipart Request Headers:", req.Header)

	// Execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.logger.Error("Failed to send multipart request", "method", method, "endpoint", endpoint, "error", err)
		return nil, err
	}

	// Check for successful status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		c.logger.Error("Received non-success status code from multipart request", "status_code", resp.StatusCode)
		return resp, fmt.Errorf("received non-success status code: %d", resp.StatusCode)
	}

	// Unmarshal the response
	if err := handler.UnmarshalResponse(resp, out); err != nil {
		c.logger.Error("Failed to unmarshal HTTP response", "method", method, "endpoint", endpoint, "error", err)
		return resp, err
	}

	return resp, nil
}
