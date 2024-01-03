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
type ResourceMacOSConfigurationProfiles struct {
	General     MacOSConfigurationProfilesSubsetGeneral     `xml:"general,omitempty"`
	Scope       MacOSConfigurationProfilesSubsetScope       `xml:"scope,omitempty"`
	SelfService MacOSConfigurationProfilesSubsetSelfService `xml:"self_service,omitempty"`
}

// Subsets and Containers

type MacOSConfigurationProfilesSubsetGeneral struct {
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

type MacOSConfigurationProfilesSubsetScope struct {
	AllComputers   bool                                            `xml:"all_computers,omitempty"`
	AllJSSUsers    bool                                            `xml:"all_jss_users,omitempty"`
	Computers      []MacOSConfigurationProfilesSubsetComputer      `xml:"computers>computer,omitempty"`
	ComputerGroups []MacOSConfigurationProfilesSubsetComputerGroup `xml:"computer_groups>computer_group,omitempty"`
	JSSUsers       []MacOSConfigurationProfilesSubsetJSSUser       `xml:"jss_users>jss_user,omitempty"`
	JSSUserGroups  []MacOSConfigurationProfilesSubsetJSSUserGroup  `xml:"jss_user_groups>jss_user_group,omitempty"`
	Buildings      []MacOSConfigurationProfilesSubsetBuilding      `xml:"buildings>building,omitempty"`
	Departments    []MacOSConfigurationProfilesSubsetDepartment    `xml:"departments>department,omitempty"`
	Limitations    MacOSConfigurationProfilesSubsetLimitations     `xml:"limitations,omitempty"`
	Exclusions     MacOSConfigurationProfilesSubsetExclusions      `xml:"exclusions,omitempty"`
}

type MacOSConfigurationProfilesSubsetSelfService struct {
	InstallButtonText           string                                                `xml:"install_button_text,omitempty"`
	SelfServiceDescription      string                                                `xml:"self_service_description,omitempty"`
	ForceUsersToViewDescription bool                                                  `xml:"force_users_to_view_description,omitempty"`
	SelfServiceIcon             SharedResourceSelfServiceIcon                         `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                                                  `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories       MacOSConfigurationProfilesSubsetSelfServiceCategories `xml:"self_service_categories,omitempty"`
	Notification                string                                                `xml:"notification,omitempty"`
	NotificationSubject         string                                                `xml:"notification_subject,omitempty"`
	NotificationMessage         string                                                `xml:"notification_message,omitempty"`
}

type MacOSConfigurationProfilesSubsetSelfServiceCategories struct {
	Category MacOSConfigurationProfilesSubsetSelfServiceCategory `xml:"category,omitempty"`
}

type MacOSConfigurationProfilesSubsetSelfServiceCategory struct {
	ID        int    `xml:"id,omitempty"`
	Name      string `xml:"name,omitempty"`
	DisplayIn bool   `xml:"display_in,omitempty"`
	FeatureIn bool   `xml:"feature_in,omitempty"`
}

type MacOSConfigurationProfilesSubsetComputer struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UDID string `xml:"udid,omitempty"`
}

type MacOSConfigurationProfilesSubsetComputerGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesSubsetJSSUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesSubsetJSSUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesSubsetBuilding struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesSubsetDepartment struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
type MacOSConfigurationProfilesSubsetLimitations struct {
	Users           []MacOSConfigurationProfilesSubsetUser           `xml:"users>user,omitempty"`
	UserGroups      []MacOSConfigurationProfilesSubsetUserGroup      `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []MacOSConfigurationProfilesSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	IBeacons        []MacOSConfigurationProfilesSubsetIBeacon        `xml:"ibeacons>ibeacon,omitempty"`
}
type MacOSConfigurationProfilesSubsetIBeacon struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}
type MacOSConfigurationProfilesSubsetNetworkSegment struct {
	ID   int    `xml:"id,omitempty"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name,omitempty"`
}
type MacOSConfigurationProfilesSubsetUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesSubsetUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesSubsetExclusions struct {
	Computers       []MacOSConfigurationProfilesSubsetComputer       `xml:"computers,omitempty"`
	ComputerGroups  []MacOSConfigurationProfilesSubsetComputerGroup  `xml:"computer_groups,omitempty"`
	Users           []MacOSConfigurationProfilesSubsetUser           `xml:"users,omitempty"`
	UserGroups      []MacOSConfigurationProfilesSubsetUserGroup      `xml:"user_groups,omitempty"`
	Buildings       []MacOSConfigurationProfilesSubsetBuilding       `xml:"buildings,omitempty"`
	Departments     []MacOSConfigurationProfilesSubsetDepartment     `xml:"departments,omitempty"`
	NetworkSegments []MacOSConfigurationProfilesSubsetNetworkSegment `xml:"network_segments,omitempty"`
	JSSUsers        []MacOSConfigurationProfilesSubsetJSSUser        `xml:"jss_users,omitempty"`
	JSSUserGroups   []MacOSConfigurationProfilesSubsetJSSUserGroup   `xml:"jss_user_groups,omitempty"`
	IBeacons        []MacOSConfigurationProfilesSubsetIBeacon        `xml:"ibeacons,omitempty"`
}

// CRUD

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
func (c *Client) GetMacOSConfigurationProfileByID(id int) (*ResourceMacOSConfigurationProfiles, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	var profile ResourceMacOSConfigurationProfiles
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
func (c *Client) GetMacOSConfigurationProfileByName(name string) (*ResourceMacOSConfigurationProfiles, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	var profile ResourceMacOSConfigurationProfiles
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS Configuration Profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMacOSConfigurationProfileByNameByID retrieves the details of a macOS Configuration Profile by its name.
func (c *Client) GetMacOSConfigurationProfileByNameByID(name string) (*ResourceMacOSConfigurationProfiles, error) {

	profilesList, err := c.GetMacOSConfigurationProfiles()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS Configuration Profiles: %v", err)
	}

	var profileID int
	for _, profile := range profilesList.Results {
		if profile.Name == name {
			profileID = profile.ID
			break
		}
	}

	if profileID == 0 {
		return nil, fmt.Errorf("no macOS Configuration Profile found with the name %s", name)
	}

	detailedProfile, err := c.GetMacOSConfigurationProfileByID(profileID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS Configuration Profile by ID: %v", err)
	}

	return detailedProfile, nil
}

// CreateMacOSConfigurationProfile creates a new macOS Configuration Profile on the Jamf Pro server and returns the profile with its ID updated.
// It sends a POST request to the Jamf Pro server with the profile details and expects a response with the ID of the newly created profile.
// CreateMacOSConfigurationProfile creates a new macOS Configuration Profile on the Jamf Pro server and returns the ID of the newly created profile.
func (c *Client) CreateMacOSConfigurationProfile(profile *ResourceMacOSConfigurationProfiles) (int, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMacOSConfigurationProfiles)

	if profile.General.Site.ID == 0 && profile.General.Site.Name == "" {
		profile.General.Site.ID = -1
		profile.General.Site.Name = "none"

	}
	if profile.General.Category.ID == 0 && profile.General.Category.Name == "" {
		profile.General.Category.ID = -1
		profile.General.Category.Name = "No Category"
	}

	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResourceMacOSConfigurationProfiles
	}{
		ResourceMacOSConfigurationProfiles: profile,
	}

	var response ResponseMacOSConfigurationProfileCreationUpdate

	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to create macOS Configuration Profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return response.ID, nil
}

// UpdateMacOSConfigurationProfileByID updates an existing macOS Configuration Profile by its ID on the Jamf Pro server
// and returns the ID of the updated profile.
func (c *Client) UpdateMacOSConfigurationProfileByID(id int, profile *ResourceMacOSConfigurationProfiles) (int, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResourceMacOSConfigurationProfiles
	}{
		ResourceMacOSConfigurationProfiles: profile,
	}

	var response ResponseMacOSConfigurationProfileCreationUpdate

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to update macOS Configuration Profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return response.ID, nil
}

// UpdateMacOSConfigurationProfileByName updates an existing macOS Configuration Profile by its name on the Jamf Pro server
// and returns the ID of the updated profile.
func (c *Client) UpdateMacOSConfigurationProfileByName(name string, profile *ResourceMacOSConfigurationProfiles) (int, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResourceMacOSConfigurationProfiles
	}{
		ResourceMacOSConfigurationProfiles: profile,
	}

	var response ResponseMacOSConfigurationProfileCreationUpdate

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to update macOS Configuration Profile by name: %v", err)
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
