// classicapi_macos_configuration_profiles.go
// Jamf Pro Classic Api - osx configuration profiles
// api reference: https://developer.jamf.com/jamf-pro/reference/osxconfigurationprofiles
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMacOSConfigurationProfiles = "/JSSResource/osxconfigurationprofiles"

// ResponseMacOSConfigurationProfileList represents the response structure for a list of macOS configuration profiles.
type ResponseMacOSConfigurationProfileList struct {
	Size    int                                 `json:"size" xml:"size"`
	Results []MacOSConfigurationProfileListItem `json:"os_x_configuration_profile" xml:"os_x_configuration_profile"`
}

type MacOSConfigurationProfileListItem struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// ResponseMacOSConfigurationProfiles represents the response structure for a macOS configuration profile.
type ResponseMacOSConfigurationProfiles struct {
	General     MacOSConfigurationProfilesDataSubsetGeneral     `json:"general"`
	Scope       MacOSConfigurationProfilesDataSubsetScope       `json:"scope"`
	SelfService MacOSConfigurationProfilesDataSubsetSelfService `json:"self_service"`
}

type MacOSConfigurationProfilesDataSubsetGeneral struct {
	ID                 int                                          `json:"id"`
	Name               string                                       `json:"name"`
	Description        string                                       `json:"description"`
	Site               MacOSConfigurationProfilesDataSubsetSite     `json:"site"`
	Category           MacOSConfigurationProfilesDataSubsetCategory `json:"category"`
	DistributionMethod string                                       `json:"distribution_method"`
	UserRemovable      bool                                         `json:"user_removable"`
	Level              string                                       `json:"level"`
	UUID               string                                       `json:"uuid"`
	RedeployOnUpdate   string                                       `json:"redeploy_on_update"`
	Payloads           string                                       `json:"payloads"`
}

type MacOSConfigurationProfilesDataSubsetSite struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetScope struct {
	AllComputers   bool                                                `json:"all_computers"`
	AllJSSUsers    bool                                                `json:"all_jss_users"`
	Computers      []MacOSConfigurationProfilesDataSubsetComputer      `json:"computers"`
	Buildings      []MacOSConfigurationProfilesDataSubsetBuilding      `json:"buildings"`
	Departments    []MacOSConfigurationProfilesDataSubsetDepartment    `json:"departments"`
	ComputerGroups []MacOSConfigurationProfilesDataSubsetComputerGroup `json:"computer_groups"`
	JSSUsers       []MacOSConfigurationProfilesDataSubsetJSSUser       `json:"jss_users"`
	JSSUserGroups  []MacOSConfigurationProfilesDataSubsetJSSUserGroup  `json:"jss_user_groups"`
	Limitations    MacOSConfigurationProfilesDataSubsetLimitations     `json:"limitations"`
	Exclusions     MacOSConfigurationProfilesDataSubsetExclusions      `json:"exclusions"`
}

type MacOSConfigurationProfilesDataSubsetComputer struct {
	Computer MacOSConfigurationProfilesDataSubsetComputerItem `json:"computer"`
}

type MacOSConfigurationProfilesDataSubsetComputerItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	UDID string `json:"udid"`
}

type MacOSConfigurationProfilesDataSubsetBuilding struct {
	Building MacOSConfigurationProfilesDataSubsetBuildingItem `json:"building"`
}

type MacOSConfigurationProfilesDataSubsetBuildingItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetDepartment struct {
	Department MacOSConfigurationProfilesDataSubsetDepartmentItem `json:"department"`
}

type MacOSConfigurationProfilesDataSubsetDepartmentItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetComputerGroup struct {
	ComputerGroup MacOSConfigurationProfilesDataSubsetComputerGroupItem `json:"computer_group"`
}

type MacOSConfigurationProfilesDataSubsetComputerGroupItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetLimitations struct {
	Users           []MacOSConfigurationProfilesDataSubsetUser           `json:"users"`
	UserGroups      []MacOSConfigurationProfilesDataSubsetUserGroup      `json:"user_groups"`
	NetworkSegments []MacOSConfigurationProfilesDataSubsetNetworkSegment `json:"network_segments"`
	IBeacons        []MacOSConfigurationProfilesDataSubsetIBeacon        `json:"ibeacons"`
}

type MacOSConfigurationProfilesDataSubsetUser struct {
	User MacOSConfigurationProfilesDataSubsetUserItem `json:"user"`
}

type MacOSConfigurationProfilesDataSubsetUserItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetUserGroup struct {
	UserGroup MacOSConfigurationProfilesDataSubsetUserGroupItem `json:"user_group"`
}

type MacOSConfigurationProfilesDataSubsetUserGroupItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetNetworkSegment struct {
	NetworkSegment MacOSConfigurationProfilesDataSubsetNetworkSegmentItem `json:"network_segment"`
}

type MacOSConfigurationProfilesDataSubsetNetworkSegmentItem struct {
	ID   int    `json:"id"`
	UID  string `json:"uid,omitempty"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetIBeacon struct {
	IBeacon MacOSConfigurationProfilesDataSubsetIBeaconItem `json:"ibeacon"`
}

type MacOSConfigurationProfilesDataSubsetIBeaconItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetJSSUser struct {
	JSSUser MacOSConfigurationProfilesDataSubsetJSSUserItem `json:"jss_user"`
}

type MacOSConfigurationProfilesDataSubsetJSSUserItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetJSSUserGroup struct {
	JSSUserGroup MacOSConfigurationProfilesDataSubsetJSSUserGroupItem `json:"jss_user_group"`
}

type MacOSConfigurationProfilesDataSubsetJSSUserGroupItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type MacOSConfigurationProfilesDataSubsetExclusions struct {
	Computers       []MacOSConfigurationProfilesDataSubsetComputer       `json:"computers"`
	Buildings       []MacOSConfigurationProfilesDataSubsetBuilding       `json:"buildings"`
	Departments     []MacOSConfigurationProfilesDataSubsetDepartment     `json:"departments"`
	ComputerGroups  []MacOSConfigurationProfilesDataSubsetComputerGroup  `json:"computer_groups"`
	Users           []MacOSConfigurationProfilesDataSubsetUser           `json:"users"`
	UserGroups      []MacOSConfigurationProfilesDataSubsetUserGroup      `json:"user_groups"`
	NetworkSegments []MacOSConfigurationProfilesDataSubsetNetworkSegment `json:"network_segments"`
	IBeacons        []MacOSConfigurationProfilesDataSubsetIBeacon        `json:"ibeacons"`
	JSSUsers        []MacOSConfigurationProfilesDataSubsetJSSUser        `json:"jss_users"`
	JSSUserGroups   []MacOSConfigurationProfilesDataSubsetJSSUserGroup   `json:"jss_user_groups"`
}

type MacOSConfigurationProfilesDataSubsetSelfService struct {
	InstallButtonText           string                                                    `json:"install_button_text"`
	SelfServiceDescription      string                                                    `json:"self_service_description"`
	ForceUsersToViewDescription bool                                                      `json:"force_users_to_view_description"`
	SelfServiceIcon             MacOSConfigurationProfilesDataSubsetSelfServiceIcon       `json:"self_service_icon"`
	FeatureOnMainPage           bool                                                      `json:"feature_on_main_page"`
	SelfServiceCategories       MacOSConfigurationProfilesDataSubsetSelfServiceCategories `json:"self_service_categories"`
	Notification                string                                                    `json:"notification"`
	NotificationSubject         string                                                    `json:"notification_subject"`
	NotificationMessage         string                                                    `json:"notification_message"`
}

type MacOSConfigurationProfilesDataSubsetSelfServiceIcon struct {
	ID   int    `json:"id"`
	URI  string `json:"uri"`
	Data string `json:"data"`
}

type MacOSConfigurationProfilesDataSubsetSelfServiceCategories struct {
	Category MacOSConfigurationProfilesDataSubsetSelfServiceCategory `json:"category"`
}

type MacOSConfigurationProfilesDataSubsetSelfServiceCategory struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	DisplayIn bool   `json:"display_in"`
	FeatureIn bool   `json:"feature_in"`
}

// GetMacOSConfigurationProfiles fetches a list of all macOS Configuration Profiles from the Jamf Pro server.
func (c *Client) GetMacOSConfigurationProfiles() (*ResponseMacOSConfigurationProfileList, error) {
	endpoint := uriMacOSConfigurationProfiles

	var profilesList ResponseMacOSConfigurationProfileList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profilesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all OS X Configuration Profiles: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profilesList, nil
}

// GetMacOSConfigurationProfileByID fetches a specific macOS Configuration Profile by its ID from the Jamf Pro server.
func (c *Client) GetMacOSConfigurationProfileByID(id int) (*ResponseMacOSConfigurationProfiles, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	var profile ResponseMacOSConfigurationProfiles
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS Configuration Profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMacOSConfigurationProfileByName fetches a specific macOS Configuration Profile by its name from the Jamf Pro server.
func (c *Client) GetMacOSConfigurationProfileByName(name string) (*ResponseMacOSConfigurationProfiles, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	var profile ResponseMacOSConfigurationProfiles
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS Configuration Profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMacOSConfigurationProfileNameByID retrieves the name of a macOS Configuration Profile by its ID.
func (c *Client) GetMacOSConfigurationProfileNameByID(id int) (string, error) {
	profile, err := c.GetMacOSConfigurationProfileByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to fetch macOS Configuration Profile by ID: %v", err)
	}
	return profile.General.Name, nil
}

// CreateMacOSConfigurationProfile creates a new macOS Configuration Profile on the Jamf Pro server.
func (c *Client) CreateMacOSConfigurationProfile(profile *ResponseMacOSConfigurationProfiles) (*ResponseMacOSConfigurationProfiles, error) {
	endpoint := uriMacOSConfigurationProfiles

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfiles
	}{
		ResponseMacOSConfigurationProfiles: profile,
	}

	var responseProfile ResponseMacOSConfigurationProfiles
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to create macOS Configuration Profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// UpdateMacOSConfigurationProfileByID updates an existing macOS Configuration Profile by its ID on the Jamf Pro server.
func (c *Client) UpdateMacOSConfigurationProfileByID(id int, profile *ResponseMacOSConfigurationProfiles) (*ResponseMacOSConfigurationProfiles, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfiles
	}{
		ResponseMacOSConfigurationProfiles: profile,
	}

	var updatedProfile ResponseMacOSConfigurationProfiles
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update macOS Configuration Profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

// UpdateMacOSConfigurationProfileByName updates an existing macOS Configuration Profile by its name on the Jamf Pro server.
func (c *Client) UpdateMacOSConfigurationProfileByName(name string, profile *ResponseMacOSConfigurationProfiles) (*ResponseMacOSConfigurationProfiles, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfiles
	}{
		ResponseMacOSConfigurationProfiles: profile,
	}

	var updatedProfile ResponseMacOSConfigurationProfiles
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update macOS Configuration Profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

// DeleteMacOSConfigurationProfileByID deletes a macOS Configuration Profile by its ID from the Jamf Pro server.
func (c *Client) DeleteMacOSConfigurationProfileByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete macOS Configuration Profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMacOSConfigurationProfileByName deletes a macOS Configuration Profile by its name from the Jamf Pro server.
func (c *Client) DeleteMacOSConfigurationProfileByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete macOS Configuration Profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
