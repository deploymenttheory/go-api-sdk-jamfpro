// jamfproapi_jamf_pro_information.go
// Jamf Pro Api - Jamf Pro Version
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-jamf-pro-information
// Classic API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriJamfProInformation = "/api/v2/jamf-pro-information"

// Response

type ResponseJamfProInformation struct {
	VppTokenEnabled         *bool `json:"vppTokenEnabled,omitempty"`
	DepAccountEnabled       *bool `json:"depAccountEnabled,omitempty"`
	ByodEnabled             *bool `json:"byodEnabled,omitempty"`
	UserMigrationEnabled    *bool `json:"userMigrationEnabled,omitempty"`
	CloudDeploymentsEnabled *bool `json:"cloudDeploymentsEnabled,omitempty"`
	PatchEnabled            *bool `json:"patchEnabled,omitempty"`
	SsoSamlEnabled          *bool `json:"ssoSamlEnabled,omitempty"`
	SmtpEnabled             *bool `json:"smtpEnabled,omitempty"`
}

// CRUD

func (c *Client) GetJamfProInformation() (*ResponseJamfProInformation, error) {
	endpoint := uriJamfProInformation

	var info ResponseJamfProInformation
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &info)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf Pro information: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &info, nil
}
