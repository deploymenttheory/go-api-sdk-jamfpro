// classicapi_mobile_device_configuration_profiles.go
// Jamf Pro Classic Api - Mobile Device Configuration Profiles
// API reference: https://developer.jamf.com/jamf-pro/reference/mobiledeviceconfigurationprofiles
// Jamf Pro Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
- SharedResourceCategory
- SharedResourceSelfServiceIcon
- SharedResourceSelfServiceCategories
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMobileDeviceConfigurationProfiles = "/JSSResource/mobiledeviceconfigurationprofiles"

// List

// ResponseMobileDeviceConfigurationProfilesList represents the response for a list of mobile device configuration profiles.
type ResponseMobileDeviceConfigurationProfilesList struct {
	ConfigurationProfiles []MobileDeviceConfigurationProfilesListItem `xml:"configuration_profile"`
}

type MobileDeviceConfigurationProfilesListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// ResponseMobileDeviceConfigurationProfile represents the detailed structure of a single mobile device configuration profile.
type ResourceMobileDeviceConfigurationProfile struct {
	General     MobileDeviceConfigurationProfileSubsetGeneral     `xml:"general"`
	Scope       MobileDeviceConfigurationProfileSubsetScope       `xml:"scope,omitempty"`
	SelfService MobileDeviceConfigurationProfileSubsetSelfService `xml:"self_service,omitempty"`
}

// Subsets and Containers

type MobileDeviceConfigurationProfileSubsetGeneral struct {
	ID                            int                    `xml:"id"`
	Name                          string                 `xml:"name"`
	Description                   string                 `xml:"description,omitempty"`
	Level                         string                 `xml:"level,omitempty"`
	Site                          SharedResourceSite     `xml:"site"`
	Category                      SharedResourceCategory `xml:"category"`
	UUID                          string                 `xml:"uuid,omitempty"`
	DeploymentMethod              string                 `xml:"deployment_method,omitempty"`
	RedeployOnUpdate              string                 `xml:"redeploy_on_update,omitempty"`
	RedeployDaysBeforeCertExpires int                    `xml:"redeploy_Dayss_before_certificate_expires,omitempty"`
	Payloads                      string                 `xml:"payloads,omitempty"`
}

type MobileDeviceConfigurationProfileSubsetScope struct {
	AllMobileDevices   bool                                                      `xml:"all_mobile_devices,omitempty"`
	AllJSSUsers        bool                                                      `xml:"all_jss_users,omitempty"`
	MobileDevices      []MobileDeviceConfigurationProfileSubsetMobileDevice      `xml:"mobile_devices>mobile_device,omitempty"`
	Buildings          []MobileDeviceConfigurationProfileSubsetBuilding          `xml:"buildings>building,omitempty"`
	Departments        []MobileDeviceConfigurationProfileSubsetDepartment        `xml:"departments>department,omitempty"`
	MobileDeviceGroups []MobileDeviceConfigurationProfileSubsetMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	JSSUsers           []MobileDeviceConfigurationProfileSubsetJSSUser           `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []MobileDeviceConfigurationProfileSubsetJSSUserGroup      `xml:"jss_user_groups>user_group,omitempty"`
	Limitations        MobileDeviceConfigurationProfileSubsetLimitation          `xml:"limitations,omitempty"`
	Exclusions         MobileDeviceConfigurationProfileSubsetExclusion           `xml:"exclusions,omitempty"`
}

type MobileDeviceConfigurationProfileSubsetSelfService struct {
	SelfServiceDescription string                                                        `xml:"self_service_description,omitempty"`
	SecurityName           MobileDeviceConfigurationProfileSubsetSelfServiceSecurityName `xml:"security_name,omitempty"`
	SelfServiceIcon        SharedResourceSelfServiceIcon                                 `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage      bool                                                          `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories  []SharedResourceSelfServiceCategories                         `xml:"self_service_categories>category,omitempty"`
}

type MobileDeviceConfigurationProfileSubsetSelfServiceSecurityName struct {
	RemovalDisallowed string `xml:"removal_disallowed,omitempty"`
}

// Shared in Resource

type MobileDeviceConfigurationProfileSubsetLimitation struct {
	Users           []MobileDeviceConfigurationProfileSubsetUser           `xml:"users>user,omitempty"`
	UserGroups      []MobileDeviceConfigurationProfileSubsetUserGroup      `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []MobileDeviceConfigurationProfileSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	Ibeacons        []MobileDeviceConfigurationProfileSubsetIbeacon        `xml:"ibeacons>ibeacon,omitempty"`
}

type MobileDeviceConfigurationProfileSubsetExclusion struct {
	MobileDevices      []MobileDeviceConfigurationProfileContainerMobileDevice      `xml:"mobile_devices>mobile_device,omitempty"`
	Buildings          []MobileDeviceConfigurationProfileContainerBuilding          `xml:"buildings>building,omitempty"`
	Departments        []MobileDeviceConfigurationProfileContainerDepartment        `xml:"departments>department,omitempty"`
	MobileDeviceGroups []MobileDeviceConfigurationProfileContainerMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	Users              []MobileDeviceConfigurationProfileSubsetUser                 `xml:"users>user,omitempty"`
	UserGroups         []MobileDeviceConfigurationProfileSubsetUserGroup            `xml:"user_groups>user_group,omitempty"`
	NetworkSegments    []MobileDeviceConfigurationProfileSubsetNetworkSegment       `xml:"network_segments>network_segment,omitempty"`
	Ibeacons           []MobileDeviceConfigurationProfileSubsetIbeacon              `xml:"ibeacons>ibeacon,omitempty"`
	JSSUsers           []MobileDeviceConfigurationProfileSubsetJSSUser              `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []MobileDeviceConfigurationProfileSubsetJSSUserGroup         `xml:"jss_user_groups>user_group,omitempty"`
}

// Mobile Device

type MobileDeviceConfigurationProfileContainerMobileDevice struct {
	MobileDevice MobileDeviceConfigurationProfileSubsetMobileDevice `xml:"mobile_device"`
}

type MobileDeviceConfigurationProfileSubsetMobileDevice struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name"`
	UDID           string `xml:"udid,omitempty"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
}

// Building

type MobileDeviceConfigurationProfileContainerBuilding struct {
	Building MobileDeviceConfigurationProfileSubsetBuilding `xml:"building"`
}

type MobileDeviceConfigurationProfileSubsetBuilding struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Department

type MobileDeviceConfigurationProfileContainerDepartment struct {
	Department MobileDeviceConfigurationProfileSubsetDepartment `xml:"department"`
}

type MobileDeviceConfigurationProfileSubsetDepartment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Mobile Device Group

type MobileDeviceConfigurationProfileContainerMobileDeviceGroup struct {
	MobileDeviceGroup MobileDeviceConfigurationProfileSubsetMobileDeviceGroup `xml:"mobile_device_group"`
}

type MobileDeviceConfigurationProfileSubsetMobileDeviceGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// JSS User

type MobileDeviceConfigurationProfileSubsetJSSUser struct {
	User MobileDeviceConfigurationProfileSubsetUser `xml:"user"`
}

// User

type MobileDeviceConfigurationProfileSubsetUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// JSS User group

type MobileDeviceConfigurationProfileSubsetJSSUserGroup struct {
	UserGroup MobileDeviceConfigurationProfileSubsetUserGroup `xml:"user_group"`
}

// User group

type MobileDeviceConfigurationProfileSubsetUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Network Segment

type MobileDeviceConfigurationProfileSubsetNetworkSegment struct {
	ID   int    `xml:"id"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name"`
}

// IBeacon

type MobileDeviceConfigurationProfileSubsetIbeacon struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// CRUD

// GetMobileDeviceConfigurationProfiles retrieves a serialized list of mobile device configuration profiles.
func (c *Client) GetMobileDeviceConfigurationProfiles() (*ResponseMobileDeviceConfigurationProfilesList, error) {
	endpoint := uriMobileDeviceConfigurationProfiles

	var profiles ResponseMobileDeviceConfigurationProfilesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "mobile device configuration profiles", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profiles, nil
}

// GetMobileDeviceConfigurationProfileByID fetches a specific mobile device configuration profile by its ID.
func (c *Client) GetMobileDeviceConfigurationProfileByID(id int) (*ResourceMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceConfigurationProfiles, id)

	var profile ResourceMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device configuration profile", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceConfigurationProfileByName fetches a specific mobile device configuration profile by its name.
func (c *Client) GetMobileDeviceConfigurationProfileByName(name string) (*ResourceMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceConfigurationProfiles, name)

	var profile ResourceMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device configuration profile", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceConfigurationProfileByIDBySubset fetches a specific mobile device configuration profile by its ID and a specified subset.
func (c *Client) GetMobileDeviceConfigurationProfileByIDWithSubset(id int, subset string) (*ResourceMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriMobileDeviceConfigurationProfiles, id, subset)

	var profile ResourceMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device configuration profile with data subset", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceConfigurationProfileByNameBySubset fetches a specific mobile device configuration profile by its name and a specified subset.
func (c *Client) GetMobileDeviceConfigurationProfileByNameWithSubset(name string, subset string) (*ResourceMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriMobileDeviceConfigurationProfiles, name, subset)

	var profile ResourceMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device configuration profile with data subset", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// CreateMobileDeviceConfigurationProfile creates a new mobile device configuration profile on the Jamf Pro server.
func (c *Client) CreateMobileDeviceConfigurationProfile(profile *ResourceMobileDeviceConfigurationProfile) (*ResourceMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceConfigurationProfiles)

	// Set default values for site and category if not included within request
	if profile.General.Site.ID == 0 && profile.General.Site.Name == "" {
		profile.General.Site.ID = -1
		profile.General.Site.Name = "none"
	}
	if profile.General.Category.ID == 0 && profile.General.Category.Name == "" {
		profile.General.Category.ID = -1
		profile.General.Category.Name = "no category"

	}

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"configuration_profile"`
		*ResourceMobileDeviceConfigurationProfile
	}{
		ResourceMobileDeviceConfigurationProfile: profile,
	}

	var responseProfile ResourceMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "mobile device configuration profile", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// UpdateMobileDeviceConfigurationProfileByID updates a mobile device configuration profile by its ID on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceConfigurationProfileByID(id int, profile *ResourceMobileDeviceConfigurationProfile) (*ResourceMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceConfigurationProfiles, id)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"configuration_profile"`
		*ResourceMobileDeviceConfigurationProfile
	}{
		ResourceMobileDeviceConfigurationProfile: profile,
	}

	var responseProfile ResourceMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "mobile device configuration profile", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// UpdateMobileDeviceConfigurationProfileByName updates a mobile device configuration profile by its name on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceConfigurationProfileByName(name string, profile *ResourceMobileDeviceConfigurationProfile) (*ResourceMobileDeviceConfigurationProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceConfigurationProfiles, name)

	requestBody := struct {
		XMLName xml.Name `xml:"configuration_profile"`
		*ResourceMobileDeviceConfigurationProfile
	}{
		ResourceMobileDeviceConfigurationProfile: profile,
	}

	var responseProfile ResourceMobileDeviceConfigurationProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "mobile device configuration profile", name, err)
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
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device configuration profile", id, err)
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
		return fmt.Errorf(errMsgFailedDeleteByName, "mobile device configuration profile", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
