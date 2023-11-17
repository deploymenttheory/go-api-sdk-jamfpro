// classicapi_mobile_device_configuration_profiles.go
// Jamf Pro Classic Api - Mobile Device Configuration Profiles
// API reference: https://developer.jamf.com/jamf-pro/reference/mobiledeviceconfigurationprofiles
// Jamf Pro Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMobileDeviceConfigurationProfiles = "/JSSResource/mobiledeviceconfigurationprofiles"

// ResponseMobileDeviceConfigurationProfilesList represents the response for a list of mobile device configuration profiles.
type ResponseMobileDeviceConfigurationProfilesList struct {
	ConfigurationProfiles []MobileDeviceConfigurationProfiles `xml:"configuration_profile"`
}

// ConfigurationProfile represents a single mobile device configuration profile.
type MobileDeviceConfigurationProfiles struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResponseMobileDeviceConfigurationProfile represents the detailed structure of a single mobile device configuration profile.
type ResponseMobileDeviceConfigurationProfile struct {
	General     MobileDeviceConfigurationProfileGeneral     `xml:"general"`
	Scope       MobileDeviceConfigurationProfileScope       `xml:"scope"`
	SelfService MobileDeviceConfigurationProfileSelfService `xml:"self_service"`
}

// ConfigurationProfileGeneral contains general information about the configuration profile.
type MobileDeviceConfigurationProfileGeneral struct {
	ID                            int                                      `xml:"id"`
	Name                          string                                   `xml:"name"`
	Description                   string                                   `xml:"description,omitempty"`
	Level                         string                                   `xml:"level,omitempty"`
	Site                          MobileDeviceConfigurationProfileSite     `xml:"site"`
	Category                      MobileDeviceConfigurationProfileCategory `xml:"category"`
	UUID                          string                                   `xml:"uuid,omitempty"`
	DeploymentMethod              string                                   `xml:"deployment_method,omitempty"`
	RedeployOnUpdate              string                                   `xml:"redeploy_on_update,omitempty"`
	RedeployDaysBeforeCertExpires int                                      `xml:"redeploy_Dayss_before_certificate_expires,omitempty"`
	Payloads                      string                                   `xml:"payloads,omitempty"`
}

type MobileDeviceConfigurationProfileSite struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MobileDeviceConfigurationProfileCategory struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ConfigurationProfileScope defines the scope of the configuration profile.
type MobileDeviceConfigurationProfileScope struct {
	AllMobileDevices   bool                                                `xml:"all_mobile_devices,omitempty"`
	AllJSSUsers        bool                                                `xml:"all_jss_users,omitempty"`
	MobileDevices      []MobileDeviceConfigurationProfileMobileDevice      `xml:"mobile_devices>mobile_device,omitempty"`
	Buildings          []MobileDeviceConfigurationProfileBuilding          `xml:"buildings>building,omitempty"`
	Departments        []MobileDeviceConfigurationProfileDepartment        `xml:"departments>department,omitempty"`
	MobileDeviceGroups []MobileDeviceConfigurationProfileMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	JSSUsers           []MobileDeviceConfigurationProfileJSSUser           `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []MobileDeviceConfigurationProfileJSSUserGroup      `xml:"jss_user_groups>user_group,omitempty"`
	Limitations        MobileDeviceConfigurationProfileLimitation          `xml:"limitations,omitempty"`
	Exclusions         MobileDeviceConfigurationProfileExclusion           `xml:"exclusions,omitempty"`
	// Continue with nested structs for each of the above arrays
}

type MobileDeviceConfigurationProfileMobileDevice struct {
	MobileDevice MobileDeviceConfigurationProfileDevice `xml:"mobile_device"`
}

type MobileDeviceConfigurationProfileDevice struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name"`
	UDID           string `xml:"udid,omitempty"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
}

type MobileDeviceConfigurationProfileBuilding struct {
	Building MobileDeviceConfigurationProfileBuildingDetail `xml:"building"`
}

type MobileDeviceConfigurationProfileBuildingDetail struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MobileDeviceConfigurationProfileDepartment struct {
	Department MobileDeviceConfigurationProfileDepartmentDetail `xml:"department"`
}

type MobileDeviceConfigurationProfileDepartmentDetail struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MobileDeviceConfigurationProfileMobileDeviceGroup struct {
	MobileDeviceGroup MobileDeviceConfigurationProfileDeviceGroup `xml:"mobile_device_group"`
}

type MobileDeviceConfigurationProfileDeviceGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MobileDeviceConfigurationProfileJSSUser struct {
	User MobileDeviceConfigurationProfileUser `xml:"user"`
}

type MobileDeviceConfigurationProfileUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MobileDeviceConfigurationProfileJSSUserGroup struct {
	UserGroup MobileDeviceConfigurationProfileUserGroup `xml:"user_group"`
}

type MobileDeviceConfigurationProfileUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MobileDeviceConfigurationProfileLimitation struct {
	Users           []MobileDeviceConfigurationProfileUser           `xml:"users>user,omitempty"`
	UserGroups      []MobileDeviceConfigurationProfileUserGroup      `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []MobileDeviceConfigurationProfileNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	Ibeacons        []MobileDeviceConfigurationProfileIbeacon        `xml:"ibeacons>ibeacon,omitempty"`
}

type MobileDeviceConfigurationProfileNetworkSegment struct {
	ID   int    `xml:"id"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name"`
}

type MobileDeviceConfigurationProfileIbeacon struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MobileDeviceConfigurationProfileExclusion struct {
	MobileDevices      []MobileDeviceConfigurationProfileMobileDevice      `xml:"mobile_devices>mobile_device,omitempty"`
	Buildings          []MobileDeviceConfigurationProfileBuilding          `xml:"buildings>building,omitempty"`
	Departments        []MobileDeviceConfigurationProfileDepartment        `xml:"departments>department,omitempty"`
	MobileDeviceGroups []MobileDeviceConfigurationProfileMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	Users              []MobileDeviceConfigurationProfileUser              `xml:"users>user,omitempty"`
	UserGroups         []MobileDeviceConfigurationProfileUserGroup         `xml:"user_groups>user_group,omitempty"`
	NetworkSegments    []MobileDeviceConfigurationProfileNetworkSegment    `xml:"network_segments>network_segment,omitempty"`
	Ibeacons           []MobileDeviceConfigurationProfileIbeacon           `xml:"ibeacons>ibeacon,omitempty"`
	JSSUsers           []MobileDeviceConfigurationProfileJSSUser           `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []MobileDeviceConfigurationProfileJSSUserGroup      `xml:"jss_user_groups>user_group,omitempty"`
}

type MobileDeviceConfigurationProfileSelfService struct {
	SelfServiceDescription string                                                `xml:"self_service_description,omitempty"`
	SecurityName           MobileDeviceConfigurationProfileSecurityName          `xml:"security_name,omitempty"`
	SelfServiceIcon        MobileDeviceConfigurationProfileSelfServiceIcon       `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage      bool                                                  `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories  []MobileDeviceConfigurationProfileSelfServiceCategory `xml:"self_service_categories>category,omitempty"`
}

type MobileDeviceConfigurationProfileSecurityName struct {
	RemovalDisallowed string `xml:"removal_disallowed,omitempty"`
}

type MobileDeviceConfigurationProfileSelfServiceIcon struct {
	Filename string `xml:"filename,omitempty"`
	URI      string `xml:"uri,omitempty"`
	Data     string `xml:"data,omitempty"`
}

type MobileDeviceConfigurationProfileSelfServiceCategory struct {
	Category MobileDeviceConfigurationProfileCategoryDetail `xml:"category"`
}

type MobileDeviceConfigurationProfileCategoryDetail struct {
	ID       int    `xml:"id"`
	Name     string `xml:"name"`
	Priority int    `xml:"priority,omitempty"`
}

// GetMobileDeviceConfigurationProfiles retrieves a serialized list of mobile device configuration profiles.
func (c *Client) GetMobileDeviceConfigurationProfiles() (*ResponseMobileDeviceConfigurationProfilesList, error) {
	endpoint := uriMobileDeviceConfigurationProfiles

	var profiles ResponseMobileDeviceConfigurationProfilesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device configuration profiles: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profiles, nil
}

// GetMobileDeviceConfigurationProfileByID fetches a specific mobile device configuration profile by its ID.
func (c *Client) GetMobileDeviceConfigurationProfileByID(id int) (*ResponseMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceConfigurationProfiles, id)

	var profile ResponseMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device configuration profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceConfigurationProfileByName fetches a specific mobile device configuration profile by its name.
func (c *Client) GetMobileDeviceConfigurationProfileByName(name string) (*ResponseMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceConfigurationProfiles, name)

	var profile ResponseMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device configuration profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceConfigurationProfileByIDBySubset fetches a specific mobile device configuration profile by its ID and a specified subset.
func (c *Client) GetMobileDeviceConfigurationProfileByIDBySubset(id int, subset string) (*ResponseMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriMobileDeviceConfigurationProfiles, id, subset)

	var profile ResponseMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device configuration profile by ID and subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceConfigurationProfileByNameBySubset fetches a specific mobile device configuration profile by its name and a specified subset.
func (c *Client) GetMobileDeviceConfigurationProfileByNameBySubset(name string, subset string) (*ResponseMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriMobileDeviceConfigurationProfiles, name, subset)

	var profile ResponseMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device configuration profile by name and subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// CreateMobileDeviceConfigurationProfile creates a new mobile device configuration profile on the Jamf Pro server.
func (c *Client) CreateMobileDeviceConfigurationProfile(profile *ResponseMobileDeviceConfigurationProfile) (*ResponseMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceConfigurationProfiles)

	// Set default values for site and category if not included within request
	if profile.General.Site.ID == 0 && profile.General.Site.Name == "" {
		profile.General.Site = MobileDeviceConfigurationProfileSite{
			ID:   -1,
			Name: "None",
		}
	}
	if profile.General.Category.ID == 0 && profile.General.Category.Name == "" {
		profile.General.Category = MobileDeviceConfigurationProfileCategory{
			ID:   -1,
			Name: "No category assigned",
		}
	}

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"configuration_profile"`
		*ResponseMobileDeviceConfigurationProfile
	}{
		ResponseMobileDeviceConfigurationProfile: profile,
	}

	var responseProfile ResponseMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to create mobile device configuration profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// UpdateMobileDeviceConfigurationProfileByID updates a mobile device configuration profile by its ID on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceConfigurationProfileByID(id int, profile *ResponseMobileDeviceConfigurationProfile) (*ResponseMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceConfigurationProfiles, id)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"configuration_profile"`
		*ResponseMobileDeviceConfigurationProfile
	}{
		ResponseMobileDeviceConfigurationProfile: profile,
	}

	var responseProfile ResponseMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device configuration profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// UpdateMobileDeviceConfigurationProfileByName updates a mobile device configuration profile by its name on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceConfigurationProfileByName(name string, profile *ResponseMobileDeviceConfigurationProfile) (*ResponseMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceConfigurationProfiles, name)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"configuration_profile"`
		*ResponseMobileDeviceConfigurationProfile
	}{
		ResponseMobileDeviceConfigurationProfile: profile,
	}

	var responseProfile ResponseMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device configuration profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// DeleteMobileDeviceConfigurationProfileByID deletes a mobile device configuration profile by its ID from the Jamf Pro server.
func (c *Client) DeleteMobileDeviceConfigurationProfileByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceConfigurationProfiles, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device configuration profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceConfigurationProfileByName deletes a mobile device configuration profile by its name from the Jamf Pro server.
func (c *Client) DeleteMobileDeviceConfigurationProfileByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceConfigurationProfiles, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device configuration profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
