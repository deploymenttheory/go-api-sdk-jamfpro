package jamfpro

import (
	"fmt"
	"net/http"
	"path"
)

// PingResource sends a ping to a specified endpoint and resource ID to check its availability.
// This function utilizes the DoPing method from the httpclient package to perform the operation.
//
// Parameters:
//   - endpoint: The target API endpoint for the ping request. This should be a relative path that will be appended
//     to the base URL configured for the HTTP client.
//   - resourceID: The specific ID of the resource to ping. It will be appended to the endpoint to form the complete path.
//
// Returns:
//   - *http.Response: The HTTP response from the server. In case of a successful ping (200 OK), this response contains
//     the status code, headers, and body of the response. In case of errors, this will be the last received HTTP response.
//   - error: An error object indicating failure during the execution of the ping operation. This could be due to network
//     issues, server errors, or reaching the maximum number of retry attempts without receiving a 200 OK response.
//
// Usage:
// This function is intended for use in scenarios where it's necessary to confirm the availability or health of a specific
// resource within an endpoint.
func (c *Client) PingResource(endpoint, resourceID string) (*http.Response, error) {
	// Combine the endpoint and the resource ID to form the full path
	fullPath := path.Join(endpoint, resourceID)

	// Here we use a nil body and a nil output variable since we're just "pinging" the endpoint
	// and don't need to send or receive any specific data.
	resp, err := c.HTTP.DoPing("GET", fullPath, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping resource at %s: %v", fullPath, err)
	}

	return resp, nil
}
