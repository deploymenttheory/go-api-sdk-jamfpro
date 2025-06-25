// jamfproapi_self_service_plus.go
// Jamf Pro API - Self Service Plus Settings
// api reference: /v1/self-service-plus/settings
package jamfpro

import (
	"fmt"
)

const (
	uriSelfServicePlusSettings = "/api/v1/self-service-plus/settings"
)

// Structs

type ResourceSelfServicePlusSettings struct {
	Enabled bool `json:"enabled"`
}

// CRUD

// GetSelfServicePlusSettings retrieves the current Self Service Plus settings
func (c *Client) GetSelfServicePlusSettings() (*ResourceSelfServicePlusSettings, error) {
	endpoint := uriSelfServicePlusSettings
	var out ResourceSelfServicePlusSettings

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "self service plus settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateSelfServicePlusSettings updates the Self Service Plus settings
func (c *Client) UpdateSelfServicePlusSettings(settings ResourceSelfServicePlusSettings) error {
	endpoint := uriSelfServicePlusSettings

	resp, _ := c.HTTP.DoRequest("PUT", endpoint, settings, nil)

	if resp == nil {
		return fmt.Errorf("failed to update Self Service Plus settings: received nil response")
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to update Self Service Plus settings: unexpected status code %d", resp.StatusCode)
	}

	return nil
}
