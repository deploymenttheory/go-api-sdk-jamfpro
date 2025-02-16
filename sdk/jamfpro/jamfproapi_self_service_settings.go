// jamfproapi_self_service_settings.go
// Jamf Pro Api - Self Service Settings
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-settings
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriSelfServiceSettings = "/api/v1/self-service/settings"

// Resource

type ResourceSelfServiceSettings struct {
	InstallSettings       InstallSettings       `json:"installSettings"`
	LoginSettings         LoginSettings         `json:"loginSettings"`
	ConfigurationSettings ConfigurationSettings `json:"configurationSettings"`
}

type InstallSettings struct {
	InstallAutomatically bool   `json:"installAutomatically"`
	InstallLocation      string `json:"installLocation"`
}

type LoginSettings struct {
	UserLoginLevel  string `json:"userLoginLevel"`
	AllowRememberMe bool   `json:"allowRememberMe"`
	UseFido2        bool   `json:"useFido2"`
	AuthType        string `json:"authType"`
}

type ConfigurationSettings struct {
	NotificationsEnabled  bool   `json:"notificationsEnabled"`
	AlertUserApprovedMdm  bool   `json:"alertUserApprovedMdm"`
	DefaultLandingPage    string `json:"defaultLandingPage"`
	DefaultHomeCategoryId int    `json:"defaultHomeCategoryId"`
	BookmarksName         string `json:"bookmarksName"`
}

// CRUD

// GetSelfServiceSettings retrives a self service list.
func (c *Client) GetSelfServiceSettings() (*ResourceSelfServiceSettings, error) {
	endpoint := uriSelfServiceSettings

	var settings ResourceSelfServiceSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &settings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "self service settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &settings, nil
}

// UpdateSelfServiceSettings updates the self service settings.
func (c *Client) UpdateSelfServiceSettings(settingsUpdate *ResourceSelfServiceSettings) (*ResourceSelfServiceSettings, error) {
	endpoint := uriSelfServiceSettings

	var updatedSettings ResourceSelfServiceSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, settingsUpdate, &updatedSettings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "self service settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSettings, nil
}

// Please note that the Create and Delete functions are not implemented here as self service settings only allow retrieval and update operations.
