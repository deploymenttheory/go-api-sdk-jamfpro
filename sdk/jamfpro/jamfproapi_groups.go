// jamfproapi_groups.go
// Jamf Pro Api - Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-groups
// Jamf Pro Api requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriGroups = "/api/v1/groups"
const uriGroupsV2 = "/api/v2/groups"

// List

// Struct for paginated response for groups
type ResponseGroupsList struct {
	TotalCount int             `json:"totalCount"`
	Results    []ResourceGroup `json:"results"`
}

// Resource

// Struct which represents Group object JSON from Pro API
type ResourceGroup struct {
	GroupPlatformId  string `json:"groupPlatformId,omitempty"`
	GroupJamfProId   string `json:"groupJamfProId,omitempty"`
	GroupName        string `json:"groupName,omitempty"`
	GroupDescription string `json:"groupDescription,omitempty"`
	GroupType        string `json:"groupType,omitempty"`
	Smart            bool   `json:"smart"`
	MembershipCount  int    `json:"membershipCount"`
}

// CRUD

// Gets full list of groups & handles pagination
func (c *Client) GetGroups(params url.Values) (*ResponseGroupsList, error) {
	resp, err := c.DoPaginatedGet(uriGroups, params)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "groups", err)
	}

	var out ResponseGroupsList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceGroup
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "group", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// Retrieves group from provided ID & returns ResourceGroup
func (c *Client) GetGroupByID(id string) (*ResourceGroup, error) {
	endpoint := fmt.Sprintf("%s/%s", uriGroups, id)
	var group ResourceGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &group)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &group, nil
}

// Retrieves computer group by JamfProName (groupName) and groupType COMPUTER
func (c *Client) GetComputerGroupByJamfProName(name string) (*ResourceGroup, error) {
	params := url.Values{}
	params.Set("filter", fmt.Sprintf("groupName==\"%s\"", name))
	groups, err := c.GetGroups(params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "groups", err)
	}

	for _, value := range groups.Results {
		if value.GroupType == "COMPUTER" {
			return &value, nil
		}
	}

	return nil, fmt.Errorf("failed to get computer group by name: %s, error: resource with name does not exist", name)
}

// Retrieves mobile group by JamfProName (groupName) and groupType MOBILE
func (c *Client) GetMobileGroupByJamfProName(name string) (*ResourceGroup, error) {
	params := url.Values{}
	params.Set("filter", fmt.Sprintf("groupName==\"%s\"", name))
	groups, err := c.GetGroups(params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "groups", err)
	}

	for _, value := range groups.Results {
		if value.GroupType == "MOBILE" {
			return &value, nil
		}
	}

	return nil, fmt.Errorf("failed to get mobile group by name: %s, error: resource with name does not exist", name)
}

// Retrieves computer group by JamfProID (groupJamfProId) and groupType COMPUTER
func (c *Client) GetComputerGroupByJamfProID(id string) (*ResourceGroup, error) {
	groups, err := c.GetGroups(nil)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "groups", err)
	}

	for _, value := range groups.Results {
		if value.GroupJamfProId == id && value.GroupType == "COMPUTER" {
			return &value, nil
		}
	}

	return nil, fmt.Errorf("failed to get computer group by id: %s, error: resource with id does not exist", id)
}

// Retrieves mobile group by JamfProID (groupJamfProId) and groupType MOBILE
func (c *Client) GetMobileGroupByJamfProID(id string) (*ResourceGroup, error) {
	groups, err := c.GetGroups(nil)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "groups", err)
	}

	for _, value := range groups.Results {
		if value.GroupJamfProId == id && value.GroupType == "MOBILE" {
			return &value, nil
		}
	}

	return nil, fmt.Errorf("failed to get mobile group by id: %s, error: resource with id does not exist", id)
}

// ---------------------------------------------------------------------------
// Unified Groups v2 (/api/v2/groups)
// ---------------------------------------------------------------------------

// List (v2)

// ResponseGroupSearchResultV2 represents the search results for unified groups v2.
type ResponseGroupSearchResultV2 struct {
	TotalCount int             `json:"totalCount"`
	Results    []ResourceGroup `json:"results"`
}

// Resource (v2)

// ResourceGroupWithCriteriaV1 represents a unified group with smart group criteria.
type ResourceGroupWithCriteriaV1 struct {
	GroupPlatformId  string                     `json:"groupPlatformId,omitempty"`
	GroupJamfProId   string                     `json:"groupJamfProId,omitempty"`
	GroupName        string                     `json:"groupName,omitempty"`
	GroupDescription string                     `json:"groupDescription,omitempty"`
	GroupType        string                     `json:"groupType,omitempty"`
	Smart            bool                       `json:"smart"`
	MembershipCount  int                        `json:"membershipCount"`
	Criteria         []SubsetSmartGroupCriteria `json:"criteria,omitempty"`
}

// SubsetSmartGroupCriteria represents a single smart group criterion (v1 format).
type SubsetSmartGroupCriteria struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen bool   `json:"openingParen,omitempty"`
	ClosingParen bool   `json:"closingParen,omitempty"`
}

// Request (v2)

// RequestGroupUpdateV2 represents the request body for updating a unified group v2.
type RequestGroupUpdateV2 struct {
	GroupName        string                              `json:"groupName,omitempty"`
	GroupDescription string                              `json:"groupDescription,omitempty"`
	Criteria         []SubsetUnifiedSmartGroupCriteriaV2 `json:"criteria,omitempty"`
	Assignments      []SubsetGroupAssignmentV1           `json:"assignments,omitempty"`
}

// SubsetUnifiedSmartGroupCriteriaV2 represents a single unified smart group criterion (v2 format).
type SubsetUnifiedSmartGroupCriteriaV2 struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen bool   `json:"openingParen,omitempty"`
	ClosingParen bool   `json:"closingParen,omitempty"`
}

// SubsetGroupAssignmentV1 represents a device assignment for a unified group v2 update.
type SubsetGroupAssignmentV1 struct {
	DeviceID string `json:"deviceId"`
	Selected bool   `json:"selected"`
}

// CRUD (v2)

// GetGroupsV2 retrieves the list of unified groups using the v2 API.
func (c *Client) GetGroupsV2(params url.Values) (*ResponseGroupSearchResultV2, error) {
	endpoint := uriGroupsV2
	if params != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var out ResponseGroupSearchResultV2
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetGroupByIDV2 retrieves a single unified group with criteria by ID using the v2 API.
func (c *Client) GetGroupByIDV2(id string) (*ResourceGroupWithCriteriaV1, error) {
	endpoint := fmt.Sprintf("%s/%s", uriGroupsV2, id)

	var out ResourceGroupWithCriteriaV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateGroupByIDV2 updates a unified group by ID using the v2 API. Returns no body on success (204).
func (c *Client) UpdateGroupByIDV2(id string, request RequestGroupUpdateV2) error {
	endpoint := fmt.Sprintf("%s/%s", uriGroupsV2, id)

	resp, err := c.HTTP.DoRequest("PATCH", endpoint, request, nil)
	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedUpdateByID, "group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteGroupByIDV2 deletes a unified group by ID using the v2 API.
func (c *Client) DeleteGroupByIDV2(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriGroupsV2, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDeleteByID, "group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
