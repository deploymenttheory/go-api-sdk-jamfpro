// userGroups.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
)

const uriUserGroups = "/JSSResource/usergroups"

type ResponseUserGroup struct {
	ID               int                 `json:"id" xml:"id"`
	Name             string              `json:"name" xml:"name"`
	IsSmart          bool                `json:"is_smart" xml:"is_smart"`
	IsNotifyOnChange bool                `json:"is_notify_on_change" xml:"is_notify_on_change"`
	Site             Site                `json:"site,omitempty" xml:"site,omitempty"`
	Criteria         []UserGroupCriteria `json:"criteria,omitempty" xml:"criteria,omitempty"`
	Users            []UserGroupUser     `json:"users,omitempty" xml:"users,omitempty"`
}

type UserGroupCriteria struct {
	Size      int                `json:"size" xml:"size"`
	Criterion UserGroupCriterion `json:"criterion" xml:"criterion"`
}

type UserGroupCriterion struct {
	Name         string `json:"name" xml:"name"`
	Priority     int    `json:"priority,omitempty" xml:"priority,omitempty"`
	AndOr        string `json:"and_or,omitempty" xml:"and_or,omitempty"`
	SearchType   string `json:"search_type,omitempty" xml:"search_type,omitempty"`
	Value        string `json:"value,omitempty" xml:"value,omitempty"`
	OpeningParen bool   `json:"opening_paren,omitempty" xml:"opening_paren,omitempty"`
	ClosingParen bool   `json:"closing_paren,omitempty" xml:"closing_paren,omitempty"`
}

type UserGroupUser struct {
	Size int  `json:"size" xml:"size"`
	User User `json:"user" xml:"user"`
}

type User struct {
	ID           int    `json:"id" xml:"id"`
	Username     string `json:"username,omitempty" xml:"username,omitempty"`
	FullName     string `json:"full_name,omitempty" xml:"full_name,omitempty"`
	PhoneNumber  string `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	EmailAddress string `json:"email_address,omitempty" xml:"email_address,omitempty"`
}

type UserGroupScope struct {
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}

// GetUserGroupByID retrieves the User Group by its ID
func (c *Client) GetUserGroupByID(id int) (*ResponseUserGroup, error) {
	url := fmt.Sprintf("%s/id/%d", uriUserGroups, id)

	var group ResponseUserGroup
	if err := c.DoRequest("GET", url, nil, nil, &group); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &group, nil
}

// GetUserGroupByName retrieves the User Group by its Name
func (c *Client) GetUserGroupByName(name string) (*ResponseUserGroup, error) {
	url := fmt.Sprintf("%s/name/%s", uriUserGroups, name)

	var group ResponseUserGroup
	if err := c.DoRequest("GET", url, nil, nil, &group); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &group, nil
}

// GetUserGroups retrieves all User Groups
func (c *Client) GetUserGroups() ([]ResponseUserGroup, error) {
	url := uriUserGroups

	var groups []ResponseUserGroup
	if err := c.DoRequest("GET", url, nil, nil, &groups); err != nil {
		return nil, fmt.Errorf("failed to fetch all User Groups: %v", err)
	}

	return groups, nil
}

// CreateSmartUserGroup creates a new Smart User Group
func (c *Client) CreateSmartUserGroup(group *ResponseUserGroup) (*ResponseUserGroup, error) {
	url := fmt.Sprintf("%s/id/0", uriUserGroups)

	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if group.Site.ID == 0 {
		group.Site = Site{ID: -1, Name: "None"}
	}

	// Set the IsSmart field to true
	group.IsSmart = true

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"user_group"`
		*ResponseUserGroup
	}{
		ResponseUserGroup: group,
	}

	// Execute the request
	var responseGroup ResponseUserGroup
	if err := c.DoRequest("POST", url, reqBody, nil, &responseGroup); err != nil {
		return nil, fmt.Errorf("failed to create Smart User Group: %v", err)
	}

	return &responseGroup, nil
}

// CreateStaticUserGroup creates a new Static User Group
func (c *Client) CreateStaticUserGroup(group *ResponseUserGroup) (*ResponseUserGroup, error) {
	url := fmt.Sprintf("%s/id/0", uriUserGroups)

	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if group.Site.ID == 0 {
		group.Site = Site{ID: -1, Name: "None"}
	}

	// Set the IsSmart field to false
	group.IsSmart = false

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"user_group"`
		*ResponseUserGroup
	}{
		ResponseUserGroup: group,
	}

	// Execute the request
	var responseGroup ResponseUserGroup
	if err := c.DoRequest("POST", url, reqBody, nil, &responseGroup); err != nil {
		return nil, fmt.Errorf("failed to create Static User Group: %v", err)
	}

	return &responseGroup, nil
}
