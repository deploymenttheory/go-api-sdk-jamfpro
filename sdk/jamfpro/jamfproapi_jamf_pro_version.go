// jamfproapi_jamf_pro_version.go
// Jamf Pro Api - Jamf Pro Version
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-jamf-pro-version
// Classic API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriJamfProVersion = "/api/v1/jamf-pro-version"

// Response

type ResponseJamfProVersion struct {
	Version *string `json:"Version,omitempty"`
}

// CRUD

func (c *Client) GetJamfProVersion() (*ResponseJamfProVersion, error) {
	endpoint := uriJamfProVersion

	var version ResponseJamfProVersion
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &version)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf Pro version: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &version, nil
}
