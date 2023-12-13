// jamfproapi_api_roles.go
// Jamf Pro Api - API Roles
// api reference: https://developer.jamf.com/jamf-pro/reference/getallapiroles
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriApiRoles = "/api/v1/api-roles"

// ResponseApiRoles represents the structure of the response for fetching API roles
type ResponseApiRolesList struct {
	Size    int               `json:"totalCount"`
	Results []ResourceAPIRole `json:"results"`
}

// Role represents the details of an individual API role
type ResourceAPIRole struct {
	ID          string   `json:"id,omitempty"`
	DisplayName string   `json:"displayName,omitempty"`
	Privileges  []string `json:"privileges,omitempty"`
}

// GetJamfAPIRoles fetches a list of Jamf API roles
func (c *Client) GetJamfAPIRoles() (*ResponseApiRolesList, error) {
	endpoint := uriApiRoles

	resp, err := c.DoPaginatedGet(endpoint, standardPageSize, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf API roles: %v", err)
	}

	var outStruct ResponseApiRolesList
	outStruct.Size = resp.Size
	for _, value := range resp.Results {
		var newObj ResourceAPIRole
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf("failed to map structure, %v", err)
		}
		outStruct.Results = append(outStruct.Results, newObj)
	}

	return &outStruct, nil
}

// GetJamfApiRolesByID fetches a Jamf API role by its ID.
func (c *Client) GetJamfApiRoleByID(id string) (*ResourceAPIRole, error) {
	endpoint := fmt.Sprintf(uriApiRoles+"/%s", id)

	var ApiRole ResourceAPIRole
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ApiRole)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf API role with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ApiRole, nil
}

/////// PROGRESS TO HERE

// GetJamfApiRolesNameById fetches a Jamf API role by its display name and then retrieves its details using its ID.
func (c *Client) GetJamfApiRoleByName(name string) (*ResourceAPIRole, error) {
	roles, err := c.GetJamfAPIRoles()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Jamf API roles: %v", err)
	}

	for _, value := range roles.Results {
		if value.DisplayName == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf("no Jamf API role found with the name %s", name)
}

// CreateJamfApiRole creates a new Jamf API role
func (c *Client) CreateJamfApiRole(role *ResourceAPIRole) (*ResourceAPIRole, error) {
	endpoint := uriApiRoles
	var response ResourceAPIRole

	resp, err := c.HTTP.DoRequest("POST", endpoint, role, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create Jamf API role: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateJamfApiRoleByID updates a Jamf API role by its ID
func (c *Client) UpdateJamfApiRoleByID(id string, role *ResourceAPIRole) (*ResourceAPIRole, error) {
	endpoint := fmt.Sprintf(uriApiRoles+"/%s", id)

	var updatedRole ResourceAPIRole
	resp, err := c.HTTP.DoRequest("PUT", endpoint, role, &updatedRole)
	if err != nil {
		return nil, fmt.Errorf("failed to update Jamf Api Role with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedRole, nil
}

// UpdateJamfApiRoleByName updates a Jamf API role based on its display name
func (c *Client) UpdateJamfApiRoleByName(name string, updatedRole *ResourceAPIRole) (*ResourceAPIRole, error) {
	target, err := c.GetJamfApiRoleByName(name)

	if err != nil {
		return nil, fmt.Errorf("failed to get api role by name, %v", err)
	}

	target_id := target.ID
	resp, err := c.UpdateJamfApiRoleByID(target_id, updatedRole)

	if err != nil {
		return nil, fmt.Errorf("failed to update role by id (by name), %v", err)
	}

	return resp, nil
}

// DeleteJamfApiRoleByID deletes a Jamf API role by its ID
func (c *Client) DeleteJamfApiRoleByID(id string) error {
	endpoint := fmt.Sprintf(uriApiRoles+"/%s", id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Jamf Api Role with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteJamfApiRoleByName deletes a Jamf API role by its display name
func (c *Client) DeleteJamfApiRoleByName(name string) error {
	target, err := c.GetJamfApiRoleByName(name)
	if err != nil {
		return fmt.Errorf("failed to get api role by name, %v", err)
	}

	target_id := target.ID

	err = c.DeleteJamfApiRoleByID(target_id)

	if err != nil {
		return fmt.Errorf("failed to delete api role, %v", err)
	}

	return nil
}
