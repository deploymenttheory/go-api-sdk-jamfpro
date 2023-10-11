// activationCode.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import "fmt"

const uriActivationCode = "/JSSResource/activationcode"

type ActivationCode struct {
	OrganizationName string `json:"organization_name" xml:"organization_name"`
	Code             string `json:"code" xml:"code"`
}

func (c *Client) GetActivationCode() (*ActivationCode, error) {
	out := &ActivationCode{}
	err := c.DoRequest("GET", uriActivationCode, nil, nil, out)
	if err != nil {
		return nil, fmt.Errorf("failed to get activation code: %v", err)
	}

	return out, nil
}

func (c *Client) UpdateActivationCode() error {
	out := &ActivationCode{}
	err := c.DoRequest("PUT", uriActivationCode, nil, nil, out)
	if err != nil {
		return fmt.Errorf("failed to update activation code: %v", err)
	}

	return nil
}
