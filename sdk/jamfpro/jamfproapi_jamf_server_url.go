// jamfproapi_jamf_server_url.go
// Jamf Pro Api - JAMF Server URL
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-server-url
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriJamfProServerUrl = "/api/v1/jamf-pro-server-url"

// Structs

type ResourceJamfProServerURL struct {
	URL                    string `json:"url"`
	UnsecuredEnrollmentUrl string `json:"unsecuredEnrollmentUrl"`
}

// CRUD

// Returns ResourceJamfProServerURL
func (c *Client) GetJamfProServerUrlSettings() (*ResourceJamfProServerURL, error) {
	endpoint := uriJamfProServerUrl
	var out ResourceJamfProServerURL

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "jamf pro server url settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// Updates Jamf Pro Server URL settings
func (c *Client) UpdateJamfProServerUrlSettings(updatedSettings ResourceJamfProServerURL) (*ResourceJamfProServerURL, error) {
	endpoint := uriJamfProServerUrl
	var out ResourceJamfProServerURL

	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedSettings, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "jamf pro server url settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
