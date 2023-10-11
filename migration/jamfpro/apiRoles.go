// apiRoles.go
// Jamf Pro Api
// Jamf Pro Api requires the structs to support JSON.

package jamfpro

import (
	"fmt"
)

const uriAPIRoles = "/api/v1/api-roles"

type ResponseAPIRole struct {
	TotalCount int       `json:"totalCount"`
	Results    []APIRole `json:"results"`
}

type APIRole struct {
	ID          string   `json:"id"`
	DisplayName string   `json:"displayName"`
	Privileges  []string `json:"privileges"`
}

type APIRoleUpdateRequest struct {
	DisplayName string   `json:"displayName"`
	Privileges  []string `json:"privileges"`
}

func (c *Client) GetApiRoleIdByName(name string) (string, error) {
	var id string
	roles, err := c.GetApiRoles()
	if err != nil {
		return "", err
	}

	for _, v := range roles.Results {
		if v.DisplayName == name {
			id = v.ID
			break
		}
	}
	return id, err
}

func (c *Client) GetApiRoleByName(name string) (*APIRole, error) {
	allRolesResponse, err := c.GetApiRoles()
	if err != nil {
		return nil, err
	}

	for _, role := range allRolesResponse.Results {
		if role.DisplayName == name {
			return &role, nil
		}
	}

	return nil, fmt.Errorf("API role with name '%s' not found", name)
}

func (c *Client) GetApiRoles() (*ResponseAPIRole, error) {
	uri := fmt.Sprintf("%s?page=0&page-size=100&sort=id%%3Aasc", uriAPIRoles)

	var out ResponseAPIRole
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get API roles: %v", err)
	}

	return &out, nil
}

func (c *Client) GetApiRoleByID(roleID int) (*APIRole, error) {
	uri := fmt.Sprintf("%s/%d", uriAPIRoles, roleID)

	var out APIRole
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get API role by ID: %v", err)
	}

	return &out, nil
}

func (c *Client) CreateApiRole(displayName *string, privileges *[]string) (*APIRole, error) {

	in := struct {
		DisplayName *string   `json:"displayName"`
		Privileges  *[]string `json:"privileges"`
	}{
		DisplayName: displayName,
		Privileges:  privileges,
	}

	var out *APIRole

	err := c.DoRequest("POST", uriAPIRoles, in, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to create API role: %v", err)
	}
	return out, nil
}

func (c *Client) UpdateApiRole(d *APIRole) (*APIRole, error) {
	uri := fmt.Sprintf("%s/%s", uriAPIRoles, d.ID)

	updateRequest := &APIRoleUpdateRequest{
		DisplayName: d.DisplayName,
		Privileges:  d.Privileges,
	}

	updatedRole := &APIRole{}

	// Perform the PUT request
	err := c.DoRequest("PUT", uri, updateRequest, nil, updatedRole)
	if err != nil {
		return nil, fmt.Errorf("failed to update API role: %v", err)
	}

	return updatedRole, nil
}

func (c *Client) DeleteApiRole(roleID int) error {
	uri := fmt.Sprintf("%s/%d", uriAPIRoles, roleID)

	// Perform the DELETE request
	err := c.DoRequest("DELETE", uri, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete API role: %v", err)
	}
	return nil
}
