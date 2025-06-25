// jamfproapi_access_management_settings.go
// Jamf Pro API - Access Management
// api reference: /v4/enrollment/access-management
package jamfpro

import (
	"fmt"
)

const (
	uriAccessManagementSettings = "/api/v4/enrollment/access-management"
)

// Structs

type ResourceAccessManagementSettings struct {
	AutomatedDeviceEnrollmentServerUuid string `json:"automatedDeviceEnrollmentServerUuid"`
}

// CRUD

// GetAccessManagementSettings retrieves the current Access Management settings
func (c *Client) GetAccessManagementSettings() (*ResourceAccessManagementSettings, error) {
	endpoint := uriAccessManagementSettings
	var out ResourceAccessManagementSettings

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "access management settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateAccessManagementSettings configures the Access Management settings on the Jamf Pro server.
func (c *Client) CreateAccessManagementSettings(settingsCreate ResourceAccessManagementSettings) (*ResourceAccessManagementSettings, error) {
	endpoint := uriAccessManagementSettings

	var out ResourceAccessManagementSettings
	resp, err := c.HTTP.DoRequest("POST", endpoint, settingsCreate, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "access management settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
