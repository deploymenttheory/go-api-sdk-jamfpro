// classicapi_activation_code.go
// Jamf Pro Classic Api - activationcode
// api reference: https://developer.jamf.com/jamf-pro/reference/activationcode
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIActivationCode = "/JSSResource/activationcode"

// Responses

// ResponseActivationCode represents the structure of the response for an activation code.
type ResourceActivationCode struct {
	OrganizationName string `xml:"organization_name"`
	Code             string `xml:"code"`
}

// CRUD

// GetActivationCode retrieves the activation code.
func (c *Client) GetActivationCode() (*ResourceActivationCode, error) {
	endpoint := uriAPIActivationCode

	var activationCode ResourceActivationCode
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &activationCode)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch activation code: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &activationCode, nil
}

// UpdateActivationCode updates the activation code.
func (c *Client) UpdateActivationCode(organizationName, code string) error {
	endpoint := uriAPIActivationCode

	requestBody := struct {
		XMLName xml.Name `xml:"activation_code"`
		ResourceActivationCode
	}{
		ResourceActivationCode: ResourceActivationCode{
			OrganizationName: organizationName,
			Code:             code,
		},
	}

	_, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, nil)
	if err != nil {
		return fmt.Errorf("failed to update activation code: %v", err)
	}

	return nil
}
