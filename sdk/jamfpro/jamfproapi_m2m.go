// jamfproapi_m2m.go
// Jamf Pro Api - M2M (Machine to Machine)
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-m2m-tenant-id
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import "fmt"

const uriM2m = "/api/v1/m2m"

// ResponseM2mTenantID represents the M2M-sourced tenant ID information.
type ResponseM2mTenantID struct {
	TenantID string `json:"tenantId"`
}

// GetM2mTenantID retrieves the M2M-sourced tenant ID.
func (c *Client) GetM2mTenantID() (*ResponseM2mTenantID, error) {
	endpoint := fmt.Sprintf("%s/tenant-id", uriM2m)

	var out ResponseM2mTenantID
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "m2m tenant id", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
