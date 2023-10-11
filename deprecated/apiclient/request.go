// request.go
// Provides functions to interact with the Jamf API.
package apiClient

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/cenkalti/backoff"
)

type StatusCodeCategory int

const (
	Success StatusCodeCategory = iota
	ClientError
	ServerError
	Unknown
	maxErrorMessageLength = 500 // Maximum characters for the error message
	ClassicAPI            = "/JSSResource"
	JamfProAPI            = "/api"
)

type Error interface {
	Error() string
	StatusCode() int
	URI() string
	Body() string
}
type errorInfo struct {
	statusCode int
	uri        string
	body       string
}

func newError(statusCode int, uri, body string) Error {
	var e = new(errorInfo)
	e.statusCode = statusCode
	e.uri = uri
	e.body = body
	return e
}

func (e *errorInfo) Error() string {
	return fmt.Sprintf("Encountered a %v error when accessing the API at %s. Response body: %s", e.statusCode, e.uri, e.body)
}

func (e *errorInfo) StatusCode() int {
	return e.statusCode
}

func (e *errorInfo) URI() string {
	return e.uri
}

func (e *errorInfo) Body() string {
	return e.body
}

func categorizeStatusCode(statusCode int) StatusCodeCategory {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return Success
	case statusCode >= 400 && statusCode < 500:
		return ClientError
	case statusCode >= 500 && statusCode < 600:
		return ServerError
	default:
		return Unknown
	}
}

// uriForApi ... Generate uri for api
func (c *Client) uriForAPI(api string) string {
	return fmt.Sprintf("https://%s%s", c.url, api)
}

// sanitizeErrorBody removes newline characters from the body for cleaner error messages
func sanitizeErrorBody(body []byte) string {
	// Remove unwanted characters
	sanitized := strings.ReplaceAll(string(body), "\n", " ")
	sanitized = strings.ReplaceAll(sanitized, "\r", " ")
	sanitized = strings.ReplaceAll(sanitized, "\t", " ")

	// Truncate overly long error messages
	if len(sanitized) > maxErrorMessageLength {
		sanitized = sanitized[:maxErrorMessageLength] + "..."
	}

	// Provide a concise and actionable default message if sanitized message is empty
	sanitized = strings.TrimSpace(sanitized)
	if sanitized == "" {
		sanitized = "The server returned an error without a specific message."
	}

	return sanitized
}

// DoRequest - A method to send a request to the jamf api
func (c *Client) DoRequest(method, api string, reqbody interface{}, params *url.Values, out interface{}) error {
	req, err := c.createRequest(method, api, params, reqbody)
	if err != nil {
		return err
	}

	// Direct methods without retries
	directMethods := map[string]bool{
		http.MethodPost: true,
		http.MethodPut:  true,
	}

	var resp *http.Response
	if directMethods[method] {
		resp, err = c.HttpClient.Do(req)
	} else {
		resp, err = c.doRequestWithRetries(req)
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	statusCategory := categorizeStatusCode(resp.StatusCode)
	if statusCategory != Success {
		return newError(resp.StatusCode, api, sanitizeErrorBody(body))
	}

	// Return early if no place to put the response
	if out == nil {
		return nil
	}

	// If we got no body, by default let's just make an empty JSON dict.
	if len(body) == 0 {
		body = []byte{'{', '}'}
	}

	if strings.Contains(api, "JSSResource") {
		return xml.Unmarshal(body, out)
	}
	return json.Unmarshal(body, out)
}

// doRequestWithRetries - Use backoff to extend the wait interval for retrying exponentially.
func (c *Client) doRequestWithRetries(req *http.Request) (*http.Response, error) {
	backOff := backoff.NewExponentialBackOff()
	backOff.MaxElapsedTime = c.HttpRetryTimeout
	backOff.MaxInterval = 30 * time.Second    // Max wait interval between retries. Example: 30 seconds
	backOff.InitialInterval = 1 * time.Second // Initial wait interval. Example: 1 second

	var resp *http.Response
	var bodyCache []byte
	var err error

	if req.Body != nil {
		bodyCache, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("failed to read request body: %w", err)
		}
	}

	backOffReq := func() error {
		if bodyCache != nil {
			req.Body = io.NopCloser(bytes.NewReader(bodyCache))
		}

		resp, err = c.HttpClient.Do(req)
		if err != nil {
			return backoff.Permanent(err) // Permanent error, no retries
		}

		// Handling rate limit
		if resp.StatusCode == 429 { // 429 Too Many Requests
			if resetHeader := resp.Header.Get("X-RateLimit-Reset"); resetHeader != "" {
				resetTime, err := strconv.ParseInt(resetHeader, 10, 64)
				if err == nil {
					waitDuration := time.Until(time.Unix(resetTime, 0))
					time.Sleep(waitDuration)
				}
			}
			return fmt.Errorf("rate limit exceeded, retrying after a pause")
		}

		switch categorizeStatusCode(resp.StatusCode) {
		case Success:
			return nil
		case ClientError:
			return backoff.Permanent(fmt.Errorf("client-side error with status code %d. Request will not be retried", resp.StatusCode)) // Permanent error, no retries
		case ServerError:
			// For server errors, we will allow retries. So, we won't return a permanent error.
			return fmt.Errorf("server-side error with status code %d. retrying the request", resp.StatusCode)
		default:
			return fmt.Errorf("received an unexpected HTTP status code %d. Please check the API documentation or contact support", resp.StatusCode)
		}
	}

	err = backoff.Retry(backOffReq, backOff)
	if err != nil {
		return nil, fmt.Errorf("request failed after retries: %w", err)
	}

	return resp, nil
}

// Orchestrates the entire request-response cycle. Generates a http request for api.
// Supports Oauth, bearer token and username and password.
func (c *Client) createRequest(method, api string, params *url.Values, reqbody interface{}) (*http.Request, error) {
	var bodyReader io.Reader

	// Convert the request body to the appropriate type
	bodyReader, err := c.setRequestBody(method, api, reqbody)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.uriForAPI(api), bodyReader)
	if err != nil {
		return req, err
	}

	if params != nil {
		req.URL.RawQuery = params.Encode()
	}

	err = c.ensureValidToken(req)
	if err != nil {
		return nil, err
	}

	c.setRequestHeaders(req, api, method)

	return req, nil
}

// setRequestBody prepares the body of the HTTP request based on the given method, API endpoint, and request body data.
// The function returns an io.Reader that represents the prepared request body, and an error if any issues arise during the process.
func (c *Client) setRequestBody(method, api string, reqbody interface{}) (io.Reader, error) {
	if method != http.MethodGet && reqbody != nil {
		switch {
		case api == uriOauthAuthToken:
			formData, ok := reqbody.(map[string]string)
			if !ok {
				return nil, fmt.Errorf("expected map[string]string for request body, got %T", reqbody)
			}
			values := url.Values{}
			for key, value := range formData {
				values.Add(key, value)
			}
			return strings.NewReader(values.Encode()), nil
		case strings.Contains(api, "JSSResource"):
			b, err := xml.Marshal(reqbody)
			if err != nil {
				return nil, err
			}
			return bytes.NewReader(b), nil
		default:
			b, err := json.Marshal(reqbody)
			if err != nil {
				return nil, err
			}
			return bytes.NewReader(b), nil
		}
	}
	return nil, nil
}

// ensureValidToken checks and ensures that the client has a valid authentication token.
// If the client has a token and it has expired, it attempts to refresh it.
func (c *Client) ensureValidToken(req *http.Request) error {
	// First, handle token expiration checks
	if c.token != nil && c.tokenExpiration != nil && c.tokenExpiration.Before(time.Now()) {
		if c.isUsingOAuth {
			err := c.refreshOAuthToken()
			if err != nil {
				return err
			}
		} else {
			err := c.refreshAuthToken()
			if err != nil {
				return err
			}
		}
	}

	// Next, set the appropriate Authorization headers
	if c.token != nil {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", *c.token))
	} else if c.username != "" && c.password != "" {
		req.SetBasicAuth(c.username, c.password)
	}

	return nil
}

// setRequestHeaders sets the necessary headers for the given API request.
// Depending on the API endpoint, it determines the appropriate content type and other headers.
func (c *Client) setRequestHeaders(req *http.Request, api string, method string) {
	// Handle the OAuth token special case first
	if api == uriOauthAuthToken {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		return
	}

	// Handle the Classic API
	if strings.HasPrefix(api, ClassicAPI) {
		if method == http.MethodPost || method == http.MethodPut {
			req.Header.Add("Content-Type", "application/xml")
		} else {
			// For GET, this could be configurable based on client's choice
			req.Header.Add("Accept", "application/xml") // or application/json
		}
	} else if strings.HasPrefix(api, JamfProAPI) {
		req.Header.Add("Content-Type", "application/json")
		// Special cases for endpoints supporting other data types should be handled here
	}

	// Add any extra headers
	for k, v := range c.ExtraHeader {
		if req.Header.Get(k) == "" {
			req.Header.Add(k, v)
		}
	}
}

// DoRequestDebug provides complete debugging information during the request-response cycle.
// It uses createRequestDebug to create and log the request, then handles the HTTP call
// and response logging any necessary debugging information.
func (c *Client) DoRequestDebug(method, api string, reqbody interface{}, params *url.Values, out interface{}) error {
	// Create and log the request using createRequestDebug
	req, err := c.createRequestDebug(method, api, params, reqbody)
	if err != nil {
		return fmt.Errorf("failed to create debug request: %w", err)
	}

	// Direct methods without retries
	directMethods := map[string]bool{
		http.MethodPost: true,
		http.MethodPut:  true,
	}

	var resp *http.Response
	if directMethods[method] {
		resp, err = c.HttpClient.Do(req)
		if err != nil {
			log.Println("Failed to execute the HTTP request due to:", err.Error()) // Log the error
			return fmt.Errorf("failed to execute HTTP request: %w", err)
		}
	} else {
		resp, err = c.doRequestWithRetries(req)
		if err != nil {
			log.Println("Failed to execute the request even after several retries due to:", err.Error()) // Log the error
			return fmt.Errorf("failed to execute request with retries: %w", err)
		}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to read the response body from the server due to:", err.Error()) // Log the error
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// Always log the raw response for debugging purposes
	log.Println("Raw Response:", string(body))

	statusCategory := categorizeStatusCode(resp.StatusCode)
	switch statusCategory {
	case Success:
		// Do nothing as the response is successful
	case ClientError, ServerError, Unknown:
		return newError(resp.StatusCode, api, sanitizeErrorBody(body))
	}

	// Return early if no place to put the response
	if out == nil {
		return nil
	}

	// Response handling logic reused from DoRequest
	// If we got no body, by default let's just make an empty JSON dict.
	if len(body) == 0 {
		body = []byte{'{', '}'}
	}

	if strings.Contains(api, "JSSResource") {
		return xml.Unmarshal(body, out)
	}
	return json.Unmarshal(body, out)
}

// createRequestDebug is a wrapper around the standard createRequest function. It creates the request,
// logs key details like method, endpoint, headers, and body, which can be invaluable when trying to
// figure out issues with the request creation process.
func (c *Client) createRequestDebug(method, api string, params *url.Values, reqbody interface{}) (*http.Request, error) {
	// Use the standard createRequest for main logic
	req, err := c.createRequest(method, api, params, reqbody)
	if err != nil {
		// If there's an error while creating the request, log it
		log.Println("Failed to construct the API request due to:", err.Error())
		return nil, err
	}

	// Debugging logs
	log.Println("Debug Request Info:")
	log.Println("Method:", method)
	log.Println("API Endpoint:", api)

	// Print request headers for debugging
	PrintRequestHeaders(req)

	// Print request body for debugging
	if req.Body != nil {
		bodyBytes, _ := io.ReadAll(req.Body)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		log.Println("Request Body:", string(bodyBytes))
	}

	return req, nil
}

// DoRawRequest sends a GET request to the specified API endpoint and returns the raw string response.
// This function is specialized for handling raw string responses like extracting certificate public keys.
func (c *Client) DoRawRequest(api string, params *url.Values) (string, error) {
	req, err := c.createRequest("GET", api, params, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.doRequestWithRetries(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	statusCategory := categorizeStatusCode(resp.StatusCode)
	switch statusCategory {
	case Success:
		// Do nothing as the response is successful
	case ClientError, ServerError, Unknown:
		return "", newError(resp.StatusCode, api, sanitizeErrorBody(body))
	}

	return string(body), nil
}
