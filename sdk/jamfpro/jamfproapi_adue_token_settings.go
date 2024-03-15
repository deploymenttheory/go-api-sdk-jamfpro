// jamfproapi_account_driven_user_enrollment_token_settings.go
// Jamf Pro Api - Enrollment
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-enrollment-history
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriUserEnrollmentTokenSettings = "/api/v1/adue-session-token-settings"

// structs

type ResourceADUETokenSettings struct {
	Enabled                   bool `json:"enabled"`
	ExpirationIntervalDays    int  `json:"expirationIntervalDays,omitempty"`
	ExpirationIntervalSeconds int  `json:"expirationIntervalSeconds,omitempty"`
}

// CRUD

func (c *Client) GetADUESessionTokenSettings() (*ResourceADUETokenSettings, error) {
	endpoint := uriUserEnrollmentTokenSettings
	var out ResourceADUETokenSettings

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "ADUE token settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}

func (c *Client) UpdateADUESessionTokenSettings(updatedSettings ResourceADUETokenSettings) (*ResourceADUETokenSettings, error) {
	endpoint := uriUserEnrollmentTokenSettings
	var out ResourceADUETokenSettings

	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedSettings, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "ADUE token settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
