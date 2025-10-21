// jamfproapi_smart_computer_groups.go
// Jamf Pro Api - Smart Computer Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-smart-groups
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

const (
	uriAPIV2ComputerGroups = "/api/v2/computer-groups"
)

// Request

// ResourceSmartComputerGroupV2 represents the request structure for creating a Smart Computer Group
type ResourceSmartComputerGroupV2 struct {
	Name        string                           `json:"name"`
	Description string                           `json:"description,omitempty"`
	Criteria    []SharedSubsetCriteriaJamfProAPI `json:"criteria"`
	SiteId      *string                          `json:"siteId,omitempty"`
}

// Response

// ResponseSmartComputerGroupListItemV2 represents individual Smart Computer Group items
type ResponseSmartComputerGroupListItemV2 struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	SiteID          string `json:"siteId"`
	MembershipCount int    `json:"membershipCount"`
}

// ResponseSmartComputerGroupMembershipV2 represents the membership response for a Smart Computer Group
type ResponseSmartComputerGroupMembershipV2 struct {
	Members []int `json:"members"`
}

// ResponseSmartComputerGroupsListV2 represents the paginated response for Smart Computer Groups v2
type ResponseSmartComputerGroupsListV2 struct {
	TotalCount int                                    `json:"totalCount"`
	Results    []ResponseSmartComputerGroupListItemV2 `json:"results"`
}

// ResourceSmartComputerGroupV2 represents a Smart Computer Group in v2 API
type ResponseSmartComputerGroupV2 struct {
	ID              string `json:"id"`
	SiteId          string `json:"siteId"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	MembershipCount int    `json:"membershipCount"`
}

// ResponseSmartComputerGroupCreateV2 represents the response structure for creating a Smart Computer Group
type ResponseSmartComputerGroupCreateV2 struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD

// GetSmartComputerGroupsV2 retrieves a paginated list of all Smart Computer Groups using V2 API
func (c *Client) GetSmartComputerGroupsV2(params url.Values) (*ResponseSmartComputerGroupsListV2, error) {
	resp, err := c.DoPaginatedGet(fmt.Sprintf("%s/smart-groups", uriAPIV2ComputerGroups), params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "smart computer groups", err)
	}

	var out ResponseSmartComputerGroupsListV2
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResponseSmartComputerGroupListItemV2
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "smart computer group", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetSmartComputerGroupMembershipByIDV2 retrieves the membership of a Smart Computer Group by ID
func (c *Client) GetSmartComputerGroupMembershipByIDV2(id string) (*ResponseSmartComputerGroupMembershipV2, error) {
	endpoint := fmt.Sprintf("%s/smart-group-membership/%s", uriAPIV2ComputerGroups, id)

	var response ResponseSmartComputerGroupMembershipV2
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "smart computer group membership", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetSmartComputerGroupByIDV2 retrieves a specific Smart Computer Group by ID
func (c *Client) GetSmartComputerGroupByIDV2(id string) (*ResourceSmartComputerGroupV2, error) {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV2ComputerGroups, id)

	var response ResourceSmartComputerGroupV2
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "smart computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetSmartComputerGroupByNameV2 retrieves a Smart Computer Group by name
func (c *Client) GetSmartComputerGroupByNameV2(name string) (*ResponseSmartComputerGroupListItemV2, error) {
	params := url.Values{}
	params.Set("filter", fmt.Sprintf("name==\"%s\"", name))

	groups, err := c.GetSmartComputerGroupsV2(params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "smart computer groups", err)
	}

	if len(groups.Results) == 0 {
		return nil, fmt.Errorf(errMsgFailedGetByName, "smart computer group", name, errMsgNoName)
	}

	return &groups.Results[0], nil
}

// CreateSmartComputerGroupV2 creates a new Smart Computer Group
func (c *Client) CreateSmartComputerGroupV2(request ResourceSmartComputerGroupV2) (*ResponseSmartComputerGroupCreateV2, error) {
	endpoint := fmt.Sprintf("%s/smart-groups", uriAPIV2ComputerGroups)

	var response ResponseSmartComputerGroupCreateV2
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "smart computer group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateSmartComputerGroupByIDV2 updates an existing Smart Computer Group by ID
func (c *Client) UpdateSmartComputerGroupByIDV2(id string, request ResourceSmartComputerGroupV2) (*ResourceSmartComputerGroupV2, error) {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV2ComputerGroups, id)

	var response ResourceSmartComputerGroupV2
	resp, err := c.HTTP.DoRequest("PUT", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "smart computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteSmartComputerGroupByIDV2 deletes a Smart Computer Group by ID
func (c *Client) DeleteSmartComputerGroupByIDV2(id string) error {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV2ComputerGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "smart computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
