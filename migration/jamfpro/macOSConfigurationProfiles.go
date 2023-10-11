// macOSConfigurationProfiles.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
	"strings"
)

const uriMacOSConfigurationProfiles = "/JSSResource/osxconfigurationprofiles"

type ResponseMacOSConfigurationProfile struct {
	General     MacOSConfigurationProfileGeneral     `json:"general" xml:"general"`
	Scope       MacOSConfigurationProfileScope       `json:"scope" xml:"scope"`
	SelfService MacOSConfigurationProfileSelfService `json:"self_service,omitempty" xml:"self_service,omitempty"`
}

type MacOSConfigurationProfileGeneral struct {
	ID                 int             `json:"id,omitempty" xml:"id,omitempty"`
	Name               string          `json:"name" xml:"name"`
	Description        string          `json:"description" xml:"description"`
	Site               Site            `json:"site" xml:"site"`
	Category           GeneralCategory `json:"category,omitempty" xml:"category,omitempty"`
	DistributionMethod string          `json:"distribution_method" xml:"distribution_method"`
	UserRemovable      bool            `json:"user_removable" xml:"user_removable"`
	Level              string          `json:"level" xml:"level"`
	UUID               string          `json:"uuid" xml:"uuid"`
	RedeployOnUpdate   string          `json:"redeploy_on_update" xml:"redeploy_on_update"`
	Payload            string          `json:"payloads" xml:"payloads"`
}

type MacOSConfigurationProfileScope struct {
	AllComputers   bool                                      `json:"all_computers" xml:"all_computers"`
	AllUsers       bool                                      `json:"all_jss_users" xml:"all_jss_users"`
	Computers      []ComputerScope                           `json:"computers,omitempty" xml:"computers>computer,omitempty"`
	Buildings      []BuildingScope                           `json:"buildings,omitempty" xml:"buildings>building,omitempty"`
	Departments    []DepartmentScope                         `json:"departments,omitempty" xml:"departments>department,omitempty"`
	ComputerGroups []ComputerGroupListResponse               `json:"computer_groups,omitempty" xml:"computer_groups>computer_group,omitempty"`
	JamfUsers      []JamfUserScope                           `json:"jss_users,omitempty" xml:"jss_users>jss_user,omitempty"`
	JamfUserGroups []UserGroupScope                          `json:"jss_user_groups,omitempty" xml:"jss_user_groups>jss_user_group,omitempty"`
	Limitiations   MacOSConfigurationProfileScopeLimitations `json:"limitations" xml:"limitations"`
	Exclusions     MacOSConfigurationProfileScopeExclusions  `json:"exclusions" xml:"exclusions"`
}

type MacOSConfigurationProfileScopeLimitations struct {
	Users           []UserScope           `json:"users,omitempty" xml:"users>user,omitempty"`
	UserGroups      []UserGroupScope      `json:"user_groups,omitempty" xml:"user_groups>user_group,omitempty"`
	NetworkSegments []NetworkSegmentScope `json:"network_segments,omitempty" xml:"network_segments>network_segment,omitempty"`
	IBeacons        []IBeaconScope        `json:"ibeacons,omitempty" xml:"ibeacons>ibeacon,omitempty"`
}
type MacOSConfigurationProfileScopeExclusions struct {
	Computers       []ComputerScope             `json:"computers,omitempty" xml:"computers>computer,omitempty"`
	Buildings       []BuildingScope             `json:"buildings,omitempty" xml:"buildings>building,omitempty"`
	Departments     []DepartmentScope           `json:"departments,omitempty" xml:"departments>department,omitempty"`
	ComputerGroups  []ComputerGroupListResponse `json:"computer_groups,omitempty" xml:"computer_groups>computer_group,omitempty"`
	Users           []UserScope                 `json:"users,omitempty" xml:"users>user,omitempty"`
	UserGroups      []UserGroupScope            `json:"user_groups,omitempty" xml:"user_groups>user_group,omitempty"`
	NetworkSegments []NetworkSegmentScope       `json:"network_segments,omitempty" xml:"network_segments>network_segment,omitempty"`
	IBeacons        []IBeaconScope              `json:"ibeacons,omitempty" xml:"ibeacons>ibeacon,omitempty"`
	JamfUsers       []JamfUserScope             `json:"jss_users,omitempty" xml:"jss_users>jss_user,omitempty"`
	JamfUserGroups  []UserGroupScope            `json:"jss_user_groups,omitempty" xml:"jss_user_groups>jss_user_group,omitempty"`
}

type MacOSConfigurationProfileSelfService struct {
	SelfServiceDisplayName      string                `json:"self_service_display_name,omitempty" xml:"self_service_display_name,omitempty"`
	InstallButtonText           string                `json:"install_button_text,omitempty" xml:"install_button_text,omitempty"`
	SelfServiceDescription      string                `json:"self_service_description,omitempty" xml:"self_service_description,omitempty"`
	ForceUsersToViewDescription bool                  `json:"force_users_to_view_description,omitempty" xml:"force_users_to_view_description,omitempty"`
	Security                    SelfServiceSecurity   `json:"security,omitempty" xml:"security,omitempty"`
	SelfServiceIcon             SelfServiceIcon       `json:"self_service_icon,omitempty" xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                  `json:"feature_on_main_page,omitempty" xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories       []SelfServiceCategory `json:"self_service_categories,omitempty" xml:"self_service_categories>category,omitempty"`
}

type SelfServiceSecurity struct {
	RemovalDisallowed string `json:"removal_disallowed,omitempty" xml:"removal_disallowed,omitempty"`
}

type ResponseMacOSConfigurationProfileList struct {
	Size    int                                 `json:"size" xml:"size"`
	Results []MacOSConfigurationProfileListItem `json:"os_x_configuration_profile" xml:"os_x_configuration_profile"`
}

type MacOSConfigurationProfileListItem struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// GetMacOSConfigurationProfileByID retrieves the macOS ConfigurationProfile by its ID
func (c *Client) GetMacOSConfigurationProfileByID(id int) (*ResponseMacOSConfigurationProfile, error) {
	url := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	var profile ResponseMacOSConfigurationProfile
	if err := c.DoRequest("GET", url, nil, nil, &profile); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &profile, nil
}

// GetMacOSConfigurationProfileByName retrieves the macOS ConfigurationProfile by its Name
func (c *Client) GetMacOSConfigurationProfileByName(name string) (*ResponseMacOSConfigurationProfile, error) {
	url := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	var profile ResponseMacOSConfigurationProfile
	if err := c.DoRequest("GET", url, nil, nil, &profile); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &profile, nil
}

// GetMacOSConfigurationProfiles retrieves all macOS ConfigurationProfiles
func (c *Client) GetMacOSConfigurationProfiles() (*ResponseMacOSConfigurationProfileList, error) {
	url := uriMacOSConfigurationProfiles

	var profiles ResponseMacOSConfigurationProfileList
	if err := c.DoRequest("GET", url, nil, nil, &profiles); err != nil {
		return nil, fmt.Errorf("failed to fetch all OS X Configuration Profiles: %v", err)
	}

	return &profiles, nil
}

// GetMacOSConfigurationProfileByDataSubset retrieves all macOS ConfigurationProfiles by data subset
func (c *Client) GetMacOSConfigurationProfileIdAndDataSubset(id int, subsets ...string) (*ResponseMacOSConfigurationProfile, error) {
	// Combining all subsets with "&"
	subsetString := strings.Join(subsets, "&")
	url := fmt.Sprintf("%s/id/%d/subset/%s", uriMacOSConfigurationProfiles, id, subsetString)

	var profile ResponseMacOSConfigurationProfile
	if err := c.DoRequest("GET", url, nil, nil, &profile); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &profile, nil
}

// GetMacOSConfigurationProfileNameAndDataSubset retrieves macOS ConfigurationProfile by its Name and subset
func (c *Client) GetMacOSConfigurationProfileNameAndDataSubset(name string, subsets ...string) (*ResponseMacOSConfigurationProfile, error) {
	// Combining all subsets with "&"
	subsetString := strings.Join(subsets, "&")
	url := fmt.Sprintf("%s/name/%s/subset/%s", uriMacOSConfigurationProfiles, name, subsetString)

	var profile ResponseMacOSConfigurationProfile
	if err := c.DoRequest("GET", url, nil, nil, &profile); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &profile, nil
}

// CreateMacOSConfigurationProfile creates a new macOS Configuration Profile
func (c *Client) CreateMacOSConfigurationProfile(profile *ResponseMacOSConfigurationProfile) (*ResponseMacOSConfigurationProfile, error) {
	url := fmt.Sprintf("%s/id/0", uriMacOSConfigurationProfiles)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfile
	}{
		ResponseMacOSConfigurationProfile: profile,
	}

	// Execute the request
	var responseProfile ResponseMacOSConfigurationProfile
	if err := c.DoRequest("POST", url, reqBody, nil, &responseProfile); err != nil {
		return nil, fmt.Errorf("failed to create OS X Configuration Profile: %v", err)
	}

	return &responseProfile, nil
}

// UpdateMacOSConfigurationProfileById updates an existing macOS Configuration Profile by its ID
func (c *Client) UpdateMacOSConfigurationProfileById(id int, profile *ResponseMacOSConfigurationProfile) (*ResponseMacOSConfigurationProfile, error) {
	url := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfile
	}{
		ResponseMacOSConfigurationProfile: profile,
	}

	// Execute the request
	var responseProfile ResponseMacOSConfigurationProfile
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseProfile); err != nil {
		return nil, fmt.Errorf("failed to update OS X Configuration Profile: %v", err)
	}

	return &responseProfile, nil
}

// UpdateMacOSConfigurationProfileByName updates an existing macOS Configuration Profile by its name
func (c *Client) UpdateMacOSConfigurationProfileByName(name string, profile *ResponseMacOSConfigurationProfile) (*ResponseMacOSConfigurationProfile, error) {
	url := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfile
	}{
		ResponseMacOSConfigurationProfile: profile,
	}

	// Execute the request
	var responseProfile ResponseMacOSConfigurationProfile
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseProfile); err != nil {
		return nil, fmt.Errorf("failed to update OS X Configuration Profile by name: %v", err)
	}

	return &responseProfile, nil
}

// DeleteMacOSConfigurationProfileById deletes an existing macOS Configuration Profile by its ID
func (c *Client) DeleteMacOSConfigurationProfileById(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete OS X Configuration Profile: %v", err)
	}

	return nil
}

// DeleteMacOSConfigurationProfileByName deletes an existing macOS Configuration Profile by its name
func (c *Client) DeleteMacOSConfigurationProfileByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete macOS Configuration Profile by name: %v", err)
	}

	return nil
}
