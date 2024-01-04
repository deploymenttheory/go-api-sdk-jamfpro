// jamfproapi_jamf_sso_failover.go
// Jamf Pro Api - JAMF SSO Failover URL
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-sso-failover
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import "fmt"

const uriSSOFailover = "/api/v1/sso/failover"

type ResponseSSOFailover struct {
	FailoverURL    string `json:"failoverUrl"`
	GenerationTime int64  `json:"generationTime"`
}

// GetSSOFailoverSettings fetches SSO failover settings from Jamf Pro
func (c *Client) GetSSOFailoverSettings() (*ResponseSSOFailover, error) {
	var out ResponseSSOFailover

	resp, err := c.HTTP.DoRequest("GET", uriSSOFailover, nil, &out)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf(errMsgFailedGet, "sso failover settings", err)
		return nil, err
	}

	return &out, nil
}

// UpdateFailoverUrl regenerates the failover URL by changing the failover key to a new one and returns the new failover settings.
func (c *Client) UpdateFailoverUrl() (*ResponseSSOFailover, error) {
	var out ResponseSSOFailover

	endpoint := uriSSOFailover + "/generate"

	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &out)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf(errMsgFailedUpdate, "sso failover url", err)
		return nil, err
	}

	return &out, nil
}
