// jamfproapi_account_groups.go
// Jamf Pro Api - Account Groups
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-account-groups
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
)

const uriAccountGroupsV1 = "/api/v1/account-groups"

// List

// ResponseAccountGroupsListV1 represents the search results for Jamf Pro API account groups.
type ResponseAccountGroupsListV1 struct {
	TotalCount int                      `json:"totalCount"`
	Results    []ResourceAccountGroupV1 `json:"results"`
}

// Resource

// ResourceAccountGroupV1 represents a Jamf Pro API account group.
type ResourceAccountGroupV1 struct {
	ID               string                         `json:"id,omitempty"`
	Name             string                         `json:"name"`
	AccessLevel      string                         `json:"accessLevel,omitempty"`
	PrivilegeLevel   string                         `json:"privilegeLevel,omitempty"`
	SiteID           string                         `json:"siteId,omitempty"`
	LdapServerID     string                         `json:"ldapServerId,omitempty"`
	DirectoryGroupID string                         `json:"directoryGroupId,omitempty"`
	Members          []ResourceAccountGroupMemberV1 `json:"members,omitempty"`
	Privileges       []string                       `json:"privileges,omitempty"`
}

// ResourceAccountGroupMemberV1 represents a member of a Jamf Pro API account group.
type ResourceAccountGroupMemberV1 struct {
	ID       string `json:"id"`
	Username string `json:"username,omitempty"`
	Realname string `json:"realname,omitempty"`
	Email    string `json:"email,omitempty"`
}

// GetAccountGroupsV1 retrieves the list of Jamf Pro API account groups.
func (c *Client) GetAccountGroupsV1(params url.Values) (*ResponseAccountGroupsListV1, error) {
	endpoint := uriAccountGroupsV1
	if params != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var out ResponseAccountGroupsListV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "account groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetAccountGroupByIDV1 retrieves a single Jamf Pro API account group by ID.
func (c *Client) GetAccountGroupByIDV1(id string) (*ResourceAccountGroupV1, error) {
	endpoint := fmt.Sprintf("%s/%s", uriAccountGroupsV1, id)

	var out ResourceAccountGroupV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "account group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
