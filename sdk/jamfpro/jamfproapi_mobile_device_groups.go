// jamfproapi_mobile_device_groups.go
// Jamf Pro Api - Mobile Device Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-mobile-device-groups
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
)

const uriMobileDeviceGroupsV2 = "/api/v2/mobile-device-groups"

// List (v2)

// ResponseMobileDeviceGroupsListV2 represents the list of unified mobile device groups (v2).
type ResponseMobileDeviceGroupsListV2 struct {
	TotalCount int                           `json:"totalCount"`
	Results    []ResourceMobileDeviceGroupV2 `json:"results"`
}

// ResourceMobileDeviceGroupV2 represents a single unified mobile device group summary (v2).
type ResourceMobileDeviceGroupV2 struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	IsSmartGroup bool   `json:"isSmartGroup"`
}

// Shared response wrappers (v2)

// ResponseInventoryListMobileDeviceSearchResultsV2 represents the inventory list membership response for mobile device groups (v2).
type ResponseInventoryListMobileDeviceSearchResultsV2 struct {
	TotalCount int                                       `json:"totalCount"`
	Results    []SharedResourceInventoryListMobileDevice `json:"results"`
}

// Request (v2)

// RequestMobileDeviceGroupResetV2 represents the request body for erasing the members of a mobile device group (v2).
type RequestMobileDeviceGroupResetV2 struct {
	PreserveDataPlan       bool `json:"preserveDataPlan"`
	DisallowProximitySetup bool `json:"disallowProximitySetup"`
	ClearActivationLock    bool `json:"clearActivationLock"`
	ReturnToService        bool `json:"returnToService"`
}

// CRUD (v2)

// GetMobileDeviceGroupsV2 retrieves the list of unified mobile device groups using the v2 API.
func (c *Client) GetMobileDeviceGroupsV2(params url.Values) (*ResponseMobileDeviceGroupsListV2, error) {
	endpoint := uriMobileDeviceGroupsV2
	if params != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var results []ResourceMobileDeviceGroupV2
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &results)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "mobile device groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	out := ResponseMobileDeviceGroupsListV2{
		TotalCount: len(results),
		Results:    results,
	}

	return &out, nil
}

// GetMobileDeviceSmartGroupMembershipByIDV2 retrieves the membership of a smart mobile device group by ID using the v2 API.
func (c *Client) GetMobileDeviceSmartGroupMembershipByIDV2(id string) (*ResponseInventoryListMobileDeviceSearchResultsV2, error) {
	endpoint := fmt.Sprintf("%s/smart-group-membership/%s", uriMobileDeviceGroupsV2, id)

	var out ResponseInventoryListMobileDeviceSearchResultsV2
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device smart group membership", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetMobileDeviceStaticGroupMembershipByIDV2 retrieves the membership of a static mobile device group by ID using the v2 API.
func (c *Client) GetMobileDeviceStaticGroupMembershipByIDV2(id string) (*ResponseInventoryListMobileDeviceSearchResultsV2, error) {
	endpoint := fmt.Sprintf("%s/static-group-membership/%s", uriMobileDeviceGroupsV2, id)

	var out ResponseInventoryListMobileDeviceSearchResultsV2
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device static group membership", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// EraseMobileDeviceGroupByIDV2 erases the members of a mobile device group by ID using the v2 API. Returns no body on success (202).
func (c *Client) EraseMobileDeviceGroupByIDV2(id string, request RequestMobileDeviceGroupResetV2) error {
	endpoint := fmt.Sprintf("%s/%s/erase", uriMobileDeviceGroupsV2, id)

	resp, err := c.HTTP.DoRequest("POST", endpoint, request, nil)
	if err != nil || resp.StatusCode != 202 {
		return fmt.Errorf(errMsgFailedActionByID, "mobile device group erase", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
