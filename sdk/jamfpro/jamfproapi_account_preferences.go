// jamfproapi_account_preferences.go
// Jamf Pro Api - Account Preferences
// api reference:
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

// Responses

const uriAccountPreferences = "/v2/account-preferences"

type ResponseAccountPreferencesList struct {
	TotalCount int                          `json:"totalCount"`
	Results    []ResourceAccountPreferences `json:"results"`
}

type ResponseAccountPreferencesCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

type ResourceAccountPreferences struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetAccountPreferencesretrieves the default server configuration for the Cloud Identity Provider.
func (c *Client) GetAccountPreferences() (*ResourceAccountPreferences, error) {
	endpoint := uriAccountPreferences

	var accountPreferences ResourceAccountPreferences
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &accountPreferences)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Account Preferences", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &accountPreferences, nil
}
