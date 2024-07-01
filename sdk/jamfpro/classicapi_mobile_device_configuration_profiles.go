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

// Responses

// ResponseMacOSConfigurationProfileCreation represents the response structure for a new macOS configuration profile.
type ResponseMobileDeviceConfigurationProfileCreateAndUpdate struct {
	ID int `xml:"id"`
}

// Resource

// ResourceMobileDeviceConfigurationProfile represents the detailed structure of a single mobile device configuration profile.
type ResourceMobileDeviceConfigurationProfile struct {
	General     MobileDeviceConfigurationProfileSubsetGeneral     `xml:"general"`
	Scope       MobileDeviceConfigurationProfileSubsetScope       `xml:"scope,omitempty"`
	SelfService MobileDeviceConfigurationProfileSubsetSelfService `xml:"self_service,omitempty"`
}

// Subsets and Containers

type MobileDeviceConfigurationProfileSubsetGeneral struct {
	ID                            int                     `xml:"id"`
	Name                          string                  `xml:"name"`
	Description                   string                  `xml:"description,omitempty"`
	Level                         string                  `xml:"level,omitempty"`
	Site                          *SharedResourceSite     `xml:"site"`
	Category                      *SharedResourceCategory `xml:"category"`
	UUID                          string                  `xml:"uuid,omitempty"`
	DeploymentMethod              string                  `xml:"deployment_method,omitempty"`
	RedeployOnUpdate              string                  `xml:"redeploy_on_update,omitempty"`
	RedeployDaysBeforeCertExpires int                     `xml:"redeploy_Dayss_before_certificate_expires,omitempty"`
	Payloads                      string                  `xml:"payloads,omitempty"`
}

type MobileDeviceConfigurationProfileSubsetScope struct {
	AllMobileDevices   bool                                                 `xml:"all_mobile_devices,omitempty"`
	AllJSSUsers        bool                                                 `xml:"all_jss_users,omitempty"`
	MobileDevices      []MobileDeviceConfigurationProfileSubsetMobileDevice `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceGroups []MobileDeviceConfigurationProfileSubsetScopeEntity  `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	JSSUsers           []MobileDeviceConfigurationProfileSubsetScopeEntity  `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []MobileDeviceConfigurationProfileSubsetScopeEntity  `xml:"jss_user_groups>user_group,omitempty"`
	Buildings          []MobileDeviceConfigurationProfileSubsetScopeEntity  `xml:"buildings>building,omitempty"`
	Departments        []MobileDeviceConfigurationProfileSubsetScopeEntity  `xml:"departments>department,omitempty"`
	Limitations        MobileDeviceConfigurationProfileSubsetLimitation     `xml:"limitations,omitempty"`
	Exclusions         MobileDeviceConfigurationProfileSubsetExclusion      `xml:"exclusions,omitempty"`
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
	NetworkSegments []MobileDeviceConfigurationProfileSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	Users           []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"users>user,omitempty"`
	UserGroups      []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"user_groups>user_group,omitempty"`
	Ibeacons        []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"ibeacons>ibeacon,omitempty"`
}

type MobileDeviceConfigurationProfileSubsetExclusion struct {
	MobileDevices      []MobileDeviceConfigurationProfileSubsetMobileDevice   `xml:"mobile_devices>mobile_device,omitempty"`
	MobileDeviceGroups []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	Users              []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"users>user,omitempty"`
	UserGroups         []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"user_groups>user_group,omitempty"`
	Buildings          []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"buildings>building,omitempty"`
	Departments        []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"departments>department,omitempty"`
	NetworkSegments    []MobileDeviceConfigurationProfileSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
	JSSUsers           []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"jss_user_groups>user_group,omitempty"`
	IBeacons           []MobileDeviceConfigurationProfileSubsetScopeEntity    `xml:"ibeacons>ibeacon,omitempty"`
}

// Mobile Device

type MobileDeviceConfigurationProfileSubsetMobileDevice struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name,omitempty"`
	UDID           string `xml:"udid,omitempty"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
}

// Entity

// Generic Entity struct for common use
type MobileDeviceConfigurationProfileSubsetScopeEntity struct {
	ID   int    `xml:"id"`
	Name string `xml:"name,omitempty"`
}

// Specific struct for NetworkSegment due to its unique attribute 'UID'
type MobileDeviceConfigurationProfileSubsetNetworkSegment struct {
	MobileDeviceConfigurationProfileSubsetScopeEntity
	UID string `xml:"uid,omitempty"`
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
func (c *Client) CreateMobileDeviceConfigurationProfile(profile *ResourceMobileDeviceConfigurationProfile) (*ResponseMobileDeviceConfigurationProfileCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceConfigurationProfiles)

	requestBody := struct {
		XMLName xml.Name `xml:"configuration_profile"`
		*ResourceMobileDeviceConfigurationProfile
	}{
		ResourceMobileDeviceConfigurationProfile: profile,
	}

	var responseProfile ResponseMobileDeviceConfigurationProfileCreateAndUpdate
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
func (c *Client) UpdateMobileDeviceConfigurationProfileByID(id int, profile *ResourceMobileDeviceConfigurationProfile) (*ResponseMobileDeviceConfigurationProfileCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceConfigurationProfiles, id)

	requestBody := struct {
		XMLName xml.Name `xml:"configuration_profile"`
		*ResourceMobileDeviceConfigurationProfile
	}{
		ResourceMobileDeviceConfigurationProfile: profile,
	}

	var responseProfile ResponseMobileDeviceConfigurationProfileCreateAndUpdate
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
func (c *Client) UpdateMobileDeviceConfigurationProfileByName(name string, profile *ResourceMobileDeviceConfigurationProfile) (*ResponseMobileDeviceConfigurationProfileCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceConfigurationProfiles, name)

	requestBody := struct {
		XMLName xml.Name `xml:"configuration_profile"`
		*ResourceMobileDeviceConfigurationProfile
	}{
		ResourceMobileDeviceConfigurationProfile: profile,
	}

	var responseProfile ResponseMobileDeviceConfigurationProfileCreateAndUpdate
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
