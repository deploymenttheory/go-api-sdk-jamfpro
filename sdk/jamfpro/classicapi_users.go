// classicapi_users.go
// Jamf Pro Api - users
// api reference: https://developer.jamf.com/jamf-pro/reference/users
// Jamf Pro API requires the structs to support an XML data structure.

package jamfpro

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

// Functions TODO
