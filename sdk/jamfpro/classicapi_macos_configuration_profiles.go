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

type ResponseMacOSConfigurationProfile struct {
	General     GeneralConfig     `json:"general,omitempty" xml:"general,omitempty"`
	Scope       ScopeConfig       `json:"scope,omitempty" xml:"scope,omitempty"`
	SelfService SelfServiceConfig `json:"self_service,omitempty" xml:"self_service,omitempty"`
}

type GeneralConfig struct {
	ID                 int          `json:"id,omitempty" xml:"id,omitempty"`
	Name               string       `json:"name,omitempty" xml:"name,omitempty"`
	Description        string       `json:"description,omitempty" xml:"description,omitempty"`
	Site               SiteInfo     `json:"site,omitempty" xml:"site,omitempty"`
	Category           CategoryInfo `json:"category,omitempty" xml:"category,omitempty"`
	DistributionMethod string       `json:"distribution_method,omitempty" xml:"distribution_method,omitempty"`
	UserRemovable      bool         `json:"user_removable,omitempty" xml:"user_removable,omitempty"`
	Level              string       `json:"level,omitempty" xml:"level,omitempty"`
	UUID               string       `json:"uuid,omitempty" xml:"uuid,omitempty"`
	RedeployOnUpdate   string       `json:"redeploy_on_update,omitempty" xml:"redeploy_on_update,omitempty"`
	Payloads           string       `json:"payloads,omitempty" xml:"payloads,omitempty"`
}

type SiteInfo struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type CategoryInfo struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type ScopeConfig struct {
	AllComputers   bool                   `json:"all_computers,omitempty" xml:"all_computers,omitempty"`
	AllJSSUsers    bool                   `json:"all_jss_users,omitempty" xml:"all_jss_users,omitempty"`
	Computers      []ComputerAssignment   `json:"computers,omitempty" xml:"computers,omitempty"`
	Buildings      []BuildingAssignment   `json:"buildings,omitempty" xml:"buildings,omitempty"`
	Departments    []DepartmentAssignment `json:"departments,omitempty" xml:"departments,omitempty"`
	ComputerGroups []ComputerGroupInfo    `json:"computer_groups,omitempty" xml:"computer_groups,omitempty"`
	JSSUsers       []JSSUserConfig        `json:"jss_users,omitempty" xml:"jss_users,omitempty"`
	JSSUserGroups  []JSSUserGroupInfo     `json:"jss_user_groups,omitempty" xml:"jss_user_groups,omitempty"`
	Limitations    NetworkLimitation      `json:"limitations,omitempty" xml:"limitations,omitempty"`
	Exclusions     ExclusionConfig        `json:"exclusions,omitempty" xml:"exclusions,omitempty"`
}

type ComputerAssignment struct {
	Computer ComputerInfo `json:"computer,omitempty" xml:"computer,omitempty"`
}

type ComputerInfo struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
	UDID string `json:"udid,omitempty" xml:"udid,omitempty"`
}

type BuildingAssignment struct {
	Building BuildingInfo `json:"building,omitempty" xml:"building,omitempty"`
}

type BuildingInfo struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type DepartmentAssignment struct {
	Department DepartmentInfo `json:"department,omitempty" xml:"department,omitempty"`
}

type DepartmentInfo struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type ComputerGroupInfo struct {
	ComputerGroup GroupAssignmentDataSubsetGroupDetail `json:"computer_group,omitempty" xml:"computer_group,omitempty"`
}

type JSSUserConfig struct {
	JSSUser UserInfo `json:"jss_user,omitempty" xml:"jss_user,omitempty"`
}

type UserInfo struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type JSSUserGroupInfo struct {
	JSSUserGroup GroupAssignmentDataSubsetGroupDetail `json:"jss_user_group,omitempty" xml:"jss_user_group,omitempty"`
}

type NetworkLimitation struct {
	Users           []UserAssignment      `json:"users,omitempty" xml:"users,omitempty"`
	UserGroups      []UserGroupAssignment `json:"user_groups,omitempty" xml:"user_groups,omitempty"`
	NetworkSegments []NetworkSegmentInfo  `json:"network_segments,omitempty" xml:"network_segments,omitempty"`
	IBeacons        []IBeaconInfo         `json:"ibeacons,omitempty" xml:"ibeacons,omitempty"`
}

type UserAssignment struct {
	User UserInfo `json:"user,omitempty" xml:"user,omitempty"`
}

type UserGroupAssignment struct {
	UserGroup GroupAssignmentDataSubsetGroupDetail `json:"user_group,omitempty" xml:"user_group,omitempty"`
}

type GroupAssignmentDataSubsetGroupDetail struct {
	ID           int                         `json:"id,omitempty" xml:"id"`
	Name         string                      `json:"name" xml:"name"`
	AccessLevel  string                      `json:"access_level" xml:"access_level"`
	PrivilegeSet string                      `json:"privilege_set" xml:"privilege_set"`
	Site         AccountDataSubsetSite       `json:"site" xml:"site"`
	Privileges   AccountDataSubsetPrivileges `json:"privileges" xml:"privileges"`
	Members      []AccountDataSubsetMembers  `json:"members" xml:"members>user"`
}

type AccountDataSubsetMembers struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type NetworkSegmentInfo struct {
	NetworkSegment NetworkDetail `json:"network_segment,omitempty" xml:"network_segment,omitempty"`
}

type NetworkDetail struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	UID  string `json:"uid,omitempty" xml:"uid,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type IBeaconInfo struct {
	IBeacon BeaconDetail `json:"ibeacon,omitempty" xml:"ibeacon,omitempty"`
}

type BeaconDetail struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name,omitempty" xml:"name,omitempty"`
}

type ExclusionConfig struct {
	Computers       []ComputerAssignment   `json:"computers,omitempty" xml:"computers,omitempty"`
	Buildings       []BuildingAssignment   `json:"buildings,omitempty" xml:"buildings,omitempty"`
	Departments     []DepartmentAssignment `json:"departments,omitempty" xml:"departments,omitempty"`
	ComputerGroups  []ComputerGroupInfo    `json:"computer_groups,omitempty" xml:"computer_groups,omitempty"`
	Users           []UserAssignment       `json:"users,omitempty" xml:"users,omitempty"`
	UserGroups      []UserGroupAssignment  `json:"user_groups,omitempty" xml:"user_groups,omitempty"`
	NetworkSegments []NetworkSegmentInfo   `json:"network_segments,omitempty" xml:"network_segments,omitempty"`
	IBeacons        []IBeaconInfo          `json:"ibeacons,omitempty" xml:"ibeacons,omitempty"`
	JSSUsers        []JSSUserConfig        `json:"jss_users,omitempty" xml:"jss_users,omitempty"`
	JSSUserGroups   []JSSUserGroupInfo     `json:"jss_user_groups,omitempty" xml:"jss_user_groups,omitempty"`
}

type SelfServiceConfig struct {
	InstallButtonText           string                `json:"install_button_text,omitempty" xml:"install_button_text,omitempty"`
	SelfServiceDescription      string                `json:"self_service_description,omitempty" xml:"self_service_description,omitempty"`
	ForceUsersToViewDescription bool                  `json:"force_users_to_view_description,omitempty" xml:"force_users_to_view_description,omitempty"`
	SelfServiceIcon             SelfServiceIconDetail `json:"self_service_icon,omitempty" xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                  `json:"feature_on_main_page,omitempty" xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories       SelfServiceCategory   `json:"self_service_categories,omitempty" xml:"self_service_categories,omitempty"`
	Notification                string                `json:"notification,omitempty" xml:"notification,omitempty"`
	NotificationSubject         string                `json:"notification_subject,omitempty" xml:"notification_subject,omitempty"`
	NotificationMessage         string                `json:"notification_message,omitempty" xml:"notification_message,omitempty"`
}

type SelfServiceIconDetail struct {
	ID   int    `json:"id,omitempty" xml:"id,omitempty"`
	URI  string `json:"uri,omitempty" xml:"uri,omitempty"`
	Data string `json:"data,omitempty" xml:"data,omitempty"`
}

//--------structs for lists----------//

type ResponseMacOSConfigurationProfileList struct {
	Size    int                                 `json:"size" xml:"size"`
	Results []MacOSConfigurationProfileListItem `json:"os_x_configuration_profile" xml:"os_x_configuration_profile"`
}

type MacOSConfigurationProfileListItem struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

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

func (c *Client) GetMacOSConfigurationProfileByID(id int) (*ResponseMacOSConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	var profile ResponseMacOSConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS Configuration Profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

func (c *Client) GetMacOSConfigurationProfileByName(name string) (*ResponseMacOSConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	var profile ResponseMacOSConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS Configuration Profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

func (c *Client) GetMacOSConfigurationProfileNameByID(id int) (string, error) {
	profile, err := c.GetMacOSConfigurationProfileByID(id)
	if err != nil {
		return "", fmt.Errorf("failed to fetch macOS Configuration Profile by ID: %v", err)
	}
	return profile.General.Name, nil
}

func (c *Client) CreateMacOSConfigurationProfile(profile *ResponseMacOSConfigurationProfile) (*ResponseMacOSConfigurationProfile, error) {
	endpoint := uriMacOSConfigurationProfiles

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfile
	}{
		ResponseMacOSConfigurationProfile: profile,
	}

	var responseProfile ResponseMacOSConfigurationProfile
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to create macOS Configuration Profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

func (c *Client) UpdateMacOSConfigurationProfileByID(id int, profile *ResponseMacOSConfigurationProfile) (*ResponseMacOSConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfile
	}{
		ResponseMacOSConfigurationProfile: profile,
	}

	var updatedProfile ResponseMacOSConfigurationProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update macOS Configuration Profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

func (c *Client) UpdateMacOSConfigurationProfileByName(name string, profile *ResponseMacOSConfigurationProfile) (*ResponseMacOSConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfile
	}{
		ResponseMacOSConfigurationProfile: profile,
	}

	var updatedProfile ResponseMacOSConfigurationProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update macOS Configuration Profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

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
