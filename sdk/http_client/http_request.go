// http_request.go
package http_client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"time"
)

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
	handler := GetAPIHandler(endpoint, c.config.DebugMode)

	// Construct request
	requestData, err := handler.MarshalRequest(body, method)
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
	contentType := handler.GetContentType(method)
	// Define Request Headers dynamically based on handler logic
	acceptHeader := handler.GetAcceptHeader(url)

	// Set Headers
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Accept", acceptHeader)
	req.Header.Set("User-Agent", GetUserAgent())

	// Debug: Print request headers if in debug mode
	if c.config.DebugMode {
		c.logger.Debug("HTTP Request Headers:", req.Header)
	}

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

func (c *Client) DoMultipartRequest(method, endpoint string, fields map[string]string, files map[string]string, out interface{}) (*http.Response, error) {
	// Auth Token validation check
	valid, err := c.ValidAuthTokenCheck()
	if err != nil || !valid {
		return nil, fmt.Errorf("validity of the authentication token failed with error: %w", err)
	}

	// Determine which set of encoding and content-type request rules to use
	handler := GetAPIHandler(endpoint, c.config.DebugMode)

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
	req.Header.Set("User-Agent", GetUserAgent())

	// Debug: Print request headers if in debug mode
	if c.config.DebugMode {
		c.logger.Debug("HTTP Multipart Request Headers:", req.Header)
	}

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
