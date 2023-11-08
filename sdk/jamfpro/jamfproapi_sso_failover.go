// jamfproapi_sso_failover.go
// Jamf Pro Api - SSO Failover URL
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-sso-failover
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import "fmt"

const uriSSOFailover = "/api/v1/sso/failover"

type SSOFailoverResponse struct {
	FailoverURL    string `json:"failoverUrl"`
	GenerationTime int64  `json:"generationTime"`
}

// GetSSOFailoverSettings fetches SSO failover settings from Jamf Pro
func (c *Client) GetSSOFailoverSettings() (*SSOFailoverResponse, error) {
	var out SSOFailoverResponse

	resp, err := c.HTTP.DoRequest("GET", uriSSOFailover, nil, &out)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf("Failed to fetch SSO failover settings: %v\n", err)
		return nil, err
	}

	return &out, nil
}

// UpdateFailoverUrl regenerates the failover URL by changing the failover key to a new one and returns the new failover settings.
func (c *Client) UpdateFailoverUrl() (*SSOFailoverResponse, error) {
	var out SSOFailoverResponse

	// Extend the existing uriSSOFailover constant for the failover generation endpoint
	endpoint := uriSSOFailover + "/generate"

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &out)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf("Failed to regenerate the failover URL: %v\n", err)
		return nil, err
	}

	return &out, nil
}
