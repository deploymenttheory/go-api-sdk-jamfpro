// classicapi_macos_configuration_profiles.go
// Jamf Pro Classic Api - osx configuration profiles
// api reference: https://developer.jamf.com/jamf-pro/reference/osxconfigurationprofiles
// Entity_Equivalents_for_Disallowed_XML_Characters: https://learn.jamf.com/bundle/technical-articles/page/Entity_Equivalents_for_Disallowed_XML_Characters.html
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMacOSConfigurationProfiles = "/JSSResource/osxconfigurationprofiles"

// ResponseMacOSConfigurationProfileList represents the response structure for a list of macOS configuration profiles.
type ResponseMacOSConfigurationProfileList struct {
	Size    int                                 `xml:"size,omitempty"`
	Results []MacOSConfigurationProfileListItem `xml:"os_x_configuration_profile,omitempty"`
}

type MacOSConfigurationProfileListItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name" `
}

// ResponseMacOSConfigurationProfileCreation represents the response structure for a new macOS configuration profile.
type ResponseMacOSConfigurationProfileCreationAndUpdate struct {
	XMLName xml.Name `xml:"os_x_configuration_profile"`
	ID      int      `xml:"id"`
}

// ResponseMacOSConfigurationProfiles represents the response structure for a macOS configuration profile.
type ResponseMacOSConfigurationProfiles struct {
	General     MacOSConfigurationProfilesDataSubsetGeneral     `xml:"general,omitempty"`
	Scope       MacOSConfigurationProfilesDataSubsetScope       `xml:"scope,omitempty"`
	SelfService MacOSConfigurationProfilesDataSubsetSelfService `xml:"self_service,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetGeneral struct {
	ID                 int                                          `xml:"id,omitempty"`
	Name               string                                       `xml:"name"`
	Description        string                                       `xml:"description,omitempty"`
	Site               MacOSConfigurationProfilesDataSubsetSite     `xml:"site,omitempty"`
	Category           MacOSConfigurationProfilesDataSubsetCategory `xml:"category,omitempty"`
	DistributionMethod string                                       `xml:"distribution_method,omitempty"`
	UserRemovable      bool                                         `xml:"user_removable,omitempty"`
	Level              string                                       `xml:"level,omitempty"`
	UUID               string                                       `xml:"uuid,omitempty"`
	RedeployOnUpdate   string                                       `xml:"redeploy_on_update,omitempty"`
	Payloads           string                                       `xml:"payloads,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetSite struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetCategory struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetScope struct {
	AllComputers   bool                                                `xml:"all_computers,omitempty"`
	AllJSSUsers    bool                                                `xml:"all_jss_users,omitempty"`
	Computers      []MacOSConfigurationProfilesDataSubsetComputer      `xml:"computers,omitempty"`
	Buildings      []MacOSConfigurationProfilesDataSubsetBuilding      `xml:"buildings,omitempty"`
	Departments    []MacOSConfigurationProfilesDataSubsetDepartment    `xml:"departments,omitempty"`
	ComputerGroups []MacOSConfigurationProfilesDataSubsetComputerGroup `xml:"computer_groups,omitempty"`
	JSSUsers       []MacOSConfigurationProfilesDataSubsetJSSUser       `xml:"jss_users,omitempty"`
	JSSUserGroups  []MacOSConfigurationProfilesDataSubsetJSSUserGroup  `xml:"jss_user_groups,omitempty"`
	Limitations    MacOSConfigurationProfilesDataSubsetLimitations     `xml:"limitations,omitempty"`
	Exclusions     MacOSConfigurationProfilesDataSubsetExclusions      `xml:"exclusions,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetComputer struct {
	Computer MacOSConfigurationProfilesDataSubsetComputerItem `xml:"computer,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetComputerItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
	UDID string `xml:"udid,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetBuilding struct {
	Building MacOSConfigurationProfilesDataSubsetBuildingItem `xml:"building,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetBuildingItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetDepartment struct {
	Department MacOSConfigurationProfilesDataSubsetDepartmentItem `xml:"department,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetDepartmentItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetComputerGroup struct {
	ComputerGroup MacOSConfigurationProfilesDataSubsetComputerGroupItem `xml:"computer_group,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetComputerGroupItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetLimitations struct {
	Users           []MacOSConfigurationProfilesDataSubsetUser           `xml:"users,omitempty"`
	UserGroups      []MacOSConfigurationProfilesDataSubsetUserGroup      `xml:"user_groups,omitempty"`
	NetworkSegments []MacOSConfigurationProfilesDataSubsetNetworkSegment `xml:"network_segments,omitempty"`
	IBeacons        []MacOSConfigurationProfilesDataSubsetIBeacon        `xml:"ibeacons,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetUser struct {
	User MacOSConfigurationProfilesDataSubsetUserItem `xml:"user,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetUserItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetUserGroup struct {
	UserGroup MacOSConfigurationProfilesDataSubsetUserGroupItem `xml:"user_group,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetUserGroupItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetNetworkSegment struct {
	NetworkSegment MacOSConfigurationProfilesDataSubsetNetworkSegmentItem `xml:"network_segment,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetNetworkSegmentItem struct {
	ID   int    `xml:"id,omitempty"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetIBeacon struct {
	IBeacon MacOSConfigurationProfilesDataSubsetIBeaconItem `xml:"ibeacon,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetIBeaconItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetJSSUser struct {
	JSSUser MacOSConfigurationProfilesDataSubsetJSSUserItem `xml:"jss_user,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetJSSUserItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetJSSUserGroup struct {
	JSSUserGroup MacOSConfigurationProfilesDataSubsetJSSUserGroupItem `xml:"jss_user_group,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetJSSUserGroupItem struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetExclusions struct {
	Computers       []MacOSConfigurationProfilesDataSubsetComputer       `xml:"computers,omitempty"`
	Buildings       []MacOSConfigurationProfilesDataSubsetBuilding       `xml:"buildings,omitempty"`
	Departments     []MacOSConfigurationProfilesDataSubsetDepartment     `xml:"departments,omitempty"`
	ComputerGroups  []MacOSConfigurationProfilesDataSubsetComputerGroup  `xml:"computer_groups,omitempty"`
	Users           []MacOSConfigurationProfilesDataSubsetUser           `xml:"users,omitempty"`
	UserGroups      []MacOSConfigurationProfilesDataSubsetUserGroup      `xml:"user_groups,omitempty"`
	NetworkSegments []MacOSConfigurationProfilesDataSubsetNetworkSegment `xml:"network_segments,omitempty"`
	IBeacons        []MacOSConfigurationProfilesDataSubsetIBeacon        `xml:"ibeacons,omitempty"`
	JSSUsers        []MacOSConfigurationProfilesDataSubsetJSSUser        `xml:"jss_users,omitempty"`
	JSSUserGroups   []MacOSConfigurationProfilesDataSubsetJSSUserGroup   `xml:"jss_user_groups,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetSelfService struct {
	InstallButtonText           string                                                    `xml:"install_button_text,omitempty"`
	SelfServiceDescription      string                                                    `xml:"self_service_description,omitempty"`
	ForceUsersToViewDescription bool                                                      `xml:"force_users_to_view_description,omitempty"`
	SelfServiceIcon             MacOSConfigurationProfilesDataSubsetSelfServiceIcon       `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                                                      `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories       MacOSConfigurationProfilesDataSubsetSelfServiceCategories `xml:"self_service_categories,omitempty"`
	Notification                string                                                    `xml:"notification,omitempty"`
	NotificationSubject         string                                                    `xml:"notification_subject,omitempty"`
	NotificationMessage         string                                                    `xml:"notification_message,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetSelfServiceIcon struct {
	ID   int    `xml:"id,omitempty"`
	URI  string `xml:"uri,omitempty"`
	Data string `xml:"data,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetSelfServiceCategories struct {
	Category MacOSConfigurationProfilesDataSubsetSelfServiceCategory `xml:"category,omitempty"`
}

type MacOSConfigurationProfilesDataSubsetSelfServiceCategory struct {
	ID        int    `xml:"id,omitempty"`
	Name      string `xml:"name,omitempty"`
	DisplayIn bool   `xml:"display_in,omitempty"`
	FeatureIn bool   `xml:"feature_in,omitempty"`
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

// GetMacOSConfigurationProfileByNameByID retrieves the details of a macOS Configuration Profile by its name.
func (c *Client) GetMacOSConfigurationProfileByNameByID(name string) (*ResponseMacOSConfigurationProfiles, error) {
	// Fetch all macOS Configuration Profiles
	profilesList, err := c.GetMacOSConfigurationProfiles()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS Configuration Profiles: %v", err)
	}

	// Search for the profile with the given name
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

	// Fetch the full details of the profile using its ID
	detailedProfile, err := c.GetMacOSConfigurationProfileByID(profileID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch macOS Configuration Profile by ID: %v", err)
	}

	return detailedProfile, nil
}

// CreateMacOSConfigurationProfile creates a new macOS Configuration Profile on the Jamf Pro server and returns the profile with its ID updated.
// It sends a POST request to the Jamf Pro server with the profile details and expects a response with the ID of the newly created profile.
// CreateMacOSConfigurationProfile creates a new macOS Configuration Profile on the Jamf Pro server and returns the ID of the newly created profile.
func (c *Client) CreateMacOSConfigurationProfile(profile *ResponseMacOSConfigurationProfiles) (int, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMacOSConfigurationProfiles)

	// Set default values for site and category if not included within request
	if profile.General.Site.ID == 0 && profile.General.Site.Name == "" {
		profile.General.Site = MacOSConfigurationProfilesDataSubsetSite{ID: -1, Name: "None"}
	}
	if profile.General.Category.ID == 0 && profile.General.Category.Name == "" {
		profile.General.Category = MacOSConfigurationProfilesDataSubsetCategory{ID: -1, Name: "No category assigned"}
	}

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfiles
	}{
		ResponseMacOSConfigurationProfiles: profile,
	}

	// Use ResponseMacOSConfigurationProfileCreationAndUpdate struct to handle the API response
	var response ResponseMacOSConfigurationProfileCreationAndUpdate

	// Send the request and capture the response
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to create macOS Configuration Profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Return the ID from the response
	return response.ID, nil
}

// UpdateMacOSConfigurationProfileByID updates an existing macOS Configuration Profile by its ID on the Jamf Pro server
// and returns the ID of the updated profile.
func (c *Client) UpdateMacOSConfigurationProfileByID(id int, profile *ResponseMacOSConfigurationProfiles) (int, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMacOSConfigurationProfiles, id)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfiles
	}{
		ResponseMacOSConfigurationProfiles: profile,
	}

	// Use ResponseMacOSConfigurationProfileCreationAndUpdate struct to handle the API response
	var response ResponseMacOSConfigurationProfileCreationAndUpdate

	// Send the request and capture the response
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to update macOS Configuration Profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Return the ID from the response
	return response.ID, nil
}

// UpdateMacOSConfigurationProfileByName updates an existing macOS Configuration Profile by its name on the Jamf Pro server
// and returns the ID of the updated profile.
func (c *Client) UpdateMacOSConfigurationProfileByName(name string, profile *ResponseMacOSConfigurationProfiles) (int, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMacOSConfigurationProfiles, name)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"os_x_configuration_profile"`
		*ResponseMacOSConfigurationProfiles
	}{
		ResponseMacOSConfigurationProfiles: profile,
	}

	// Use ResponseMacOSConfigurationProfileCreationAndUpdate struct to handle the API response
	var response ResponseMacOSConfigurationProfileCreationAndUpdate

	// Send the request and capture the response
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return 0, fmt.Errorf("failed to update macOS Configuration Profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Return the ID from the response
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
