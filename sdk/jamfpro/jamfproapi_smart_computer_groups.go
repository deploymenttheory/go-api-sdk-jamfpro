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
	uriAPIV3ComputerGroups = "/api/v3/computer-groups"
)

// Request

// ResourceSmartComputerGroupV2 represents the request structure for creating a Smart Computer Group
type ResourceSmartComputerGroupV2 struct {
	Name        string                           `json:"name"`
	Description string                           `json:"description,omitempty"`
	Criteria    []SharedSubsetCriteriaJamfProAPI `json:"criteria,omitempty"`
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

// ---------------------------------------------------------------------------
// Smart Computer Groups v3 (/api/v3/computer-groups/smart-groups)
// ---------------------------------------------------------------------------

// List (v3)

// ResponseSmartComputerGroupsListV3 represents the search results for smart computer groups (v3).
type ResponseSmartComputerGroupsListV3 struct {
	TotalCount int                                    `json:"totalCount"`
	Results    []ResponseSmartComputerGroupListItemV3 `json:"results"`
}

// ResponseSmartComputerGroupListItemV3 represents a smart computer group summary (v3).
type ResponseSmartComputerGroupListItemV3 struct {
	ID              string `json:"id"`
	SiteID          string `json:"siteId"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	MembershipCount int    `json:"membershipCount"`
}

// Membership (v3)

// ResponseSmartComputerGroupMembershipV3 represents the membership of a smart computer group (v3).
type ResponseSmartComputerGroupMembershipV3 struct {
	Members []int `json:"members"`
}

// Request / Response (v3)

// ResourceSmartComputerGroupV3 represents the request/response body for creating or updating a smart computer group (v3).
type ResourceSmartComputerGroupV3 struct {
	Name        string                               `json:"name"`
	Description string                               `json:"description,omitempty"`
	Criteria    []SubsetComputerSmartGroupCriteriaV3 `json:"criteria,omitempty"`
	SiteID      *string                              `json:"siteId,omitempty"`
}

// SubsetComputerSmartGroupCriteriaV3 represents a single smart computer group criterion (v3).
type SubsetComputerSmartGroupCriteriaV3 struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen bool   `json:"openingParen,omitempty"`
	ClosingParen bool   `json:"closingParen,omitempty"`
}

// ResponseSmartComputerGroupCreateV3 represents the response for creating a smart computer group (v3).
type ResponseSmartComputerGroupCreateV3 struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD (v3)

// GetSmartComputerGroupsV3 retrieves the list of smart computer groups using the v3 API.
func (c *Client) GetSmartComputerGroupsV3(params url.Values) (*ResponseSmartComputerGroupsListV3, error) {
	endpoint := fmt.Sprintf("%s/smart-groups", uriAPIV3ComputerGroups)
	if params != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var out ResponseSmartComputerGroupsListV3
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "smart computer groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetSmartComputerGroupMembershipByIDV3 retrieves the membership of a smart computer group by ID using the v3 API.
func (c *Client) GetSmartComputerGroupMembershipByIDV3(id string) (*ResponseSmartComputerGroupMembershipV3, error) {
	endpoint := fmt.Sprintf("%s/smart-group-membership/%s", uriAPIV3ComputerGroups, id)

	var out ResponseSmartComputerGroupMembershipV3
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "smart computer group membership", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetSmartComputerGroupByIDV3 retrieves a specific smart computer group by ID using the v3 API.
func (c *Client) GetSmartComputerGroupByIDV3(id string) (*ResourceSmartComputerGroupV3, error) {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV3ComputerGroups, id)

	var out ResourceSmartComputerGroupV3
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "smart computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateSmartComputerGroupV3 creates a new smart computer group using the v3 API.
func (c *Client) CreateSmartComputerGroupV3(request ResourceSmartComputerGroupV3) (*ResponseSmartComputerGroupCreateV3, error) {
	endpoint := fmt.Sprintf("%s/smart-groups", uriAPIV3ComputerGroups)

	var out ResponseSmartComputerGroupCreateV3
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "smart computer group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateSmartComputerGroupByIDV3 updates an existing smart computer group by ID using the v3 API.
func (c *Client) UpdateSmartComputerGroupByIDV3(id string, request ResourceSmartComputerGroupV3) (*ResourceSmartComputerGroupV3, error) {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV3ComputerGroups, id)

	var out ResourceSmartComputerGroupV3
	resp, err := c.HTTP.DoRequest("PUT", endpoint, request, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "smart computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// DeleteSmartComputerGroupByIDV3 deletes a smart computer group by ID using the v3 API.
func (c *Client) DeleteSmartComputerGroupByIDV3(id string) error {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV3ComputerGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDeleteByID, "smart computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
