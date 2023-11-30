// classicapi_users.go
// Jamf Pro Classic Api - users
// api reference: https://developer.jamf.com/jamf-pro/reference/users
// Jamf Pro Classic Api requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriUsers = "/JSSResource/users"

type ResponseUsersList struct {
	Size  int        `xml:"size"`
	Users []UserItem `xml:"user"`
}

type UserItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type ResponseUser struct {
	ID                  int                               `xml:"id"`
	Name                string                            `xml:"name"`
	FullName            string                            `xml:"full_name"`
	Email               string                            `xml:"email"`
	EmailAddress        string                            `xml:"email_address"`
	PhoneNumber         string                            `xml:"phone_number"`
	Position            string                            `xml:"position"`
	EnableCustomPhoto   bool                              `xml:"enable_custom_photo_url"`
	CustomPhotoURL      string                            `xml:"custom_photo_url"`
	LDAPServer          UserDataSubsetLDAPServer          `xml:"ldap_server"`
	ExtensionAttributes UserDataSubsetExtensionAttributes `xml:"extension_attributes"`
	Sites               []UserDataSubsetSite              `xml:"sites>site"`
	Links               UserDataSubsetUserLinks           `xml:"links"`
}

type UserDataSubsetLDAPServer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type UserDataSubsetExtensionAttributes struct {
	Attributes []UserDataSubsetExtensionAttributeItem `xml:"extension_attribute"`
}

type UserDataSubsetExtensionAttributeItem struct {
	ID    int    `xml:"id"`
	Name  string `xml:"name"`
	Type  string `xml:"type"`
	Value string `xml:"value"`
}

type UserDataSubsetSite struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type UserDataSubsetUserLinks struct {
	Computers         []UserDataSubsetComputer      `xml:"computers>computer"`
	Peripherals       []UserDataSubsetPeripheral    `xml:"peripherals>peripheral"`
	MobileDevices     []UserDataSubsetMobileDevice  `xml:"mobile_devices>mobile_device"`
	VPPAssignments    []UserDataSubsetVPPAssignment `xml:"vpp_assignments>vpp_assignment"`
	TotalVPPCodeCount int                           `xml:"total_vpp_code_count"`
}

type UserDataSubsetComputer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type UserDataSubsetPeripheral struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type UserDataSubsetMobileDevice struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type UserDataSubsetVPPAssignment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// GetUsers retrieves a list of all users.
func (c *Client) GetUsers() (*ResponseUsersList, error) {
	endpoint := uriUsers

	var usersList ResponseUsersList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &usersList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &usersList, nil
}

// GetUserByID retrieves the details of a user by their ID.
func (c *Client) GetUserByID(id int) (*ResponseUser, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriUsers, id)

	var userDetail ResponseUser
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userDetail, nil
}

// GetUserByName retrieves the details of a user by their name.
func (c *Client) GetUserByName(name string) (*ResponseUser, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriUsers, name)

	var userDetail ResponseUser
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userDetail, nil
}

// GetUserByEmail retrieves the details of a user by their email.
func (c *Client) GetUserByEmail(email string) (*ResponseUser, error) {
	endpoint := fmt.Sprintf("%s/email/%s", uriUsers, email)

	var userDetail ResponseUser
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &userDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user by email: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &userDetail, nil
}

// CreateUser creates a new user.
func (c *Client) CreateUser(newUser *ResponseUser) (*ResponseUser, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriUsers) // Using ID 0 for creation

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*ResponseUser
	}{
		ResponseUser: newUser,
	}

	var createdUser ResponseUser
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdUser)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdUser, nil
}

// UpdateUserByID updates a user's details by their ID.
func (c *Client) UpdateUserByID(id int, updatedUser *ResponseUser) (*ResponseUser, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriUsers, id)

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*ResponseUser
	}{
		ResponseUser: updatedUser,
	}

	var user ResponseUser
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &user, nil
}

// UpdateUserByName updates a user's details by their name.
func (c *Client) UpdateUserByName(name string, updatedUser *ResponseUser) (*ResponseUser, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriUsers, name)

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*ResponseUser
	}{
		ResponseUser: updatedUser,
	}

	var user ResponseUser
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &user, nil
}

// UpdateUserByEmail updates a user's details by their email.
func (c *Client) UpdateUserByEmail(email string, updatedUser *ResponseUser) (*ResponseUser, error) {
	endpoint := fmt.Sprintf("%s/email/%s", uriUsers, email)

	requestBody := struct {
		XMLName xml.Name `xml:"user"`
		*ResponseUser
	}{
		ResponseUser: updatedUser,
	}

	var user ResponseUser
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user by email: %v", err)
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
		return fmt.Errorf("failed to delete user by ID: %v", err)
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
		return fmt.Errorf("failed to delete user by name: %v", err)
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
		return fmt.Errorf("failed to delete user by email: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
