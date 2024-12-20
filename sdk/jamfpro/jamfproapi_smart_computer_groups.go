// jamfproapi_smart_computer_groups.go
// Jamf Pro Api - Smart Computer Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/computergroups
// Jamf Pro API requires the structs to support an JSON data structure.

/*
Shared data structure resources in this endpoint:
- SharedSubsetCriteriaJamfProAPI
*/

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const (
	uriAPISmartComputerGroups   = "/api/v1/computer-groups"
	uriAPIV2SmartComputerGroups = "/api/v2/computer-groups"
)

// Request

// ResourceSmartComputerGroup represents the request structure for creating a Smart Computer Group
type ResourceSmartComputerGroup struct {
	Name     string                           `json:"name"`
	Criteria []SharedSubsetCriteriaJamfProAPI `json:"criteria"`
	SiteId   *string                          `json:"siteId,omitempty"`
}

// Response

// ResponseSmartComputerGroupsList represents the response for list of Smart Computer Groups
type ResponseSmartComputerGroupsList []ResponseSmartComputerGroupListItem

// ResponseSmartComputerGroupListItem represents individual Smart Computer Group items
type ResponseSmartComputerGroupListItem struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	SmartGroup bool   `json:"smartGroup"`
}

// ResponseSmartComputerGroupMembership represents the membership response for a Smart Computer Group
type ResponseSmartComputerGroupMembership struct {
	Members []int `json:"members"`
}

// ResponseSmartComputerGroupsListV2 represents the paginated response for Smart Computer Groups v2
type ResponseSmartComputerGroupsListV2 struct {
	TotalCount int                            `json:"totalCount"`
	Results    []ResourceSmartComputerGroupV2 `json:"results"`
}

// ResourceSmartComputerGroupV2 represents a Smart Computer Group in v2 API
type ResourceSmartComputerGroupV2 struct {
	ID              string `json:"id"`
	SiteId          string `json:"siteId"`
	Name            string `json:"name"`
	MembershipCount int    `json:"membershipCount"`
}

// ResponseSmartComputerGroupCreate represents the response structure for creating a Smart Computer Group
type ResponseSmartComputerGroupCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD

// GetSmartComputerGroups retrieves a list of all Smart Computer Groups
func (c *Client) GetSmartComputerGroups() (*ResponseSmartComputerGroupsList, error) {
	endpoint := uriAPISmartComputerGroups

	var response ResponseSmartComputerGroupsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "smart computer groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetSmartComputerGroupMembershipByID retrieves the membership of a Smart Computer Group by ID
func (c *Client) GetSmartComputerGroupMembershipByID(id string) (*ResponseSmartComputerGroupMembership, error) {
	endpoint := fmt.Sprintf("%s/smart-group-membership/%s", uriAPIV2SmartComputerGroups, id)

	var response ResponseSmartComputerGroupMembership
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "smart computer group membership", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetSmartComputerGroupsV2 retrieves a paginated list of all Smart Computer Groups using V2 API
func (c *Client) GetSmartComputerGroupsV2(sort_filter string) (*ResponseSmartComputerGroupsListV2, error) {
	resp, err := c.DoPaginatedGet(
		fmt.Sprintf("%s/smart-groups", uriAPIV2SmartComputerGroups),
		standardPageSize,
		startingPageNumber,
		sort_filter,
	)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "smart computer groups", err)
	}

	var out ResponseSmartComputerGroupsListV2
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceSmartComputerGroupV2
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "smart computer group", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetSmartComputerGroupByID retrieves a specific Smart Computer Group by ID
func (c *Client) GetSmartComputerGroupByID(id string) (*ResourceSmartComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/%s", uriAPISmartComputerGroups, id)

	var response ResourceSmartComputerGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "smart computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetSmartComputerGroupByName retrieves a Smart Computer Group by name
func (c *Client) GetSmartComputerGroupByName(name string) (*ResourceComputerGroup, error) {
	groups, err := c.GetSmartComputerGroups()
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "smart computer groups", err)
	}

	var groupID string
	for _, group := range *groups {
		if group.Name == name {
			groupID = group.ID
			break
		}
	}

	if groupID == "" {
		return nil, fmt.Errorf(errMsgFailedGetByName, "smart computer group", name, errMsgNoName)
	}
	// mixing the pro and classic api endpoint as GetSmartComputerGroupByID is not available in pro api yet. 20/12/2024
	return c.GetComputerGroupByID(groupID)
}

// CreateSmartComputerGroup creates a new Smart Computer Group
func (c *Client) CreateSmartComputerGroup(request ResourceSmartComputerGroup) (*ResponseSmartComputerGroupCreate, error) {
	endpoint := fmt.Sprintf("%s/smart-groups", uriAPIV2SmartComputerGroups)

	var response ResponseSmartComputerGroupCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "smart computer group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateSmartComputerGroupByID updates an existing Smart Computer Group by ID
func (c *Client) UpdateSmartComputerGroupByID(id string, request ResourceSmartComputerGroup) (*ResourceSmartComputerGroup, error) {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV2SmartComputerGroups, id)

	var response ResourceSmartComputerGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "smart computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteSmartComputerGroupByID deletes a Smart Computer Group by ID
func (c *Client) DeleteSmartComputerGroupByID(id string) error {
	endpoint := fmt.Sprintf("%s/smart-groups/%s", uriAPIV2SmartComputerGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "smart computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
