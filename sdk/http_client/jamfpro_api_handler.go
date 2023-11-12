// jamfpro_api_handler.go
/* ------------------------------Summary----------------------------------------
This is a api handler module for the http_client to accommodate specifics of
jamf's api(s). It handles the encoding (marshalling) and decoding (unmarshalling)
of data. It also sets the correct content headers for the various http methods.

This module integrates with the http_client logger for wrapped error handling
for human readable return codes. It also supports the http_client tiered logging
functionality for logging support.

The logic of this module is defined as follows:
Classic API:

For requests (GET, POST, PUT, DELETE):
- Encoding (Marshalling): Use XML format.
For responses (GET, POST, PUT):
- Decoding (Unmarshalling): Use XML format.
For responses (DELETE):
- Handle response codes as response body lacks anything useful.
Headers
- Sets accept headers based on weighting. XML out weighs JSON to ensure XML is returned
- Sets content header as application/xml with edge case exceptions based on need.

JamfPro API:

For requests (GET, POST, PUT, DELETE):
- Encoding (Marshalling): Use JSON format.
For responses (GET, POST, PUT):
- Decoding (Unmarshalling): Use JSON format.
For responses (DELETE):
- Handle response codes as response body lacks anything useful.
Headers
- Sets accept headers based on weighting. Jamf Pro API doesn't support XML, so MIME type is skipped and returns JSON
- Set content header as application/json with edge case exceptions based on need.
*/
package http_client

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	_ "embed"
)

// Endpoint constants represent the URL suffixes used for Jamf API token interactions.
const (
	BaseDomain              = ".jamfcloud.com"                // BaseDomain represents the base domain for the jamf instance.
	OAuthTokenEndpoint      = "/api/oauth/token"              // OAuthTokenEndpoint: The endpoint to obtain an OAuth token.
	BearerTokenEndpoint     = "/api/v1/auth/token"            // BearerTokenEndpoint: The endpoint to obtain a bearer token.
	TokenRefreshEndpoint    = "/api/v1/auth/keep-alive"       // TokenRefreshEndpoint: The endpoint to refresh an existing token.
	TokenInvalidateEndpoint = "/api/v1/auth/invalidate-token" // TokenInvalidateEndpoint: The endpoint to invalidate an active token.
)

// ConfigMap is a map that associates endpoint URL patterns with their corresponding configurations.
// The map's keys are strings that identify the endpoint, and the values are EndpointConfig structs
// that hold the configuration for that endpoint.
type ConfigMap map[string]EndpointConfig

// Variables
var configMap ConfigMap

// Embedded Resources
//
//go:embed jamfpro_api_exceptions_configuration.json
var defaultConfig []byte

// Package-level Functions

// init is invoked automatically on package initialization and is responsible for
// setting up the default state of the package by loading the default configuration.
// If an error occurs during the loading process, the program will terminate with a fatal error log.
func init() {
	// Load the default configuration from an embedded resource.
	err := loadDefaultConfig()
	if err != nil {
		log.Fatalf("Error loading default config: %s", err)
	}
}

// loadDefaultConfig reads and unmarshals the default configuration JSON data from an embedded file
// into the configMap variable, which holds the exceptions configuration for endpoint-specific headers.
// Returns an error if the unmarshalling process fails.
func loadDefaultConfig() error {
	// Unmarshal the embedded default configuration into the global configMap.
	return json.Unmarshal(defaultConfig, &configMap)
}

// LoadUserConfig allows users to apply their own configuration by providing a JSON file.
// The custom configuration will override the default settings previously loaded.
// It reads the file from the provided filename path and unmarshals its content into the configMap.
// If reading or unmarshalling fails, an error is returned.
func LoadUserConfig(filename string) error {
	// Read the user-provided JSON configuration file and unmarshal it into the global configMap.
	userConfigBytes, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	// Override the default configuration with the user's custom settings.
	return json.Unmarshal(userConfigBytes, &configMap)
}

// Structs

// EndpointConfig is a struct that holds configuration details for a specific API endpoint.
// It includes what type of content it can accept and what content type it should send.
type EndpointConfig struct {
	Accept      string  `json:"accept"`       // Accept specifies the MIME type the endpoint can handle in responses.
	ContentType *string `json:"content_type"` // ContentType, if not nil, specifies the MIME type to set for requests sent to the endpoint. A pointer is used to distinguish between a missing field and an empty string.
}

// UnifiedAPIHandler is a struct that implements the APIHandler interface.
// It holds a Logger instance to facilitate logging across various API handling methods.
// This handler is responsible for encoding and decoding request and response data,
// determining content types, and other API interactions as defined by the APIHandler interface.
type UnifiedJamfAPIHandler struct {
	logger Logger // logger is used to output logs for the API handling processes.
}

// Functions

// ConstructAPIResourceEndpoint returns the full URL for a Jamf API resource endpoint path.
func (c *Client) ConstructAPIResourceEndpoint(endpointPath string) string {
	url := fmt.Sprintf("https://%s%s%s", c.InstanceName, BaseDomain, endpointPath)
	c.logger.Info("Request will be made to API Resource URL:", "URL", url)
	return url
}

// ConstructAPIAuthEndpoint returns the full URL for a Jamf API auth endpoint path.
func (c *Client) ConstructAPIAuthEndpoint(endpointPath string) string {
	url := fmt.Sprintf("https://%s%s%s", c.InstanceName, BaseDomain, endpointPath)
	c.logger.Info("Request will be made to API authentication URL:", "URL", url)
	return url
}

// APIHandler is an interface for encoding, decoding, and determining content types for different API implementations.
// It encapsulates behavior for encoding and decoding requests and responses.
type APIHandler interface {
	MarshalRequest(body interface{}, method string, endpoint string) ([]byte, error)
	MarshalMultipartRequest(fields map[string]string, files map[string]string) ([]byte, string, error) // New method for multipart
	UnmarshalResponse(resp *http.Response, out interface{}) error
	GetContentTypeHeader(method string) string
	GetAcceptHeader() string
	SetLogger(logger Logger)
}

// GetAPIHandler initializes and returns an APIHandler with a configured logger.
func GetAPIHandler(config Config) APIHandler {
	handler := &UnifiedJamfAPIHandler{}
	logger := NewDefaultLogger()
	logger.SetLevel(config.LogLevel) // Use the LogLevel from the config
	handler.SetLogger(logger)
	return handler
}

// SetLogger assigns a Logger instance to the UnifiedAPIHandler.
// This allows for logging throughout the handler's operations,
// enabling consistent logging that follows the configuration of the provided Logger.
func (u *UnifiedJamfAPIHandler) SetLogger(logger Logger) {
	u.logger = logger
}

/*
// NewUnifiedAPIHandler creates a new UnifiedAPIHandler with the provided logger.
func NewUnifiedAPIHandler(logger Logger) *UnifiedAPIHandler {
	return &UnifiedAPIHandler{logger: logger}
}
*/

// GetContentTypeHeader determines the appropriate Content-Type header for a given API endpoint.
// It attempts to find a content type that matches the endpoint prefix in the global configMap.
// If a match is found and the content type is defined (not nil), it returns the specified content type.
// If the content type is nil or no match is found in configMap, it falls back to default behaviors:
// - For url endpoints starting with "/JSSResource", it defaults to "application/xml" for the Classic API.
// - For url endpoints starting with "/api", it defaults to "application/json" for the JamfPro API.
// If the endpoint does not match any of the predefined patterns, "application/json" is used as a fallback.
// This method logs the decision process at various stages for debugging purposes.
func (u *UnifiedJamfAPIHandler) GetContentTypeHeader(endpoint string) string {
	// Dynamic lookup from configuration should be the first priority
	for key, config := range configMap {
		if strings.HasPrefix(endpoint, key) {
			if config.ContentType != nil {
				u.logger.Debug("Content-Type for endpoint found in configMap", "endpoint", endpoint, "content_type", *config.ContentType)
				return *config.ContentType
			}
			u.logger.Debug("Content-Type for endpoint is nil in configMap, handling as special case", "endpoint", endpoint)
			// If a nil ContentType is an expected case, do not set Content-Type header.
			return "" // Return empty to indicate no Content-Type should be set.
		}
	}

	// If no specific configuration is found, then check for standard URL patterns.
	if strings.Contains(endpoint, "/JSSResource") {
		u.logger.Debug("Content-Type for endpoint defaulting to XML for Classic API", "endpoint", endpoint)
		return "application/xml" // Classic API uses XML
	} else if strings.Contains(endpoint, "/api") {
		u.logger.Debug("Content-Type for endpoint defaulting to JSON for JamfPro API", "endpoint", endpoint)
		return "application/json" // JamfPro API uses JSON
	}

	// Fallback to JSON if no other match is found.
	u.logger.Debug("Content-Type for endpoint not found in configMap or standard patterns, using default JSON", "endpoint", endpoint)
	return "application/json"
}

// MarshalRequest encodes the request body according to the endpoint for the API.
func (u *UnifiedJamfAPIHandler) MarshalRequest(body interface{}, method string, endpoint string) ([]byte, error) {
	var (
		data []byte
		err  error
	)

	// Determine the format based on the endpoint
	format := "json"
	if strings.Contains(endpoint, "/JSSResource") {
		format = "xml"
	} else if strings.Contains(endpoint, "/api") {
		format = "json"
	}

	switch format {
	case "xml":
		data, err = xml.Marshal(body)
		if err != nil {
			return nil, err
		}

		if method == "POST" || method == "PUT" {
			u.logger.Trace("XML Request Body:", "Body", string(data))
		}

	case "json":
		data, err = json.Marshal(body)
		if err != nil {
			u.logger.Error("Failed marshaling JSON request", "error", err)
			return nil, err
		}

		if method == "POST" || method == "PUT" || method == "PATCH" {
			u.logger.Debug("JSON Request Body:", string(data))
		}
	}

	return data, nil
}

// UnmarshalResponse decodes the response body from XML or JSON format depending on the Content-Type header.
func (u *UnifiedJamfAPIHandler) UnmarshalResponse(resp *http.Response, out interface{}) error {
	// Handle DELETE method
	if resp.Request.Method == "DELETE" {
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		} else {
			return fmt.Errorf("DELETE request failed with status code: %d", resp.StatusCode)
		}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		u.logger.Error("Failed reading response body", "error", err)
		return err
	}

	// Log the raw response body and headers
	u.logger.Trace("Raw HTTP Response:", string(bodyBytes))
	u.logger.Debug("Unmarshaling response", "status", resp.Status)

	// Log headers when in debug mode
	u.logger.Debug("HTTP Response Headers:", resp.Header)

	// Check the Content-Type and Content-Disposition headers
	contentType := resp.Header.Get("Content-Type")
	contentDisposition := resp.Header.Get("Content-Disposition")

	// Handle binary data if necessary
	if err := u.handleBinaryData(contentType, contentDisposition, bodyBytes, out); err != nil {
		return err
	}

	// If content type is HTML, extract the error message
	if strings.Contains(contentType, "text/html") {
		errMsg := extractErrorMessageFromHTML(string(bodyBytes))
		u.logger.Warn("Received HTML content", "error_message", errMsg, "status_code", resp.StatusCode)
		return &APIError{
			StatusCode: resp.StatusCode,
			Message:    errMsg,
		}
	}

	// Check for non-success status codes before attempting to unmarshal
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		// Parse the error details from the response body for JSON content type
		if strings.Contains(contentType, "application/json") {
			description, err := parseJSONErrorResponse(bodyBytes)
			if err != nil {
				u.logger.Error("Failed to parse JSON error response", "error", err)
				return fmt.Errorf("received non-success status code: %d and failed to parse error response", resp.StatusCode)
			}
			return fmt.Errorf("received non-success status code: %d, error: %s", resp.StatusCode, description)
		}

		// If the response is not JSON or another error occurs, return a generic error message
		u.logger.Error("Received non-success status code", "status_code", resp.StatusCode)
		return fmt.Errorf("received non-success status code: %d", resp.StatusCode)
	}

	// Determine whether the content type is JSON or XML and unmarshal accordingly
	switch {
	case strings.Contains(contentType, "application/json"):
		err = json.Unmarshal(bodyBytes, out)
	case strings.Contains(contentType, "application/xml"), strings.Contains(contentType, "text/xml;charset=UTF-8"):
		err = xml.Unmarshal(bodyBytes, out)
	default:
		// If the content type is neither JSON nor XML nor HTML
		return fmt.Errorf("unexpected content type: %s", contentType)
	}

	// Handle any errors that occurred during unmarshaling
	if err != nil {
		// If unmarshalling fails, check if the content might be HTML
		if strings.Contains(string(bodyBytes), "<html>") {
			errMsg := extractErrorMessageFromHTML(string(bodyBytes))
			u.logger.Warn("Received HTML content instead of expected format", "error_message", errMsg, "status_code", resp.StatusCode)
			return fmt.Errorf(errMsg)
		}

		// Log the error and return it
		u.logger.Error("Failed to unmarshal response", "error", err)
		return fmt.Errorf("failed to unmarshal response: %v", err)
	}

	return nil
}

// GetAcceptHeader constructs and returns a weighted Accept header string for HTTP requests.
// The Accept header indicates the MIME types that the client can process and prioritizes them
// based on the quality factor (q) parameter. Higher q-values signal greater preference.
// This function specifies a range of MIME types with their respective weights, ensuring that
// the server is informed of the client's versatile content handling capabilities while
// indicating a preference for XML. The specified MIME types cover common content formats like
// images, JSON, XML, HTML, plain text, and certificates, with a fallback option for all other types.
func (u *UnifiedJamfAPIHandler) GetAcceptHeader() string {
	weightedAcceptHeader := "application/x-x509-ca-cert;q=0.95," +
		"application/pkix-cert;q=0.94," +
		"application/pem-certificate-chain;q=0.93," +
		"application/octet-stream;q=0.8," + // For general binary files
		"image/png;q=0.75," +
		"image/jpeg;q=0.74," +
		"image/*;q=0.7," +
		"application/xml;q=0.65," +
		"text/xml;q=0.64," +
		"text/xml;charset=UTF-8;q=0.63," +
		"application/json;q=0.5," +
		"text/html;q=0.5," +
		"text/plain;q=0.4," +
		"*/*;q=0.05" // Fallback for any other types
	return weightedAcceptHeader
}

// MarshalMultipartFormData takes a map with form fields and file paths and returns the encoded body and content type.
func (u *UnifiedJamfAPIHandler) MarshalMultipartRequest(fields map[string]string, files map[string]string) ([]byte, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the simple fields to the form data
	for field, value := range fields {
		if err := writer.WriteField(field, value); err != nil {
			return nil, "", err
		}
	}

	// Add the files to the form data
	for formField, filepath := range files {
		file, err := os.Open(filepath)
		if err != nil {
			return nil, "", err
		}
		defer file.Close()

		part, err := writer.CreateFormFile(formField, filepath)
		if err != nil {
			return nil, "", err
		}
		if _, err := io.Copy(part, file); err != nil {
			return nil, "", err
		}
	}

	// Close the writer before returning
	contentType := writer.FormDataContentType()
	if err := writer.Close(); err != nil {
		return nil, "", err
	}

	return body.Bytes(), contentType, nil
}

// handleBinaryData checks if the response should be treated as binary data and assigns to out if so.
func (u *UnifiedJamfAPIHandler) handleBinaryData(contentType, contentDisposition string, bodyBytes []byte, out interface{}) error {
	if strings.Contains(contentType, "application/octet-stream") || strings.HasPrefix(contentDisposition, "attachment") {
		if outPointer, ok := out.(*[]byte); ok {
			*outPointer = bodyBytes
			return nil
		} else {
			return fmt.Errorf("output parameter is not a *[]byte for binary data")
		}
	}
	return nil // If not binary data, no action needed
}
