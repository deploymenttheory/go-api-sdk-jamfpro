// jamfproapi_api_role_privileges.go
// Jamf Pro Api - API Role Privileges
// API reference: https://developer.jamf.com/jamf-pro/reference/api-role-privileges
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
)

const uriApiRolePrivileges = "/api/v1/api-role-privileges"

// Resource

// ResponseApiRolePrivileges represents the structure of the response for fetching API role privileges
type ResourceApiRolePrivilegesList struct {
	Privileges []string `json:"privileges"`
}

// CRUD

// GetJamfAPIPrivileges fetches a list of Jamf API role privileges
func (c *Client) GetJamfAPIPrivileges() (*ResourceApiRolePrivilegesList, error) {
	var privilegesList ResourceApiRolePrivilegesList
	resp, err := c.HTTP.DoRequest("GET", uriApiRolePrivileges, nil, &privilegesList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "API Privileges", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &privilegesList, nil
}

// GetJamfAPIPrivilegesByName fetches a list of Jamf API role privileges by name
func (c *Client) GetJamfAPIPrivilegesByName(name string, limit int) (*ResourceApiRolePrivilegesList, error) {
	// Encode the name parameter to handle special characters
	encodedName := url.QueryEscape(name)

	// Construct the URL with the provided name and limit
	endpoint := fmt.Sprintf(uriApiRolePrivileges+"/search?name=%s&limit=%d", encodedName, limit)

	var privilegesList ResourceApiRolePrivilegesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &privilegesList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "API Privilege", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &privilegesList, nil
}
