// jamfproapi_api_roles.go
// Jamf Pro Api - API Roles
// api reference: https://developer.jamf.com/jamf-pro/reference/getallapiroles
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"strconv"
)

const uriApiRoles = "/api/v1/api-roles"

// ResponseApiRoles represents the structure of the response for fetching API roles
type ResponseApiRoles struct {
	TotalCount int       `json:"totalCount"`
	Results    []APIRole `json:"results"`
}

// Role represents the details of an individual API role
type APIRole struct {
	ID          string   `json:"id,omitempty"`
	DisplayName string   `json:"displayName,omitempty"`
	Privileges  []string `json:"privileges,omitempty"`
}

// GetJamfAPIRoles fetches a list of Jamf API roles
func (c *Client) GetJamfAPIRoles() (*ResponseApiRoles, error) {
	var rolesList ResponseApiRoles
	resp, err := c.HTTP.DoRequest("GET", uriApiRoles, nil, &rolesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf API roles: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &rolesList, nil
}

// GetJamfApiRolesByID fetches a Jamf API role by its ID.
func (c *Client) GetJamfApiRolesByID(id int) (*APIRole, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf(uriApiRoles+"/%d", id)

	var profile APIRole
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Jamf API role with ID %d: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetJamfApiRolesNameById fetches a Jamf API role by its display name and then retrieves its details using its ID.
func (c *Client) GetJamfApiRolesNameById(name string) (*APIRole, error) {
	rolesList, err := c.GetJamfAPIRoles()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Jamf API roles: %v", err)
	}

	// Search for the role with the given name
	for _, role := range rolesList.Results {
		fmt.Printf("Comparing desired name '%s' with role name '%s'\n", name, role.DisplayName) // Debug log
		if role.DisplayName == name {
			// Convert the ID from string to int before passing it to GetJamfApiRolesByID
			roleID, convErr := strconv.Atoi(role.ID)
			if convErr != nil {
				return nil, fmt.Errorf("failed to convert role ID '%s' to integer: %v", role.ID, convErr)
			}
			return c.GetJamfApiRolesByID(roleID)
		}
	}

	return nil, fmt.Errorf("no Jamf API role found with the name %s", name)
}

// CreateJamfApiRole creates a new Jamf API role
func (c *Client) CreateJamfApiRole(role *APIRole) (*APIRole, error) {
	endpoint := uriApiRoles

	var response APIRole
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
func (c *Client) UpdateJamfApiRoleByID(id string, roleUpdate *APIRole) (*APIRole, error) {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf(uriApiRoles+"/%s", id)

	var updatedRole APIRole
	resp, err := c.HTTP.DoRequest("PUT", endpoint, roleUpdate, &updatedRole)
	if err != nil {
		return nil, fmt.Errorf("failed to update Jamf Api Role with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedRole, nil
}

// UpdateJamfApiRoleByName updates a Jamf API role based on its display name
func (c *Client) UpdateJamfApiRoleByName(name string, updatedRole *APIRole) (*APIRole, error) {
	rolesList, err := c.GetJamfAPIRoles()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all Jamf API roles: %v", err)
	}

	// Search for the role with the given name
	for _, role := range rolesList.Results {
		if role.DisplayName == name {
			// Update the role with the provided ID
			return c.UpdateJamfApiRoleByID(role.ID, updatedRole)
		}
	}

	return nil, fmt.Errorf("no Jamf API role found with the name %s", name)
}

// DeleteJamfApiRoleByID deletes a Jamf API role by its ID
func (c *Client) DeleteJamfApiRoleByID(id string) error {
	// Construct the URL with the provided ID
	endpoint := fmt.Sprintf(uriApiRoles+"/%s", id)

	// Perform the DELETE request
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
	rolesList, err := c.GetJamfAPIRoles()
	if err != nil {
		return fmt.Errorf("failed to fetch all Jamf API roles: %v", err)
	}

	// Search for the role with the given name
	for _, role := range rolesList.Results {
		if role.DisplayName == name {
			return c.DeleteJamfApiRoleByID(role.ID)
		}
	}

	return fmt.Errorf("no Jamf API role found with the name %s", name)
}

// DeleteSelfServiceBrandingMacOSByName deletes a self-service branding configuration for macOS by name.
func (c *Client) DeleteSelfServiceBrandingMacOSByName(name string) error {
	brandingList, err := c.GetSelfServiceBrandingMacOS()
	if err != nil {
		return fmt.Errorf("failed to fetch all self-service branding configurations: %v", err)
	}

	// Search for the branding with the given name
	for _, branding := range brandingList.Results {
		if branding.BrandingName == name || branding.BrandingNameSecondary == name {
			return c.DeleteSelfServiceBrandingMacOSByID(branding.ID)
		}
	}

	return fmt.Errorf("no self-service branding found with the name %s", name)
}
