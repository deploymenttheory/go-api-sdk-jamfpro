// jamfproapi_sso_failover.go
// uses jamf pro api and therefore structs support JSON only.

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
		fmt.Printf("Failed to regenerate the failover URL", "Error", err)
		return nil, err
	}

	return &out, nil
}
