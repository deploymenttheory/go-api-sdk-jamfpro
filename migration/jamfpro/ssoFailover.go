// ssoFailover.go
package jamfpro

const uriSSOFailover = "/api/v1/sso/failover"

type FailoverResponse struct {
	FailoverURL    string `json:"failoverUrl"`
	GenerationTime int64  `json:"generationTime"`
}

// GetSSOFailoverSettings fetches SSO failover settings from Jamf Pro
func (c *Client) GetSSOFailoverSettings() (*FailoverResponse, error) {
	var out FailoverResponse
	_, err := c.http.DoRequest("GET", uriSSOFailover, nil, &out)
	return &out, err
}
