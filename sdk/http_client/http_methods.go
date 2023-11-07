// http_methods.go
package http_client

import "net/http"

// Get sends a GET request to the specified endpoint and unmarshals the response into 'out'.
// The caller is responsible for closing the response body.
func (c *Client) Get(endpoint string, out interface{}) (*http.Response, error) {
	c.logger.Info("Sending GET request", "endpoint", endpoint)

	resp, err := c.DoRequest(http.MethodGet, endpoint, nil, out)
	if err != nil {
		c.logger.Error("GET request failed", "endpoint", endpoint, "error", err)
		return nil, err
	}
	return resp, nil
}

// Post sends a POST request to the specified endpoint with the provided body and unmarshals the response into 'out'.
// The caller is responsible for closing the response body.
func (c *Client) Post(endpoint string, body, out interface{}) (*http.Response, error) {
	c.logger.Info("Sending POST request", "endpoint", endpoint, "body", body)

	resp, err := c.DoRequest(http.MethodPost, endpoint, body, out)
	if err != nil {
		c.logger.Error("POST request failed", "endpoint", endpoint, "error", err)
		return nil, err
	}
	return resp, nil
}

// Put sends a PUT request to the specified endpoint with the provided body and unmarshals the response into 'out'.
// The caller is responsible for closing the response body.
func (c *Client) Put(endpoint string, body, out interface{}) (*http.Response, error) {

	c.logger.Debug("Sending PUT request", "endpoint", endpoint, "body", body)

	resp, err := c.DoRequest(http.MethodPut, endpoint, body, out)
	if err != nil {
		c.logger.Error("PUT request failed", "endpoint", endpoint, "error", err)
		return nil, err
	}
	return resp, nil
}

// Delete sends a DELETE request to the specified endpoint and unmarshals the response into 'out'.
// The caller is responsible for closing the response body.
func (c *Client) Delete(endpoint string, out interface{}) (*http.Response, error) {
	c.logger.Debug("Sending DELETE request", "endpoint", endpoint)

	resp, err := c.DoRequest(http.MethodDelete, endpoint, nil, out)
	if err != nil {
		c.logger.Error("DELETE request failed", "endpoint", endpoint, "error", err)
		return nil, err
	}
	return resp, nil
}

// Patch sends a PATCH request to the specified endpoint with the provided body and unmarshals the response into 'out'.
// The caller is responsible for closing the response body.
func (c *Client) Patch(endpoint string, body, out interface{}) (*http.Response, error) {
	c.logger.Debug("Sending PATCH request", "endpoint", endpoint, "body", body)

	resp, err := c.DoRequest(http.MethodPatch, endpoint, body, out)
	if err != nil {
		c.logger.Error("PATCH request failed", "endpoint", endpoint, "error", err)
		return nil, err
	}
	return resp, nil
}
