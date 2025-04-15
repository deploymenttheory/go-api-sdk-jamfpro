// jamfproapi_api_roles.go
// Jamf Pro Api - API Roles
// api reference: https://developer.jamf.com/jamf-pro/reference/getallapiroles
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriApiRoles = "/api/v1/api-roles"

// List

// ResponseApiRoles represents the structure of the response for fetching API roles
type ResponseApiRolesList struct {
	TotalCount int               `json:"totalCount"`
	Results    []ResourceAPIRole `json:"results"`
}

// Resource

// Role represents the details of an individual API role
type ResourceAPIRole struct {
	ID          string   `json:"id,omitempty"`
	DisplayName string   `json:"displayName,omitempty"`
	Privileges  []string `json:"privileges,omitempty"`
}

// CRUD

// GetJamfAPIRoles fetches a list of Jamf API roles
func (c *Client) GetJamfAPIRoles(params url.Values) (*ResponseApiRolesList, error) {
	endpoint := uriApiRoles

	resp, err := c.DoPaginatedGet(endpoint, standardPageSize, 0, params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "api roles", err)
	}

	var outStruct ResponseApiRolesList
	outStruct.TotalCount = resp.Size
	for _, value := range resp.Results {
		var newObj ResourceAPIRole
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "api role", err)
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
		return nil, fmt.Errorf(errMsgFailedGetByID, "api role", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ApiRole, nil
}

// GetJamfApiRolesNameById fetches a Jamf API role by its display name and then retrieves its details using its ID.
func (c *Client) GetJamfApiRoleByName(name string) (*ResourceAPIRole, error) {
	roles, err := c.GetJamfAPIRoles(url.Values{})
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "api role", err)
	}

	for _, value := range roles.Results {
		if value.DisplayName == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "api role", name, errMsgNoName)
}

// CreateJamfApiRole creates a new Jamf API role
func (c *Client) CreateJamfApiRole(role *ResourceAPIRole) (*ResourceAPIRole, error) {
	endpoint := uriApiRoles
	var response ResourceAPIRole

	resp, err := c.HTTP.DoRequest("POST", endpoint, role, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "api role", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateJamfApiRoleByID updates a Jamf API role by its ID
func (c *Client) UpdateJamfApiRoleByID(id string, roleUpdate *ResourceAPIRole) (*ResourceAPIRole, error) {
	endpoint := fmt.Sprintf(uriApiRoles+"/%s", id)

	var updatedRole ResourceAPIRole
	resp, err := c.HTTP.DoRequest("PUT", endpoint, roleUpdate, &updatedRole)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "api role", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedRole, nil
}

// UpdateJamfApiRoleByName updates a Jamf API role based on its display name
func (c *Client) UpdateJamfApiRoleByName(name string, roleUpdate *ResourceAPIRole) (*ResourceAPIRole, error) {
	target, err := c.GetJamfApiRoleByName(name)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "api role", name, err)
	}

	target_id := target.ID
	resp, err := c.UpdateJamfApiRoleByID(target_id, roleUpdate)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "api role", name, err)
	}

	return resp, nil
}

// DeleteJamfApiRoleByID deletes a Jamf API role by its ID
func (c *Client) DeleteJamfApiRoleByID(id string) error {
	endpoint := fmt.Sprintf(uriApiRoles+"/%s", id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "api role", id, err)
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
		return fmt.Errorf(errMsgFailedGetByName, "api role", name, err)
	}

	target_id := target.ID

	err = c.DeleteJamfApiRoleByID(target_id)

	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "api role", name, err)
	}

	return nil
}
