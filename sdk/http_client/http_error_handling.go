// http_error_handling.go
// This package provides utility functions and structures for handling and categorizing HTTP error responses.
package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// APIError represents a structured API error response.
type APIError struct {
	StatusCode int
	Message    string
}

// handleAPIError handles error responses from the API, converting them into a structured error if possible.
func (c *Client) handleAPIError(resp *http.Response) error {
	var structuredErr StructuredError
	err := json.NewDecoder(resp.Body).Decode(&structuredErr)
	if err == nil && structuredErr.Error.Message != "" {
		c.logger.Warn("API returned structured error", "status", resp.Status, "error_code", structuredErr.Error.Code, "error_message", structuredErr.Error.Message)
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    structuredErr.Error.Message,
		}
	}

	var errMsg string
	err = json.NewDecoder(resp.Body).Decode(&errMsg)
	if err != nil || errMsg == "" {
		errMsg = fmt.Sprintf("Unexpected error with status code: %d", resp.StatusCode)
		c.logger.Warn("Failed to decode API error message, using default error message", "status", resp.Status)
	} else {
		c.logger.Warn("API returned non-structured error", "status", resp.Status, "error_message", errMsg)
	}

	return &APIError{
		StatusCode: resp.StatusCode,
		Message:    errMsg,
	}
}

// Error returns a string representation of the APIError.
func (e *APIError) Error() string {
	return fmt.Sprintf("API Error (Code: %d): %s", e.StatusCode, e.Message)
}

// translateStatusCode provides a human-readable message for HTTP status codes.
func translateStatusCode(statusCode int) string {
	messages := map[int]string{
		http.StatusOK:                    "Request successful.",
		http.StatusCreated:               "Request to create or update resource successful.",
		http.StatusAccepted:              "The request was accepted for processing, but the processing has not completed.",
		http.StatusNoContent:             "Request successful. Resource successfully deleted.",
		http.StatusBadRequest:            "Bad request. Verify the syntax of the request.",
		http.StatusUnauthorized:          "Authentication failed. Verify the credentials being used for the request.",
		http.StatusForbidden:             "Invalid permissions. Verify the account being used has the proper permissions for the resource you are trying to access.",
		http.StatusNotFound:              "Resource not found. Verify the URL path is correct.",
		http.StatusConflict:              "Conflict. See the error response for additional details.",
		http.StatusPreconditionFailed:    "Precondition failed. See error description for additional details.",
		http.StatusRequestEntityTooLarge: "Payload too large.",
		http.StatusRequestURITooLong:     "Request-URI too long.",
		http.StatusInternalServerError:   "Internal server error. Retry the request or contact support if the error persists.",
		http.StatusBadGateway:            "Bad Gateway. Generally due to a timeout issue.",
		http.StatusServiceUnavailable:    "Service unavailable.",
	}

	if message, exists := messages[statusCode]; exists {
		return message
	}
	return "An unexpected error occurred. Please try again later."
}

// isNonRetryableError checks if the provided response indicates a non-retryable error.
func isNonRetryableError(resp *http.Response) bool {
	// List of non-retryable HTTP status codes
	nonRetryableStatusCodes := map[int]bool{
		http.StatusBadRequest:            true, // 400
		http.StatusUnauthorized:          true, // 401
		http.StatusForbidden:             true, // 403
		http.StatusNotFound:              true, // 404
		http.StatusConflict:              true, // 409
		http.StatusRequestEntityTooLarge: true, // 413
		http.StatusRequestURITooLong:     true, // 414
	}

	_, isNonRetryable := nonRetryableStatusCodes[resp.StatusCode]
	return isNonRetryable
}

// isRateLimitError checks if the provided response indicates a rate limit error.
func isRateLimitError(resp *http.Response) bool {
	return resp.StatusCode == http.StatusTooManyRequests
}

// isTransientError checks if an error or HTTP response indicates a transient error.
func isTransientError(resp *http.Response) bool {
	transientStatusCodes := map[int]bool{
		http.StatusInternalServerError: true,
		http.StatusBadGateway:          true,
		http.StatusServiceUnavailable:  true,
	}
	return resp != nil && transientStatusCodes[resp.StatusCode]
}

// extractErrorMessageFromHTML attempts to parse an HTML error page and extract a human-readable error message.
func extractErrorMessageFromHTML(htmlContent string) string {
	r := bytes.NewReader([]byte(htmlContent))
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "Unable to parse HTML content"
	}

	var messages []string
	doc.Find("p").Each(func(i int, s *goquery.Selection) {
		messages = append(messages, s.Text())
	})

	return strings.Join(messages, " | ")
}

// parseJSONErrorResponse parses the JSON error message from the response body.
func parseJSONErrorResponse(body []byte) (string, error) {
	var errorResponse struct {
		HTTPStatus int `json:"httpStatus"`
		Errors     []struct {
			Code        string `json:"code"`
			Description string `json:"description"`
			ID          string `json:"id"`
			Field       string `json:"field"`
		} `json:"errors"`
	}

	err := json.Unmarshal(body, &errorResponse)
	if err != nil {
		return "", err
	}

	if len(errorResponse.Errors) > 0 {
		return errorResponse.Errors[0].Description, nil
	}

	return "No error description available", nil
}
