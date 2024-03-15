// jamfproapi_local_admin_password.go
// Jamf Pro Api - JAMF local administrator password (LAPS)
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-local-admin-password-pending-rotations
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import "fmt"

const uriLocalAdminPassword = "/api/v2/local-admin-password"

// Resource
type ResourceLocalAdminPasswordSettings struct {
	AutoDeployEnabled        bool `json:"autoDeployEnabled"`
	PasswordRotationTime     int  `json:"passwordRotationTime"`
	AutoRotateEnabled        bool `json:"autoRotateEnabled"`
	AutoRotateExpirationTime int  `json:"autoRotateExpirationTime"`
}

// GetLocalAdminPasswordSettings retrieves current Jamf Pro LAPS settings
func (c *Client) GetLocalAdminPasswordSettings() (*ResourceLocalAdminPasswordSettings, error) {
	endpoint := uriLocalAdminPassword + "/settings"
	var out ResourceLocalAdminPasswordSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "LAPS settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateLocalAdminPasswordSettings updates the current Jamf Pro LAPS settings
func (c *Client) UpdateLocalAdminPasswordSettings(settings *ResourceLocalAdminPasswordSettings) error {
	endpoint := uriLocalAdminPassword + "/settings"

	var handleResponse struct{}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &settings, &handleResponse)
	if err != nil {
		return fmt.Errorf(errMsgFailedUpdate, "LAPS settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
