// jamfProInformation.go
// Jamf Pro Api
// Jamf Pro API requires the structs to support JSON.

package jamfpro

const uriJamfProInformation = "/api/v2/jamf-pro-information"

type ResponseJamfProInformation struct {
	IsVppTokenEnabled         bool `json:"isVppTokenEnabled"`
	IsDepAccountEnabled       bool `json:"isDepAccountEnabled"`
	IsByodEnabled             bool `json:"isByodEnabled"`
	IsUserMigrationEnabled    bool `json:"isUserMigrationEnabled"`
	IsCloudDeploymentsEnabled bool `json:"isCloudDeploymentsEnabled"`
	IsPatchEnabled            bool `json:"isPatchEnabled"`
	IsSsoSamlEnabled          bool `json:"isSsoSamlEnabled"`
	IsSmtpEnabled             bool `json:"isSmtpEnabled"`
}

func (c *Client) GetJamfProInformation() (*ResponseJamfProInformation, error) {
	var out *ResponseJamfProInformation
	err := c.DoRequest("GET", uriJamfProInformation, nil, nil, &out)
	return out, err
}
