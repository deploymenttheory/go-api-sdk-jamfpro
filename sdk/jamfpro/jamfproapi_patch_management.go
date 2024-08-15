// jamfproapi_patch_management.go
// Jamf Pro Api - Patch Management
// api reference: https://developer.jamf.com/jamf-pro/reference/post_v2-patch-management-accept-disclaimer
// Jamf Pro Api requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriPatchManagementDisclaimer = "/api/v2/patch-management-accept-disclaimer"

// AcceptPatchManagementDisclaimer accepts the Patch Management disclaimer
func (c *Client) AcceptPatchManagementDisclaimer() error {
	resp, err := c.HTTP.DoRequest("POST", uriPatchManagementDisclaimer, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to accept Patch Management disclaimer: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
