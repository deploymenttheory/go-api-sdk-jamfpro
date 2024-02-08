// jamfproapi_client_checkin.go
// Jamf Pro Api - Client Checkin
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v3-check-in
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const UriClientCheckinSettings = "/api/v3/check-in"

// Structs

// Resource

type ResourceClientCheckinSettings struct {
	CheckInFrequency                 int  `json:"checkInFrequency"`
	CreateHooks                      bool `json:"createHooks"`
	HookLog                          bool `json:"hookLog"`
	HookPolicies                     bool `json:"hookPolicies"`
	CreateStartupScript              bool `json:"createStartupScript"`
	StartupLog                       bool `json:"startupLog"`
	StartupPolicies                  bool `json:"startupPolicies"`
	StartupSsh                       bool `json:"startupSsh"`
	EnableLocalConfigurationProfiles bool `json:"enableLocalConfigurationProfiles"`
}

// CRUD

func (c *Client) GetClientCheckinSettings() (*ResourceClientCheckinSettings, error) {
	endpoint := UriClientCheckinSettings

	var out ResourceClientCheckinSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "client checkin settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

func (c *Client) UpdateClientCheckinSettings(settingsUpdate ResourceClientCheckinSettings) (*ResourceClientCheckinSettings, error) {
	endpoint := UriClientCheckinSettings

	var out ResourceClientCheckinSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, settingsUpdate, &out, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "client checkin settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}
