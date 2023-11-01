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

// ResponseApiRolePrivileges represents the structure of the response for fetching API role privileges
type ResponseApiRolePrivileges struct {
	Privileges []string `json:"privileges"`
}

// GetJamfAPIPrivileges fetches a list of Jamf API role privileges
func (c *Client) GetJamfAPIPrivileges() (*ResponseApiRolePrivileges, error) {
	var privilegesList ResponseApiRolePrivileges
	resp, err := c.HTTP.DoRequest("GET", uriApiRolePrivileges, nil, &privilegesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf API role privileges: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &privilegesList, nil
}

// GetJamfAPIPrivilegesByName fetches a list of Jamf API role privileges by name
func (c *Client) GetJamfAPIPrivilegesByName(name string, limit int) (*ResponseApiRolePrivileges, error) {
	// Encode the name parameter to handle special characters
	encodedName := url.QueryEscape(name)

	// Construct the URL with the provided name and limit
	endpoint := fmt.Sprintf(uriApiRolePrivileges+"/search?name=%s&limit=%d", encodedName, limit)

	var privilegesList ResponseApiRolePrivileges
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &privilegesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf API role privileges by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &privilegesList, nil
}
