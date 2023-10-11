// jamfProVersion.go
// Jamf Pro Api
// Jamf Pro API requires the structs to support JSON.

package jamfpro

const uriJamfProVersion = "/api/v1/jamf-pro-version"

type ResponseJamfProVersion struct {
	Version *string `json:"Version,omitempty"`
}

func (c *Client) GetJamfProVersion() (*ResponseJamfProVersion, error) {
	var out *ResponseJamfProVersion
	err := c.DoRequest("GET", uriJamfProVersion, nil, nil, &out)
	return out, err
}
