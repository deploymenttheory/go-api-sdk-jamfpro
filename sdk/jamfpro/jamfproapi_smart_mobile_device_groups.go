// jamfproapi_smart_mobiledevice_groups.go
// Jamf Pro Api - Smart Mobile Device Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-mobiledevice-groups-smart-groups
// Jamf Pro API requires the structs to support an JSON data structure.

/*
Shared data structure resources in this endpoint:
- SharedSubsetCriteriaJamfProAPI
*/

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

// Request

// ResourceSmartMobileDeviceGroupV1 represents the request structure for creating a Smart Mobile Device Group
type ResourceSmartMobileDeviceGroupV1 struct {
	GroupName        string                           `json:"groupName"`
	GroupDescription string                           `json:"groupDescription,omitempty"`
	Criteria         []SharedSubsetCriteriaJamfProAPI `json:"criteria"`
	SiteId           *string                          `json:"siteId,omitempty"`
}

// ResponseSmartMobileDeviceGroupMembershipV1 represents the response structure for retrieving the membership of a Smart Mobile Device Group
type ResponseSmartMobileDeviceGroupMembershipV1 struct {
	TotalCount int                                       `json:"totalCount"`
	Results    []SharedResourceInventoryListMobileDevice `json:"results"`
}

// ResponseSmartMobileDeviceGroupListItemV1 represents individual Smart Mobile Device Group items
type ResponseSmartMobileDeviceGroupListItemV1 struct {
	GroupID          string `json:"groupId"`
	GroupName        string `json:"groupName"`
	GroupDescription string `json:"groupDescription"`
	SiteID           string `json:"siteId"`
	Count            int    `json:"count"`
}

// ResponseSmartMobileDeviceGroupsListV1 represents the paginated response for Smart Mobile Device Groups v1
type ResponseSmartMobileDeviceGroupsListV1 struct {
	TotalCount int                                        `json:"totalCount"`
	Results    []ResponseSmartMobileDeviceGroupListItemV1 `json:"results"`
}

// ResourceSmartMobileDeviceGroupV1 represents a Smart Mobile Device Group in v1 API
type ResponseSmartMobileDeviceGroupV1 struct {
	GroupID          string                           `json:"groupId"`
	SiteId           string                           `json:"siteId"`
	GroupName        string                           `json:"groupName"`
	GroupDescription string                           `json:"groupDescription"`
	Count            int                              `json:"count"`
	Criteria         []SharedSubsetCriteriaJamfProAPI `json:"criteria"`
}

// ResponseSmartMobileDeviceGroupCreateV1 represents the response structure for creating a Smart Mobile Device Group
type ResponseSmartMobileDeviceGroupCreateV1 struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD

// GetSmartMobileDeviceGroupsV1 retrieves a paginated list of all Smart Mobile Device Groups using V1 API
func (c *Client) GetSmartMobileDeviceGroupsV1(params url.Values) (*ResponseSmartMobileDeviceGroupsListV1, error) {
	resp, err := c.DoPaginatedGet(fmt.Sprintf("%s/smart-groups", uriAPIV1MobileDeviceGroups), params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "smart mobile device groups", err)
	}

	var out ResponseSmartMobileDeviceGroupsListV1
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResponseSmartMobileDeviceGroupListItemV1
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "smart mobile device group", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetSmartMobileDeviceGroupMembershipByIDV1 retrieves the membership of a Smart Mobile Device Group by ID
func (c *Client) GetSmartMobileDeviceGroupMembershipByIDV1(id string) (*ResponseSmartMobileDeviceGroupMembershipV1, error) {
	endpoint := fmt.Sprintf("%s/smart-group-membership/%s", uriAPIV1MobileDeviceGroups, id)

	var response ResponseSmartMobileDeviceGroupMembershipV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "smart mobile device group membership", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetSmartMobileDeviceGroupByIDV1 retrieves a specific Smart Mobile Device Group by ID
func (c *Client) GetSmartMobileDeviceGroupByIDV1(id string) (*ResourceSmartMobileDeviceGroupV1, error) {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV1MobileDeviceGroups, id)

	var response ResourceSmartMobileDeviceGroupV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "smart mobile device group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetSmartMobileDeviceGroupByNameV1 retrieves a Smart Mobile Device Group by name
func (c *Client) GetSmartMobileDeviceGroupByNameV1(groupName string) (*ResponseSmartMobileDeviceGroupListItemV1, error) {
	params := url.Values{}
	params.Set("filter", fmt.Sprintf("groupName==\"%s\"", groupName))

	groups, err := c.GetSmartMobileDeviceGroupsV1(params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "smart mobile device groups", err)
	}

	if len(groups.Results) == 0 {
		return nil, fmt.Errorf(errMsgFailedGetByName, "smart mobile device group", groupName, errMsgNoName)
	}

	return &groups.Results[0], nil
}

// CreateSmartMobileDeviceGroupV1 creates a new Smart Mobile Device Group
func (c *Client) CreateSmartMobileDeviceGroupV1(request ResourceSmartMobileDeviceGroupV1) (*ResponseSmartMobileDeviceGroupCreateV1, error) {
	endpoint := fmt.Sprintf("%s/smart-groups", uriAPIV1MobileDeviceGroups)

	var response ResponseSmartMobileDeviceGroupCreateV1
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "smart mobile device group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateSmartMobileDeviceGroupByIDV1 updates an existing Smart Mobile Device Group by ID
func (c *Client) UpdateSmartMobileDeviceGroupByIDV1(id string, request ResourceSmartMobileDeviceGroupV1) (*ResourceSmartMobileDeviceGroupV1, error) {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV1MobileDeviceGroups, id)

	var response ResourceSmartMobileDeviceGroupV1
	resp, err := c.HTTP.DoRequest("PUT", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "smart mobile device group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteSmartMobileDeviceGroupByIDV1 deletes a Smart Mobile Device Group by ID
func (c *Client) DeleteSmartMobileDeviceGroupByIDV1(id string) error {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV1MobileDeviceGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "smart mobile device group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
