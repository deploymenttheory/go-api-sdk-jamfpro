// classicapi_users.go
// Jamf Pro Classic Api - users
// api reference: https://developer.jamf.com/jamf-pro/reference/users
// Jamf Pro Classic Api requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriUsers = "/JSSResource/users"

// List

type ResponseUsersList struct {
	Size  int             `xml:"size"`
	Users []UsersListItem `xml:"user"`
}

type UsersListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

type ResourceUser struct {
	ID                  int                           `xml:"id"`
	Name                string                        `xml:"name"`
	FullName            string                        `xml:"full_name"`
	Email               string                        `xml:"email"`
	EmailAddress        string                        `xml:"email_address"`
	PhoneNumber         string                        `xml:"phone_number"`
	Position            string                        `xml:"position"`
	EnableCustomPhoto   bool                          `xml:"enable_custom_photo_url"`
	CustomPhotoURL      string                        `xml:"custom_photo_url"`
	LDAPServer          UserSubsetLDAPServer          `xml:"ldap_server"`
	ExtensionAttributes UserSubsetExtensionAttributes `xml:"extension_attributes"`
	Sites               []SharedResourceSite          `xml:"sites>site"`
	Links               UserSubsetLinks               `xml:"links"`
}

// Subsets & Containers

// LDAP

type UserSubsetLDAPServer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Extension Attributes

type UserSubsetExtensionAttributes struct {
	Attributes []UserSubsetExtensionAttribute `xml:"extension_attribute"`
}

type UserSubsetExtensionAttribute struct {
	ID    int    `xml:"id"`
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
}

// Links

type UserSubsetLinks struct {
	Computers         []UserSubsetLinksListItem `xml:"computers>computer"`
	Peripherals       []UserSubsetLinksListItem `xml:"peripherals>peripheral"`
	MobileDevices     []UserSubsetLinksListItem `xml:"mobile_devices>mobile_device"`
	VPPAssignments    []UserSubsetLinksListItem `xml:"vpp_assignments>vpp_assignment"`
	TotalVPPCodeCount int                       `xml:"total_vpp_code_count"`
}

type UserSubsetLinksListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// CRUD

// GetUsers retrieves a list of all users.
func (c *Client) GetUsers() (*ResponseUsersList, error) {
	endpoint := uriUsers

	var usersList ResponseUsersList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &usersList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "users", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &usersList, nil
}

// GetUserByID retrieves the details of a user by their ID.
func (c *Client) GetUserByID(id int) (*ResourceUser, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriUsers, id)

	var userDetail ResourceUser
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userDetail)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "user", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userDetail, nil
}

// GetUserByName retrieves the details of a user by their name.
func (c *Client) GetUserByName(name string) (*ResourceUser, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriUsers, name)

	var userDetail ResourceUser
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userDetail)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "user", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userDetail, nil
}

// GetUserByEmail retrieves the details of a user by their email.
func (c *Client) GetUserByEmail(email string) (*ResourceUser, error) {
	endpoint := fmt.Sprintf("%s/email/%s", uriUsers, email)

	var userDetail ResourceUser
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userDetail)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByEmail, "user", email, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userDetail, nil
}

// CreateUser creates a new user.
func (c *Client) CreateUser(newUser *ResourceUser) (*ResourceUser, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriUsers)

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*ResourceUser
	}{
		ResourceUser: newUser,
	}

	var createdUser ResourceUser
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdUser)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "user", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdUser, nil
}

// UpdateUserByID updates a user's details by their ID.
func (c *Client) UpdateUserByID(id int, updatedUser *ResourceUser) (*ResourceUser, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriUsers, id)

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*ResourceUser
	}{
		ResourceUser: updatedUser,
	}

	var user ResourceUser
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &user)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "user", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &user, nil
}

// UpdateUserByName updates a user's details by their name.
func (c *Client) UpdateUserByName(name string, updatedUser *ResourceUser) (*ResourceUser, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriUsers, name)

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*ResourceUser
	}{
		ResourceUser: updatedUser,
	}

	var user ResourceUser
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &user)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "user", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &user, nil
}

// UpdateUserByEmail updates a user's details by their email.
func (c *Client) UpdateUserByEmail(email string, updatedUser *ResourceUser) (*ResourceUser, error) {
	endpoint := fmt.Sprintf("%s/email/%s", uriUsers, email)

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*ResourceUser
	}{
		ResourceUser: updatedUser,
	}

	var user ResourceUser
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &user)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByEmail, "user", email, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &user, nil
}

// DeleteUserByID deletes a user by their ID.
func (c *Client) DeleteUserByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriUsers, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "user", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteUserByName deletes a user by their name.
func (c *Client) DeleteUserByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriUsers, name)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "user", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteUserByEmail deletes a user by their email.
func (c *Client) DeleteUserByEmail(email string) error {
	endpoint := fmt.Sprintf("%s/email/%s", uriUsers, email)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByEmail, "user", email, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
