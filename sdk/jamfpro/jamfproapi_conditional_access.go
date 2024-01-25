// jamfproapi_conditional_access.go
// Jamf Pro Api - Conditional Access
// api reference: undocumented
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriConditionalAccess = "/api/v1/conditional-access"

// ResourceConditionalAccessDeviceComplianceStatus represents the resource object.
type ResourceConditionalAccessDeviceComplianceStatus struct {
	SharedDeviceFeatureEnabled bool `json:"sharedDeviceFeatureEnabled"`
}

// GetConditionalAccessDeviceComplianceFeatureEnablement retrieves the enablement state of the device compliance settinfs for CA.
func (c *Client) GetConditionalAccessDeviceComplianceFeatureEnablement() (*ResourceConditionalAccessDeviceComplianceStatus, error) {
	endpoint := uriConditionalAccess + "/device-compliance/feature-toggle"

	var defaultCloudIdPServer ResourceConditionalAccessDeviceComplianceStatus
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &defaultCloudIdPServer)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "Conditional Access Device Compliance", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &defaultCloudIdPServer, nil
}
