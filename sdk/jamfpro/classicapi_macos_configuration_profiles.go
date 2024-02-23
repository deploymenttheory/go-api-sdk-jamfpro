// classicapi_macos_configuration_profiles.go
// Jamf Pro Classic Api - osx configuration profiles
// api reference: https://developer.jamf.com/jamf-pro/reference/osxconfigurationprofiles
// Entity_Equivalents_for_Disallowed_XML_Characters: https://learn.jamf.com/bundle/technical-articles/page/Entity_Equivalents_for_Disallowed_XML_Characters.html
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
- SharedResourceCategory
- SharedResourceSelfServiceIcon
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMacOSConfigurationProfiles = "/JSSResource/osxconfigurationprofiles"

// List

// ResponseMacOSConfigurationProfileList represents the response structure for a list of macOS configuration profiles.
type ResponseMacOSConfigurationProfileList struct {
	Size    int                                 `xml:"size,omitempty"`
	Results []MacOSConfigurationProfileListItem `xml:"os_x_configuration_profile,omitempty"`
}

type MacOSConfigurationProfileListItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name"`
}

// Responses

// ResponseMacOSConfigurationProfileCreation represents the response structure for a new macOS configuration profile.
type ResponseMacOSConfigurationProfileCreationUpdate struct {
	XMLName xml.Name `xml:"os_x_configuration_profile"`
	ID      int      `xml:"id"`
}

// Resource

// ResourceMacOSConfigurationProfiles represents the response structure for a macOS configuration profile.
type ResourceMacOSConfigurationProfile struct {
	General     MacOSConfigurationProfileSubsetGeneral     `xml:"general,omitempty"`
	Scope       MacOSConfigurationProfileSubsetScope       `xml:"scope,omitempty"`
	SelfService MacOSConfigurationProfileSubsetSelfService `xml:"self_service,omitempty"`
}

// Subsets and Containers

type MacOSConfigurationProfileSubsetGeneral struct {
	ID                 int                    `xml:"id,omitempty"`
	Name               string                 `xml:"name"`
	Description        string                 `xml:"description,omitempty"`
	Site               SharedResourceSite     `xml:"site,omitempty"`
	Category           SharedResourceCategory `xml:"category,omitempty"`
	DistributionMethod string                 `xml:"distribution_method,omitempty"`
	UserRemovable      bool                   `xml:"user_removable,omitempty"`
	Level              string                 `xml:"level,omitempty"`
	UUID               string                 `xml:"uuid,omitempty"`
	RedeployOnUpdate   string                 `xml:"redeploy_on_update,omitempty"`
	Payloads           string                 `xml:"payloads,omitempty"`
}

type MacOSConfigurationProfileSubsetScope struct {
	AllComputers   bool                                           `xml:"all_computers,omitempty"`
	AllJSSUsers    bool                                           `xml:"all_jss_users,omitempty"`
	Computers      []MacOSConfigurationProfileSubsetComputer      `xml:"computers>computer,omitempty"`
	ComputerGroups []MacOSConfigurationProfileSubsetComputerGroup `xml:"computer_groups>computer_group,omitempty"`
	JSSUsers       []MacOSConfigurationProfileSubsetJSSUser       `xml:"jss_users>jss_user,omitempty"`
	JSSUserGroups  []MacOSConfigurationProfileSubsetJSSUserGroup  `xml:"jss_user_groups>jss_user_group,omitempty"`
	Buildings      []MacOSConfigurationProfileSubsetBuilding      `xml:"buildings>building,omitempty"`
	Departments    []MacOSConfigurationProfileSubsetDepartment    `xml:"departments>department,omitempty"`
	Limitations    MacOSConfigurationProfileSubsetLimitations     `xml:"limitations,omitempty"`
	Exclusions     MacOSConfigurationProfileSubsetExclusions      `xml:"exclusions,omitempty"`
}

type MacOSConfigurationProfileSubsetSelfService struct {
	InstallButtonText           string                                               `xml:"install_button_text,omitempty"`
	SelfServiceDescription      string                                               `xml:"self_service_description,omitempty"`
	ForceUsersToViewDescription bool                                                 `xml:"force_users_to_view_description,omitempty"`
	SelfServiceIcon             SharedResourceSelfServiceIcon                        `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                                                 `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories       MacOSConfigurationProfileSubsetSelfServiceCategories `xml:"self_service_categories,omitempty"`
	Notification                string                                               `xml:"notification,omitempty"`
	NotificationSubject         string                                               `xml:"notification_subject,omitempty"`
	NotificationMessage         string                                               `xml:"notification_message,omitempty"`
}

type MacOSConfigurationProfileSubsetSelfServiceCategories struct {
	Category MacOSConfigurationProfileSubsetSelfServiceCategory `xml:"category,omitempty"`
}

type MacOSConfigurationProfileSubsetSelfServiceCategory struct {
	ID        int    `xml:"id,omitempty"`
	Name      string `xml:"name,omitempty"`
	DisplayIn bool   `xml:"display_in,omitempty"`
	FeatureIn bool   `xml:"feature_in,omitempty"`
}

type MacOSConfigurationProfileSubsetComputer struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UDID string `xml:"udid,omitempty"`
}

type MacOSConfigurationProfileSubsetComputerGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfileSubsetJSSUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfileSubsetJSSUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfileSubsetBuilding struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfileSubsetDepartment struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
type MacOSConfigurationProfileSubsetLimitations struct {
	Users           []MacOSConfigurationProfileSubsetUser           `xml:"users>user,omitempty"`
	UserGroups      []MacOSConfigurationProfileSubsetUserGroup      `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []MacOSConfigurationProfileSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	IBeacons        []MacOSConfigurationProfileSubsetIBeacon        `xml:"ibeacons>ibeacon,omitempty"`
}
type MacOSConfigurationProfileSubsetIBeacon struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
type MacOSConfigurationProfileSubsetNetworkSegment struct {
	ID   int    `xml:"id,omitempty"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name,omitempty"`
}
type MacOSConfigurationProfileSubsetUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfileSubsetUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfileSubsetExclusions struct {
	Computers       []MacOSConfigurationProfileSubsetComputer       `xml:"computers,omitempty"`
	ComputerGroups  []MacOSConfigurationProfileSubsetComputerGroup  `xml:"computer_groups,omitempty"`
	Users           []MacOSConfigurationProfileSubsetUser           `xml:"users,omitempty"`
	UserGroups      []MacOSConfigurationProfileSubsetUserGroup      `xml:"user_groups,omitempty"`
	Buildings       []MacOSConfigurationProfileSubsetBuilding       `xml:"buildings,omitempty"`
	Departments     []MacOSConfigurationProfileSubsetDepartment     `xml:"departments,omitempty"`
	NetworkSegments []MacOSConfigurationProfileSubsetNetworkSegment `xml:"network_segments,omitempty"`
	JSSUsers        []MacOSConfigurationProfileSubsetJSSUser        `xml:"jss_users,omitempty"`
	JSSUserGroups   []MacOSConfigurationProfileSubsetJSSUserGroup   `xml:"jss_user_groups,omitempty"`
	IBeacons        []MacOSConfigurationProfileSubsetIBeacon        `xml:"ibeacons,omitempty"`
}

// CRUD

// GetMacOSConfigurationProfiles fetches a list of all macOS Configuration Profiles from the Jamf Pro server.
func (c *Client) GetMacOSConfigurationProfiles() (*ResponseMacOSConfigurationProfileList, error) {
	endpoint := uriMacOSConfigurationProfiles

	var profilesList ResponseMacOSConfigurationProfileList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profilesList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "mac os config profiles", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profilesList, nil
}

// GetMacOSConfigurationProfileByID fetches a specific macOS Configuration Profile by its ID from the Jamf Pro server.
func (c *Client) GetMacOSConfigurationProfileByID(id int) (*ResourceMacOSConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	var profile ResourceMacOSConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mac os config profile", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMacOSConfigurationProfileByName fetches a specific macOS Configuration Profile by its name from the Jamf Pro server.
func (c *Client) GetMacOSConfigurationProfileByName(name string) (*ResourceMacOSConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	var profile ResourceMacOSConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mac os config profile", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// QUERY Review this structure

// GetMacOSConfigurationProfileByNameByID retrieves the details of a macOS Configuration Profile by its name.
func (c *Client) GetMacOSConfigurationProfileByNameByID(name string) (*ResourceMacOSConfigurationProfile, error) {

	profilesList, err := c.GetMacOSConfigurationProfiles()
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "mac os config profiles", err)
	}

	var profileID int
	for _, profile := range profilesList.Results {
		if profile.Name == name {
			profileID = profile.ID
			break
		}
	}

	if profileID == 0 {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mac os config profile", name, err)
	}

	detailedProfile, err := c.GetMacOSConfigurationProfileByID(profileID)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mac os config profile", name, err)
	}

	return detailedProfile, nil
}

// CreateMacOSConfigurationProfile creates a new macOS Configuration Profile on the Jamf Pro server and returns the profile with its ID updated.
// It sends a POST request to the Jamf Pro server with the profile details and expects a response with the ID of the newly created profile.
// CreateMacOSConfigurationProfile creates a new macOS Configuration Profile on the Jamf Pro server and returns the ID of the newly created profile.
func (c *Client) CreateMacOSConfigurationProfile(profile *ResourceMacOSConfigurationProfile) (int, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMacOSConfigurationProfiles)

	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResourceMacOSConfigurationProfile
	}{
		ResourceMacOSConfigurationProfile: profile,
	}

	var response ResponseMacOSConfigurationProfileCreationUpdate

	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return 0, fmt.Errorf(errMsgFailedCreate, "mac os config profile", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return response.ID, nil
}

// UpdateMacOSConfigurationProfileByID updates an existing macOS Configuration Profile by its ID on the Jamf Pro server
// and returns the ID of the updated profile.
func (c *Client) UpdateMacOSConfigurationProfileByID(id int, profile *ResourceMacOSConfigurationProfile) (int, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResourceMacOSConfigurationProfile
	}{
		ResourceMacOSConfigurationProfile: profile,
	}

	var response ResponseMacOSConfigurationProfileCreationUpdate

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return 0, fmt.Errorf(errMsgFailedUpdateByID, "mac os config profile", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return response.ID, nil
}

// UpdateMacOSConfigurationProfileByName updates an existing macOS Configuration Profile by its name on the Jamf Pro server
// and returns the ID of the updated profile.
func (c *Client) UpdateMacOSConfigurationProfileByName(name string, profile *ResourceMacOSConfigurationProfile) (int, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResourceMacOSConfigurationProfile
	}{
		ResourceMacOSConfigurationProfile: profile,
	}

	var response ResponseMacOSConfigurationProfileCreationUpdate

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return 0, fmt.Errorf(errMsgFailedUpdateByName, "mac os config profile", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return response.ID, nil
}

// DeleteMacOSConfigurationProfileByID deletes a macOS Configuration Profile by its ID from the Jamf Pro server.
func (c *Client) DeleteMacOSConfigurationProfileByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "mac os config profile", id, err)
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
		return fmt.Errorf(errMsgFailedDeleteByName, "mac os config profile", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
