// classicapi_usergroups.go
// Jamf Pro Classic Api - usergroups
// api reference: https://developer.jamf.com/jamf-pro/reference/usergroups
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriUserGroups = "/JSSResource/usergroups"

// ResponseUserGroupsList represents the structure for a list of user groups.
type ResponseUserGroupsList struct {
	Size      int             `xml:"size"`
	UserGroup []UserGroupItem `xml:"user_group"`
}

// UserGroupItem represents a single user group.
type UserGroupItem struct {
	ID               int    `xml:"id"`
	Name             string `xml:"name"`
	IsSmart          bool   `xml:"is_smart"`
	IsNotifyOnChange bool   `xml:"is_notify_on_change"`
}

// ResponseUserGroup represents the detailed information of a user group.
type ResponseUserGroup struct {
	ID               int                      `xml:"id"`
	Name             string                   `xml:"name"`
	IsSmart          bool                     `xml:"is_smart"`
	IsNotifyOnChange bool                     `xml:"is_notify_on_change"`
	Site             UserGroupSite            `xml:"site"`
	Criteria         []UserGroupCriterionItem `xml:"criteria>criterion"`
	Users            []UserGroupUserItem      `xml:"users>user"`
	UserAdditions    []UserGroupUserItem      `xml:"user_additions>user"`
	UserDeletions    []UserGroupUserItem      `xml:"user_deletions>user"`
}

// UserGroupSite represents the site details of a user group.
type UserGroupSite struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// UserGroupCriterionItem represents a criterion of a user group.
type UserGroupCriterionItem struct {
	Name         string `xml:"name"`
	Priority     int    `xml:"priority"`
	AndOr        string `xml:"and_or"`
	SearchType   string `xml:"search_type"`
	Value        string `xml:"value"`
	OpeningParen bool   `xml:"opening_paren,omitempty"`
	ClosingParen bool   `xml:"closing_paren,omitempty"`
}

// UserGroupUserItem represents a user of a user group.
type UserGroupUserItem struct {
	ID           int    `xml:"id"`
	Username     string `xml:"username"`
	FullName     string `xml:"full_name"`
	PhoneNumber  string `xml:"phone_number,omitempty"`
	EmailAddress string `xml:"email_address"`
}

// GetUserGroups retrieves a list of all user groups.
func (c *Client) GetUserGroups() (*ResponseUserGroupsList, error) {
	endpoint := uriUserGroups

	var userGroupsList ResponseUserGroupsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userGroupsList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user groups: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userGroupsList, nil
}

// GetUserGroupsByID retrieves the details of a user group by its ID.
func (c *Client) GetUserGroupsByID(id int) (*ResponseUserGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriUserGroups, id)

	var userGroupDetail ResponseUserGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userGroupDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userGroupDetail, nil
}

// GetUserGroupsByName retrieves the details of a user group by its name.
func (c *Client) GetUserGroupsByName(name string) (*ResponseUserGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriUserGroups, name)

	var userGroupDetail ResponseUserGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userGroupDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userGroupDetail, nil
}

// CreateUserGroup creates a new user group.
func (c *Client) CreateUserGroup(userGroup *ResponseUserGroup) (*ResponseUserGroup, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriUserGroups) // Using ID 0 for creation

	requestBody := struct {
		XMLName xml.Name `xml:"user_group"`
		*ResponseUserGroup
	}{
		ResponseUserGroup: userGroup,
	}

	var createdUserGroup ResponseUserGroup
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdUserGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to create user group: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdUserGroup, nil
}

// UpdateUserGroupByID updates an existing user group by its ID.
func (c *Client) UpdateUserGroupByID(id int, userGroup *ResponseUserGroup) (*ResponseUserGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriUserGroups, id)

	requestBody := struct {
		XMLName xml.Name `xml:"user_group"`
		*ResponseUserGroup
	}{
		ResponseUserGroup: userGroup,
	}

	var updatedUserGroup ResponseUserGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedUserGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to update user group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedUserGroup, nil
}

// UpdateUserGroupByName updates an existing user group by its name.
func (c *Client) UpdateUserGroupByName(name string, userGroup *ResponseUserGroup) (*ResponseUserGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriUserGroups, name)

	requestBody := struct {
		XMLName xml.Name `xml:"user_group"`
		*ResponseUserGroup
	}{
		ResponseUserGroup: userGroup,
	}

	var updatedUserGroup ResponseUserGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedUserGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to update user group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedUserGroup, nil
}

// DeleteUserGroupByID deletes a user group by its ID.
func (c *Client) DeleteUserGroupByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriUserGroups, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete user group by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteUserGroupByName deletes a user group by its name.
func (c *Client) DeleteUserGroupByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriUserGroups, name)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete user group by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
