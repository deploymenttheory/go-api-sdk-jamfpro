// classicapi_usergroups.go
// Jamf Pro Classic Api - usergroups
// api reference: https://developer.jamf.com/jamf-pro/reference/usergroups
// Jamf Pro Classic Api requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
- SharedSubsetCriteria
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriUserGroups = "/JSSResource/usergroups"

// List

// ResponseUserGroupsList represents the structure for a list of user groups.
type ResponseUserGroupsList struct {
	Size      int                  `xml:"size"`
	UserGroup []UserGroupsListItem `xml:"user_group"`
}

type UserGroupsListItem struct {
	ID               int    `xml:"id"`
	Name             string `xml:"name"`
	IsSmart          bool   `xml:"is_smart"`
	IsNotifyOnChange bool   `xml:"is_notify_on_change"`
}

// Resource

// ResourceUserGroup represents the detailed information of a user group.
type ResourceUserGroup struct {
	ID               int                       `xml:"id"`
	Name             string                    `xml:"name"`
	IsSmart          bool                      `xml:"is_smart"`
	IsNotifyOnChange bool                      `xml:"is_notify_on_change"`
	Site             SharedResourceSite        `xml:"site"`
	Criteria         []SharedSubsetCriteria    `xml:"criteria>criterion"`
	Users            []UserGroupSubsetUserItem `xml:"users>user"`
	UserAdditions    []UserGroupSubsetUserItem `xml:"user_additions>user"`
	UserDeletions    []UserGroupSubsetUserItem `xml:"user_deletions>user"`
}

// Shared

// UserGroupUserItem represents a user of a user group.
type UserGroupSubsetUserItem struct {
	ID           int    `xml:"id"`
	Username     string `xml:"username"`
	FullName     string `xml:"full_name"`
	PhoneNumber  string `xml:"phone_number,omitempty"`
	EmailAddress string `xml:"email_address"`
}

// CRUD

// GetUserGroups retrieves a list of all user groups.
func (c *Client) GetUserGroups() (*ResponseUserGroupsList, error) {
	endpoint := uriUserGroups

	var userGroupsList ResponseUserGroupsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userGroupsList, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "user groups", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userGroupsList, nil
}

// GetUserGroupsByID retrieves the details of a user group by its ID.
func (c *Client) GetUserGroupByID(id int) (*ResourceUserGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriUserGroups, id)

	var userGroupDetail ResourceUserGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userGroupDetail, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "user group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userGroupDetail, nil
}

// GetUserGroupsByName retrieves the details of a user group by its name.
func (c *Client) GetUserGroupByName(name string) (*ResourceUserGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriUserGroups, name)

	var userGroupDetail ResourceUserGroup
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userGroupDetail, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "user group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userGroupDetail, nil
}

// CreateUserGroup creates a new user group.
func (c *Client) CreateUserGroup(userGroup *ResourceUserGroup) (*ResourceUserGroup, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriUserGroups)

	requestBody := struct {
		XMLName xml.Name `xml:"user_group"`
		*ResourceUserGroup
	}{
		ResourceUserGroup: userGroup,
	}

	var createdUserGroup ResourceUserGroup
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdUserGroup, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "user group", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdUserGroup, nil
}

// UpdateUserGroupByID updates an existing user group by its ID.
func (c *Client) UpdateUserGroupByID(id int, userGroup *ResourceUserGroup) (*ResourceUserGroup, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriUserGroups, id)

	requestBody := struct {
		XMLName xml.Name `xml:"user_group"`
		*ResourceUserGroup
	}{
		ResourceUserGroup: userGroup,
	}

	var updatedUserGroup ResourceUserGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedUserGroup, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "user group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedUserGroup, nil
}

// UpdateUserGroupByName updates an existing user group by its name.
func (c *Client) UpdateUserGroupByName(name string, userGroup *ResourceUserGroup) (*ResourceUserGroup, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriUserGroups, name)

	requestBody := struct {
		XMLName xml.Name `xml:"user_group"`
		*ResourceUserGroup
	}{
		ResourceUserGroup: userGroup,
	}

	var updatedUserGroup ResourceUserGroup
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedUserGroup, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "user group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedUserGroup, nil
}

// DeleteUserGroupByID deletes a user group by its ID.
func (c *Client) DeleteUserGroupByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriUserGroups, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "user group", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteUserGroupByName deletes a user group by its name.
func (c *Client) DeleteUserGroupByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriUserGroups, name)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "user group", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
