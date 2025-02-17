// jamfproapi_policy_properties.go
// Jamf Pro Api - Policy Properties appends to Client Checkin settings
// why on earth is this two different api endpoints ?
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-policy-properties
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const UriPolicyProperties = "/api/v1/policy-properties"

// Structs

// Resource

type ResourcePolicyProperties struct {
	PoliciesRequireNetworkStateChange bool `json:"policiesRequireNetworkStateChange"`
	AllowNetworkStateChangeTriggers   bool `json:"allowNetworkStateChangeTriggers"`
}

// CRUD

func (c *Client) GetPolicyProperties() (*ResourcePolicyProperties, error) {
	endpoint := UriPolicyProperties

	var out ResourcePolicyProperties
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "policy properties", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

func (c *Client) UpdatePolicyProperties(settingsUpdate ResourcePolicyProperties) (*ResourcePolicyProperties, error) {
	endpoint := UriPolicyProperties

	var out ResourcePolicyProperties
	resp, err := c.HTTP.DoRequest("PUT", endpoint, settingsUpdate, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdate, "policy properties", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil

}
