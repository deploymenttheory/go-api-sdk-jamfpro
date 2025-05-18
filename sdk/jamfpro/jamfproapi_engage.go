// jamfproapi_engage.go
// Jamf Pro Api - Engage
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-engage
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const UriEngageSettings = "/api/v2/engage"

// Structs

// Resource

type ResourceEngageSettings struct {
	IsEnabled bool `json:"isEnabled"`
}

// CRUD

// GetEngageSettings retrieves the Engage settings from the Jamf Pro server.
func (c *Client) GetEngageSettings() (*ResourceEngageSettings, error) {
	endpoint := UriEngageSettings

	var out ResourceEngageSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "engage settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateEngageSettings updates the Engage settings on the Jamf Pro server.
func (c *Client) UpdateEngageSettings(settingsUpdate ResourceEngageSettings) (*ResourceEngageSettings, error) {
	endpoint := UriEngageSettings

	var out ResourceEngageSettings
	resp, err := c.HTTP.DoRequest("PUT", endpoint, settingsUpdate, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "engage settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
