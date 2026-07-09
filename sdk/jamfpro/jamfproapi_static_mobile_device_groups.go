// jamfproapi_static_mobile_device_groups.go
// Jamf Pro Api - Static Mobile Device Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-mobile-device-groups-static-groups
// Jamf Pro API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const (
	uriAPIV1MobileDeviceGroups = "/api/v1/mobile-device-groups"
)

// Request

// ResourceStaticMobileDeviceGroupV1 represents the request structure for creating a Static Mobile Device Group
type ResourceStaticMobileDeviceGroupV1 struct {
	GroupName        string                                         `json:"groupName"`
	GroupDescription string                                         `json:"groupDescription,omitempty"`
	Assignments      []ResourceStaticMobileDeviceGroupAssignmentsV1 `json:"assignments"`
	SiteId           string                                         `json:"siteId"`
}

// ResourceStaticMobileDeviceGroupAssignmentsV1 represents the assignments structure for Static Mobile Device Groups
type ResourceStaticMobileDeviceGroupAssignmentsV1 struct {
	MobileDeviceID string `json:"mobileDeviceId"`
	Selected       bool   `json:"selected"`
}

// Response

// ResponseStaticMobileDeviceGroupsListV1 represents the paginated response for Static Mobile Device Groups v1
type ResponseStaticMobileDeviceGroupsListV1 struct {
	TotalCount int                                 `json:"totalCount"`
	Results    []ResponseStaticMobileDeviceGroupV1 `json:"results"`
}

// ResponseStaticMobileDeviceGroupV1 represents a Static Mobile Device Group in v1 API
type ResponseStaticMobileDeviceGroupV1 struct {
	GroupID          string `json:"groupId"`
	GroupName        string `json:"groupName"`
	GroupDescription string `json:"groupDescription"`
	SiteID           string `json:"siteId"`
	Count            int    `json:"count"`
}

// ResponseStaticMobileDeviceGroupCreateV1 represents the response structure for creating a Static Mobile Device Group
type ResponseStaticMobileDeviceGroupCreateV1 struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD

// GetStaticMobileDeviceGroupsV1 retrieves a paginated list of all Static Mobile Device Groups using V1 API
func (c *Client) GetStaticMobileDeviceGroupsV1(params url.Values) (*ResponseStaticMobileDeviceGroupsListV1, error) {
	resp, err := c.DoPaginatedGet(fmt.Sprintf("%s/static-groups", uriAPIV1MobileDeviceGroups), params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "static mobile device groups", err)
	}

	var out ResponseStaticMobileDeviceGroupsListV1
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResponseStaticMobileDeviceGroupV1
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "static mobile device group", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetStaticMobileDeviceGroupByIDV1 retrieves a specific Static Mobile Device Group by ID
func (c *Client) GetStaticMobileDeviceGroupByIDV1(id string) (*ResponseStaticMobileDeviceGroupV1, error) {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriAPIV1MobileDeviceGroups, id)

	var response ResponseStaticMobileDeviceGroupV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "static mobile device group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetStaticMobileDeviceGroupByNameV1 retrieves a Static Mobile Device Group by name
func (c *Client) GetStaticMobileDeviceGroupByNameV1(groupName string) (*ResponseStaticMobileDeviceGroupV1, error) {
	params := url.Values{}
	params.Set("filter", fmt.Sprintf("groupName==\"%s\"", groupName))

	groups, err := c.GetStaticMobileDeviceGroupsV1(params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "static mobile device groups", err)
	}

	if len(groups.Results) == 0 {
		return nil, fmt.Errorf(errMsgFailedGetByName, "static mobile device group", groupName, errMsgNoName)
	}

	return &groups.Results[0], nil
}

// CreateStaticMobileDeviceGroupV1 creates a new Static Mobile Device Group
func (c *Client) CreateStaticMobileDeviceGroupV1(request ResourceStaticMobileDeviceGroupV1) (*ResponseStaticMobileDeviceGroupCreateV1, error) {
	endpoint := fmt.Sprintf("%s/static-groups", uriAPIV1MobileDeviceGroups)
	if request.Assignments == nil {
		request.Assignments = []ResourceStaticMobileDeviceGroupAssignmentsV1{}
	}

	var response ResponseStaticMobileDeviceGroupCreateV1
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "static mobile device group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateStaticMobileDeviceGroupByIDV1 updates an existing Static Mobile Device Group by ID
func (c *Client) UpdateStaticMobileDeviceGroupByIDV1(id string, request ResourceStaticMobileDeviceGroupV1) (*ResourceStaticMobileDeviceGroupV1, error) {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriAPIV1MobileDeviceGroups, id)
	if request.Assignments == nil {
		request.Assignments = []ResourceStaticMobileDeviceGroupAssignmentsV1{}
	}

	var response ResourceStaticMobileDeviceGroupV1
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "static mobile device group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteStaticMobileDeviceGroupByIDV1 deletes a Static Mobile Device Group by ID
func (c *Client) DeleteStaticMobileDeviceGroupByIDV1(id string) error {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriAPIV1MobileDeviceGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "static mobile device group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// ---------------------------------------------------------------------------
// Static Mobile Device Groups v2 (/api/v2/mobile-device-groups/static-groups)
// ---------------------------------------------------------------------------

// List (v2)

// ResponseMobileDeviceStaticGroupsListV2 represents the search results for static mobile device groups (v2).
type ResponseMobileDeviceStaticGroupsListV2 struct {
	TotalCount int                                 `json:"totalCount"`
	Results    []ResourceMobileDeviceStaticGroupV2 `json:"results"`
}

// Resource / Response (v2)

// ResourceMobileDeviceStaticGroupV2 represents a static mobile device group (v2).
type ResourceMobileDeviceStaticGroupV2 struct {
	GroupID          string `json:"groupId"`
	GroupName        string `json:"groupName"`
	GroupDescription string `json:"groupDescription"`
	SiteID           string `json:"siteId"`
	Count            int    `json:"count"`
}

// Request (v2)

// ResourceStaticGroupAssignmentV2 represents the request/response body for creating or updating a static mobile device group (v2).
type ResourceStaticGroupAssignmentV2 struct {
	GroupID          string                                `json:"groupId,omitempty"`
	GroupName        string                                `json:"groupName"`
	GroupDescription string                                `json:"groupDescription,omitempty"`
	SiteID           string                                `json:"siteId,omitempty"`
	Assignments      []SubsetMobileDeviceGroupAssignmentV2 `json:"assignments,omitempty"`
}

// SubsetMobileDeviceGroupAssignmentV2 represents a single mobile device assignment for a static group (v2).
type SubsetMobileDeviceGroupAssignmentV2 struct {
	MobileDeviceID string `json:"mobileDeviceId"`
	Selected       bool   `json:"selected"`
}

// ResponseMobileDeviceStaticGroupCreateV2 represents the response for creating a static mobile device group (v2).
type ResponseMobileDeviceStaticGroupCreateV2 struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD (v2)

// GetMobileDeviceStaticGroupsV2 retrieves the list of static mobile device groups using the v2 API.
func (c *Client) GetMobileDeviceStaticGroupsV2(params url.Values) (*ResponseMobileDeviceStaticGroupsListV2, error) {
	endpoint := fmt.Sprintf("%s/static-groups", uriMobileDeviceGroupsV2)
	if params != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var out ResponseMobileDeviceStaticGroupsListV2
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "mobile device static groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetMobileDeviceStaticGroupByIDV2 retrieves a specific static mobile device group by ID using the v2 API.
func (c *Client) GetMobileDeviceStaticGroupByIDV2(id string) (*ResourceMobileDeviceStaticGroupV2, error) {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriMobileDeviceGroupsV2, id)

	var out ResourceMobileDeviceStaticGroupV2
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device static group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateMobileDeviceStaticGroupV2 creates a new static mobile device group using the v2 API.
func (c *Client) CreateMobileDeviceStaticGroupV2(request ResourceStaticGroupAssignmentV2) (*ResponseMobileDeviceStaticGroupCreateV2, error) {
	endpoint := fmt.Sprintf("%s/static-groups", uriMobileDeviceGroupsV2)

	var out ResponseMobileDeviceStaticGroupCreateV2
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "mobile device static group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateMobileDeviceStaticGroupByIDV2 updates an existing static mobile device group by ID using the v2 API.
func (c *Client) UpdateMobileDeviceStaticGroupByIDV2(id string, request ResourceStaticGroupAssignmentV2) (*ResourceStaticGroupAssignmentV2, error) {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriMobileDeviceGroupsV2, id)

	var out ResourceStaticGroupAssignmentV2
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, request, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "mobile device static group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// DeleteMobileDeviceStaticGroupByIDV2 deletes a static mobile device group by ID using the v2 API.
func (c *Client) DeleteMobileDeviceStaticGroupByIDV2(id string) error {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriMobileDeviceGroupsV2, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device static group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
