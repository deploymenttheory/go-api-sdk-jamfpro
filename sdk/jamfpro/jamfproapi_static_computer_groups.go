// jamfproapi_static_computer_groups.go
// Jamf Pro Api - Static Computer Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-groups-static-groups
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

// ResourceStaticComputerGroupV2 represents the request structure for creating a Static Computer Group
type ResourceStaticComputerGroupV2 struct {
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	Assignments []string `json:"assignments"`
	SiteId      *string  `json:"siteId,omitempty"`
}

// Response

// ResponseStaticComputerGroupListItemV2 represents individual Static Computer Group items
type ResponseStaticComputerGroupListItemV2 struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SiteID      string `json:"siteId"`
	Count       int    `json:"count"`
}

// ResponseStaticComputerGroupsListV2 represents the paginated response for Static Computer Groups v2
type ResponseStaticComputerGroupsListV2 struct {
	TotalCount int                                     `json:"totalCount"`
	Results    []ResponseStaticComputerGroupListItemV2 `json:"results"`
}

// ResourceStaticComputerGroupV2 represents a Static Computer Group in v2 API
type ResponseStaticComputerGroupV2 struct {
	ID          string `json:"id"`
	SiteId      string `json:"siteId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ResponseStaticComputerGroupCreateV2 represents the response structure for creating a Static Computer Group
type ResponseStaticComputerGroupCreateV2 struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD

// GetStaticComputerGroupsV2 retrieves a paginated list of all Static Computer Groups using V2 API
func (c *Client) GetStaticComputerGroupsV2(params url.Values) (*ResponseStaticComputerGroupsListV2, error) {
	resp, err := c.DoPaginatedGet(fmt.Sprintf("%s/static-groups", uriAPIV2ComputerGroups), params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "static computer groups", err)
	}

	var out ResponseStaticComputerGroupsListV2
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResponseStaticComputerGroupListItemV2
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "static computer group", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetStaticComputerGroupByIDV2 retrieves a specific Static Computer Group by ID
func (c *Client) GetStaticComputerGroupByIDV2(id string) (*ResponseStaticComputerGroupListItemV2, error) {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriAPIV2ComputerGroups, id)

	var response ResponseStaticComputerGroupListItemV2
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "static computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetStaticComputerGroupByNameV2 retrieves a Static Computer Group by name
func (c *Client) GetStaticComputerGroupByNameV2(name string) (*ResponseStaticComputerGroupListItemV2, error) {
	params := url.Values{}
	params.Set("filter", fmt.Sprintf("name==\"%s\"", name))

	groups, err := c.GetStaticComputerGroupsV2(params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "static computer groups", err)
	}

	if len(groups.Results) == 0 {
		return nil, fmt.Errorf(errMsgFailedGetByName, "static computer group", name, errMsgNoName)
	}

	return &groups.Results[0], nil
}

// CreateStaticComputerGroupV2 creates a new Static Computer Group
func (c *Client) CreateStaticComputerGroupV2(request ResourceStaticComputerGroupV2) (*ResponseStaticComputerGroupCreateV2, error) {
	endpoint := fmt.Sprintf("%s/static-groups", uriAPIV2ComputerGroups)

	var response ResponseStaticComputerGroupCreateV2
	if request.Assignments == nil {
		request.Assignments = []string{}
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "static computer group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateStaticComputerGroupByIDV2 updates an existing Static Computer Group by ID
func (c *Client) UpdateStaticComputerGroupByIDV2(id string, request ResourceStaticComputerGroupV2) (*ResourceStaticComputerGroupV2, error) {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriAPIV2ComputerGroups, id)

	var response ResourceStaticComputerGroupV2
	if request.Assignments == nil {
		request.Assignments = []string{}
	}

	resp, err := c.HTTP.DoRequest("PUT", endpoint, request, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "static computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteStaticComputerGroupByIDV2 deletes a Static Computer Group by ID
func (c *Client) DeleteStaticComputerGroupByIDV2(id string) error {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriAPIV2ComputerGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "static computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
