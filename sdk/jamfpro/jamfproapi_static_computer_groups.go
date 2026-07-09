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

// ---------------------------------------------------------------------------
// Static Computer Groups v3 (/api/v3/computer-groups/static-groups)
// ---------------------------------------------------------------------------

// List (v3)

// ResponseStaticComputerGroupsListV3 represents the search results for static computer groups (v3).
type ResponseStaticComputerGroupsListV3 struct {
	TotalCount int                                     `json:"totalCount"`
	Results    []ResponseStaticComputerGroupListItemV3 `json:"results"`
}

// ResponseStaticComputerGroupListItemV3 represents a static computer group summary (v3).
type ResponseStaticComputerGroupListItemV3 struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SiteID      string `json:"siteId"`
	Count       int    `json:"count"`
}

// Resource / Response (v3)

// ResponseStaticComputerGroupV3 represents a static computer group returned by GET by ID (v3).
type ResponseStaticComputerGroupV3 struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SiteID      string `json:"siteId"`
}

// Request (v3)

// ResourceStaticComputerGroupAssignmentV3 represents the request/response body for creating or updating a static computer group (v3).
type ResourceStaticComputerGroupAssignmentV3 struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name"`
	Description string   `json:"description,omitempty"`
	SiteID      *string  `json:"siteId,omitempty"`
	Assignments []string `json:"assignments,omitempty"`
}

// ResponseStaticComputerGroupCreateV3 represents the response for creating a static computer group (v3).
type ResponseStaticComputerGroupCreateV3 struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD (v3)

// GetStaticComputerGroupsV3 retrieves the list of static computer groups using the v3 API.
func (c *Client) GetStaticComputerGroupsV3(params url.Values) (*ResponseStaticComputerGroupsListV3, error) {
	endpoint := fmt.Sprintf("%s/static-groups", uriAPIV3ComputerGroups)
	if params != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var out ResponseStaticComputerGroupsListV3
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "static computer groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetStaticComputerGroupByIDV3 retrieves a specific static computer group by ID using the v3 API.
func (c *Client) GetStaticComputerGroupByIDV3(id string) (*ResponseStaticComputerGroupV3, error) {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriAPIV3ComputerGroups, id)

	var out ResponseStaticComputerGroupV3
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "static computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateStaticComputerGroupV3 creates a new static computer group using the v3 API.
func (c *Client) CreateStaticComputerGroupV3(request ResourceStaticComputerGroupAssignmentV3) (*ResponseStaticComputerGroupCreateV3, error) {
	endpoint := fmt.Sprintf("%s/static-groups", uriAPIV3ComputerGroups)
	if request.Assignments == nil {
		request.Assignments = []string{}
	}

	var out ResponseStaticComputerGroupCreateV3
	resp, err := c.HTTP.DoRequest("POST", endpoint, request, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "static computer group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateStaticComputerGroupByIDV3 updates an existing static computer group by ID using the v3 API.
func (c *Client) UpdateStaticComputerGroupByIDV3(id string, request ResourceStaticComputerGroupAssignmentV3) (*ResourceStaticComputerGroupAssignmentV3, error) {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriAPIV3ComputerGroups, id)
	if request.Assignments == nil {
		request.Assignments = []string{}
	}

	var out ResourceStaticComputerGroupAssignmentV3
	resp, err := c.HTTP.DoRequest("PUT", endpoint, request, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "static computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// DeleteStaticComputerGroupByIDV3 deletes a static computer group by ID using the v3 API.
func (c *Client) DeleteStaticComputerGroupByIDV3(id string) error {
	endpoint := fmt.Sprintf("%s/static-groups/%s", uriAPIV3ComputerGroups, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDeleteByID, "static computer group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
