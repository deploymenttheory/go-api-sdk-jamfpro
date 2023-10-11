// classicapi_usergroups.go
// Jamf Pro Api - usergroups
// api reference: https://developer.jamf.com/jamf-pro/reference/usergroups
// Jamf Pro API requires the structs to support an XML data structure.

package jamfpro

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

// Functions - TODO
