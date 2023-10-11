// users.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
)

const uriUsers = "/JSSResource/users"

type UserScope struct {
	Name string `xml:"name"`
}

type JamfUserScope struct {
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}

type ResponseUser struct {
	ID                   int                                 `json:"id" xml:"id"`
	Name                 string                              `json:"name" xml:"name"` // required
	FullName             string                              `json:"full_name,omitempty" xml:"full_name,omitempty"`
	Email                string                              `json:"email,omitempty" xml:"email,omitempty"`
	EmailAddress         string                              `json:"email_address,omitempty" xml:"email_address,omitempty"`
	PhoneNumber          string                              `json:"phone_number,omitempty" xml:"phone_number,omitempty"`
	Position             string                              `json:"position,omitempty" xml:"position,omitempty"`
	EnableCustomPhotoURL bool                                `json:"enable_custom_photo_url,omitempty" xml:"enable_custom_photo_url,omitempty"`
	CustomPhotoURL       string                              `json:"custom_photo_url,omitempty" xml:"custom_photo_url,omitempty"`
	LDAPServer           UsersDataSubsetLDAPServer           `json:"ldap_server,omitempty" xml:"ldap_server,omitempty"`
	ExtensionAttributes  []UsersDataSubsetExtensionAttribute `json:"extension_attributes,omitempty" xml:"extension_attributes,omitempty"`
	Sites                []UsersDataSubsetUserSite           `json:"sites,omitempty" xml:"sites,omitempty"`
	Links                UsersDataSubsetUserLinks            `json:"links,omitempty" xml:"links,omitempty"`
}

type UsersDataSubsetLDAPServer struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type UsersDataSubsetExtensionAttribute struct {
	ExtensionAttributeItem `json:"extension_attribute" xml:"extension_attribute"`
}

type ExtensionAttributeItem struct {
	ID    int    `json:"id,omitempty" xml:"id,omitempty"`
	Name  string `json:"name,omitempty" xml:"name,omitempty"`
	Type  string `json:"type,omitempty" xml:"type,omitempty"` // possible values: String, Integer, Date
	Value string `json:"value,omitempty" xml:"value,omitempty"`
}

type UsersDataSubsetUserSite struct {
	Site Site `json:"site,omitempty" xml:"site,omitempty"`
}

type UsersDataSubsetUserLinks struct {
	Computers         ComputerLink      `json:"computers,omitempty" xml:"computers,omitempty"`
	Peripherals       PeripheralLink    `json:"peripherals,omitempty" xml:"peripherals,omitempty"`
	MobileDevices     MobileDeviceLink  `json:"mobile_devices,omitempty" xml:"mobile_devices,omitempty"`
	VPPAssignments    VPPAssignmentLink `json:"vpp_assignments,omitempty" xml:"vpp_assignments,omitempty"`
	TotalVPPCodeCount int               `json:"total_vpp_code_count,omitempty" xml:"total_vpp_code_count,omitempty"`
}

type ComputerLink struct {
	Computer UsersDataSubsetItem `json:"computer" xml:"computer"`
}

type PeripheralLink struct {
	Peripheral UsersDataSubsetItem `json:"peripheral" xml:"peripheral"`
}

type MobileDeviceLink struct {
	MobileDevice UsersDataSubsetItem `json:"mobile_device" xml:"mobile_device"`
}

type VPPAssignmentLink struct {
	VPPAssignment UsersDataSubsetItem `json:"vpp_assignment" xml:"vpp_assignment"`
}

type UsersDataSubsetItem struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type ResponseUsers struct {
	Size  int            `xml:"size"`
	Users []ResponseUser `xml:"user"`
}

type ResponseUserList struct {
	Size  int            `json:"size" xml:"size"`
	Users []UserListItem `json:"user" xml:"user"`
}

type UserListItem struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// Functions

// GetUserByID retrieves the User by its ID
func (c *Client) GetUserByID(id int) (*ResponseUser, error) {
	url := fmt.Sprintf("%s/id/%d", uriUsers, id)

	var user ResponseUser
	if err := c.DoRequest("GET", url, nil, nil, &user); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &user, nil
}

// GetUserByName retrieves the User by its Name
func (c *Client) GetUserByName(name string) (*ResponseUser, error) {
	url := fmt.Sprintf("%s/name/%s", uriUsers, name)

	var user ResponseUser
	if err := c.DoRequest("GET", url, nil, nil, &user); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &user, nil
}

// GetUserByEmail retrieves the User by its email address
func (c *Client) GetUserByEmail(email string) (*ResponseUsers, error) {
	url := fmt.Sprintf("%s/email/%s", uriUsers, email)

	var users ResponseUsers
	if err := c.DoRequest("GET", url, nil, nil, &users); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &users, nil
}

// GetUsers retrieves all Users
func (c *Client) GetUsers() (*ResponseUserList, error) {
	url := uriUsers

	var users ResponseUserList
	if err := c.DoRequest("GET", url, nil, nil, &users); err != nil {
		return nil, fmt.Errorf("failed to fetch all Users: %v", err)
	}

	return &users, nil
}

// CreateUser creates a new User
func (c *Client) CreateUser(user *ResponseUser) (*ResponseUser, error) {
	url := fmt.Sprintf("%s/id/0", uriUsers)

	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if len(user.Sites) == 0 || user.Sites[0].Site.ID == 0 {
		user.Sites = []UsersDataSubsetUserSite{{Site: Site{ID: -1, Name: "None"}}}
	}
	// If LDAPServer ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if user.LDAPServer.ID == 0 {
		user.LDAPServer.ID = -1
	}
	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"user"`
		*ResponseUser
	}{
		ResponseUser: user,
	}

	// Execute the request
	var responseUser ResponseUser
	if err := c.DoRequest("POST", url, reqBody, nil, &responseUser); err != nil {
		return nil, fmt.Errorf("failed to create User: %v", err)
	}

	return &responseUser, nil
}

// UpdateUserById updates an existing User by its ID
func (c *Client) UpdateUserById(id int, user *ResponseUser) (*ResponseUser, error) {
	url := fmt.Sprintf("%s/id/%d", uriUsers, id)
	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if len(user.Sites) == 0 || user.Sites[0].Site.ID == 0 {
		user.Sites = []UsersDataSubsetUserSite{{Site: Site{ID: -1, Name: "None"}}}
	}
	// If LDAPServer ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if user.LDAPServer.ID == 0 {
		user.LDAPServer.ID = -1
	}
	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"user"`
		*ResponseUser
	}{
		ResponseUser: user,
	}

	// Execute the request
	var responseUser ResponseUser
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseUser); err != nil {
		return nil, fmt.Errorf("failed to update User: %v", err)
	}

	return &responseUser, nil
}

// UpdateUserByName updates an existing User by its Name
func (c *Client) UpdateUserByName(name string, user *ResponseUser) (*ResponseUser, error) {
	url := fmt.Sprintf("%s/name/%s", uriUsers, name)
	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if len(user.Sites) == 0 || user.Sites[0].Site.ID == 0 {
		user.Sites = []UsersDataSubsetUserSite{{Site: Site{ID: -1, Name: "None"}}}
	}
	// If LDAPServer ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if user.LDAPServer.ID == 0 {
		user.LDAPServer.ID = -1
	}
	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"user"`
		*ResponseUser
	}{
		ResponseUser: user,
	}

	// Execute the request
	var responseUser ResponseUser
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseUser); err != nil {
		return nil, fmt.Errorf("failed to update User by name: %v", err)
	}

	return &responseUser, nil
}

// UpdateUserByEmail updates an existing User by its Email
func (c *Client) UpdateUserByEmail(email string, user *ResponseUser) (*ResponseUser, error) {
	// Construct the URL for the request using the provided email
	url := fmt.Sprintf("%s/email/%s", uriUsers, email)

	// If Site ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if len(user.Sites) == 0 || user.Sites[0].Site.ID == 0 {
		user.Sites = []UsersDataSubsetUserSite{{Site: Site{ID: -1, Name: "None"}}}
	}
	// If LDAPServer ID is not set (or set to 0), set it to -1 to exclude from the request and set to unused
	if user.LDAPServer.ID == 0 {
		user.LDAPServer.ID = -1
	}
	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"user"`
		*ResponseUser
	}{
		ResponseUser: user,
	}

	// Execute the request
	var responseUser ResponseUser
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseUser); err != nil {
		return nil, fmt.Errorf("failed to update User by email: %v", err)
	}

	return &responseUser, nil
}

// DeleteUserById deletes an existing User by its ID
func (c *Client) DeleteUserById(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriUsers, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete User: %v", err)
	}

	return nil
}

// DeleteUserByName deletes an existing User by its Name
func (c *Client) DeleteUserByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriUsers, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete User by name: %v", err)
	}

	return nil
}
