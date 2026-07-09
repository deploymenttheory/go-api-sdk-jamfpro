// jamfproapi_last_login.go
// Jamf Pro Api - Last Login
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-last-login
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriLastLogin = "/api/v1/last-login"

// ResponseLastLogin represents the timestamp of the current user's last login.
type ResponseLastLogin struct {
	LastLogin string `json:"lastLogin"`
}

// GetLastLogin retrieves the timestamp of the last login for the current user.
func (c *Client) GetLastLogin() (*ResponseLastLogin, error) {
	endpoint := uriLastLogin

	var out ResponseLastLogin
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "last login", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
