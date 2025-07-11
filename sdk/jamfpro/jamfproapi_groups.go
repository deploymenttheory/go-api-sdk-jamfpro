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

// List

// Struct for paginated response for groups
type ResponseGroupsList struct {
	TotalCount int             `json:"totalCount"`
	Results    []ResourceGroup `json:"results"`
}

// Resource

// Struct which represents Group object JSON from Pro API
type ResourceGroup struct {
	GroupPlatformId string `json:"groupPlatformId,omitempty"`
	GroupJamfProId  string `json:"groupJamfProId,omitempty"`
	GroupName       string `json:"groupName,omitempty"`
	GroupType       string `json:"groupType,omitempty"`
	Smart           bool   `json:"smart"`
	MembershipCount int    `json:"membershipCount"`
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
