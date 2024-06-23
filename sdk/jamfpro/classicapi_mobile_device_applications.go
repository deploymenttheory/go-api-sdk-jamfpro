// classicapi_mobile_device_applications.go
// Jamf Pro Classic Api - Mobile Device Applications
// api reference: https://developer.jamf.com/jamf-pro/reference/mobiledeviceapplications
// Jamf Pro Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
- SharedResourceCategory
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMobileDeviceApplications = "/JSSResource/mobiledeviceapplications"

// List

// ResponseMobileDeviceApplicationsList represents the response for a list of mobile device applications.
type ResponseMobileDeviceApplicationsList struct {
	MobileDeviceApplications []MobileDeviceApplicationsListItem `xml:"mobile_device_application"`
}

type MobileDeviceApplicationsListItem struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	DisplayName string `xml:"display_name"`
	BundleID    string `xml:"bundle_id"`
	Version     string `xml:"version"`
	InternalApp bool   `xml:"internal_app"`
}

// Resource

// ResourceMobileDeviceApplication represents the detailed structure of a single mobile device application.
type ResourceMobileDeviceApplication struct {
	General MobileDeviceApplicationSubsetGeneral `xml:"general"`
}

// Subsets and Containers

type MobileDeviceApplicationSubsetGeneral struct {
	ID                               int                                                  `xml:"id,omitempty"`
	Name                             string                                               `xml:"name"`
	DisplayName                      string                                               `xml:"display_name"`
	Description                      string                                               `xml:"description,omitempty"`
	BundleID                         string                                               `xml:"bundle_id"`
	Version                          string                                               `xml:"version"`
	InternalApp                      bool                                                 `xml:"internal_app,omitempty"`
	OsType                           string                                               `xml:"os_type,omitempty"`
	Category                         SharedResourceCategory                               `xml:"category"`
	IPA                              MobileDeviceApplicationSubsetGeneralIPA              `xml:"ipa,omitempty"`
	Icon                             MobileDeviceApplicationSubsetIcon                    `xml:"icon"`
	ProvisioningProfile              int                                                  `xml:"mobile_device_provisioning_profile,omitempty"`
	ITunesStoreURL                   string                                               `xml:"itunes_store_url,omitempty"`
	MakeAvailableAfterInstall        bool                                                 `xml:"make_available_after_install,omitempty"`
	ITunesCountryRegion              string                                               `xml:"itunes_country_region,omitempty"`
	ITunesSyncTime                   int                                                  `xml:"itunes_sync_time,omitempty"`
	DeploymentType                   string                                               `xml:"deployment_type,omitempty"`
	DeployAutomatically              bool                                                 `xml:"deploy_automatically,omitempty"`
	DeployAsManagedApp               bool                                                 `xml:"deploy_as_managed_app,omitempty"`
	RemoveAppWhenMDMProfileIsRemoved bool                                                 `xml:"remove_app_when_mdm_profile_is_removed,omitempty"`
	PreventBackupOfAppData           bool                                                 `xml:"prevent_backup_of_app_data,omitempty"`
	KeepDescriptionAndIconUpToDate   bool                                                 `xml:"keep_description_and_icon_up_to_date,omitempty"`
	Free                             bool                                                 `xml:"free,omitempty"`
	TakeOverManagement               bool                                                 `xml:"take_over_management,omitempty"`
	HostExternally                   bool                                                 `xml:"host_externally,omitempty"`
	ExternalURL                      string                                               `xml:"external_url,omitempty"`
	Site                             *SharedResourceSite                                  `xml:"site"`
	Scope                            MobileDeviceApplicationSubsetScope                   `xml:"scope"`
	SelfService                      MobileDeviceApplicationSubsetGeneralSelfService      `xml:"self_service"`
	VPP                              MobileDeviceApplicationSubsetGeneralVPP              `xml:"vpp,omitempty"`
	AppConfiguration                 MobileDeviceApplicationSubsetGeneralAppConfiguration `xml:"app_configuration,omitempty"`
}

type MobileDeviceApplicationSubsetGeneralIPA struct {
	Name string `xml:"name,omitempty"`
	URI  string `xml:"uri,omitempty"`
	Data string `xml:"data,omitempty"`
}

type MobileDeviceApplicationSubsetGeneralSelfService struct {
	SelfServiceDescription string                              `xml:"self_service_description,omitempty"`
	SelfServiceIcon        MobileDeviceApplicationSubsetIcon   `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage      bool                                `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories  []SharedResourceSelfServiceCategory `xml:"self_service_categories>category,omitempty"`
	Notification           bool                                `xml:"notification,omitempty"`
	NotificationSubject    string                              `xml:"notification_subject,omitempty"`
	NotificationMessage    string                              `xml:"notification_message,omitempty"`
}

type MobileDeviceApplicationSubsetGeneralVPP struct {
	AssignVPPDeviceBasedLicenses bool `xml:"assign_vpp_device_based_licenses,omitempty"`
	VPPAdminAccountID            int  `xml:"vpp_admin_account_id,omitempty"`
}

type MobileDeviceApplicationSubsetGeneralAppConfiguration struct {
	Preferences string `xml:"preferences,omitempty"`
}

// Shared Structs

type MobileDeviceApplicationSubsetScope struct {
	AllMobileDevices   bool                                             `xml:"all_mobile_devices,omitempty"`
	AllJSSUsers        bool                                             `xml:"all_jss_users,omitempty"`
	MobileDevices      []MobileDeviceApplicationSubsetMobileDevice      `xml:"mobile_devices>mobile_device,omitempty"`
	Buildings          []MobileDeviceApplicationSubsetBuilding          `xml:"buildings>building,omitempty"`
	Departments        []MobileDeviceApplicationSubsetDepartment        `xml:"departments>department,omitempty"`
	MobileDeviceGroups []MobileDeviceApplicationSubsetMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	JSSUsers           []MobileDeviceApplicationSubsetJSSUser           `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []MobileDeviceApplicationSubsetJSSUserGroup      `xml:"jss_user_groups>user_group,omitempty"`
	Limitations        MobileDeviceApplicationSubsetLimitation          `xml:"limitations,omitempty"`
	Exclusions         MobileDeviceApplicationSubsetExclusion           `xml:"exclusions,omitempty"`
}

type MobileDeviceApplicationSubsetLimitation struct {
	Users           []MobileDeviceApplicationSubsetUser           `xml:"users>user,omitempty"`
	UserGroups      []MobileDeviceApplicationSubsetUserGroup      `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []MobileDeviceApplicationSubsetNetworkSegment `xml:"network_segments>network_segment,omitempty"`
}

type MobileDeviceApplicationSubsetExclusion struct {
	MobileDevices      []MobileDeviceApplicationSubsetMobileDevice      `xml:"mobile_devices>mobile_device,omitempty"`
	Buildings          []MobileDeviceApplicationSubsetBuilding          `xml:"buildings>building,omitempty"`
	Departments        []MobileDeviceApplicationSubsetDepartment        `xml:"departments>department,omitempty"`
	MobileDeviceGroups []MobileDeviceApplicationSubsetMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	JSSUsers           []MobileDeviceApplicationSubsetJSSUser           `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []MobileDeviceApplicationSubsetJSSUserGroup      `xml:"jss_user_groups>user_group,omitempty"`
}

type MobileDeviceApplicationSubsetIcon struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name"`
	URI  string `xml:"uri,omitempty"`
	Data string `xml:"data,omitempty"`
}

type MobileDeviceApplicationSubsetMobileDevice struct {
	ID             int    `xml:"id,omitempty"`
	Name           string `xml:"name,omitempty"`
	UDID           string `xml:"udid,omitempty"`
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"`
}

type MobileDeviceApplicationSubsetBuilding struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MobileDeviceApplicationSubsetDepartment struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MobileDeviceApplicationSubsetMobileDeviceGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MobileDeviceApplicationSubsetJSSUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MobileDeviceApplicationSubsetJSSUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MobileDeviceApplicationSubsetUser struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MobileDeviceApplicationSubsetUserGroup struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type MobileDeviceApplicationSubsetNetworkSegment struct {
	ID   int    `xml:"id,omitempty"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name,omitempty"`
}

// GetMobileDeviceApplications retrieves a serialized list of mobile device applications.
func (c *Client) GetMobileDeviceApplications() (*ResponseMobileDeviceApplicationsList, error) {
	endpoint := uriMobileDeviceApplications

	var mobileDeviceApps ResponseMobileDeviceApplicationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &mobileDeviceApps)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "mobile device applications", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &mobileDeviceApps, nil
}

// CRUD

// GetMobileDeviceApplicationByID fetches a specific mobile device application by its ID from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByID(id int) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceApplications, id)

	var app ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device application", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// GetMobileDeviceApplicationByName fetches a specific mobile device application by its name from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByName(name string) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceApplications, name)

	var app ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device application", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// GetMobileDeviceApplicationByAppBundleID fetches a specific mobile device application by its bundle ID from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByAppBundleID(id string) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/bundleid/%s", uriMobileDeviceApplications, id)

	var app ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device application (app bundle id)", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// GetMobileDeviceApplicationByAppBundleIDAndVersion fetches a specific mobile device application by its bundle ID and version from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByAppBundleIDAndVersion(id string, version string) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/bundleid/%s/version/%s", uriMobileDeviceApplications, id, version)

	var app ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device application (by bundle id and version)", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// GetMobileDeviceApplicationByIDAndDataSubset fetches a specific mobile device application by its ID and a specified data subset from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByIDAndDataSubset(id int, subset string) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriMobileDeviceApplications, id, subset)

	var app ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "mobile device application with data subset", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// GetMobileDeviceApplicationByNameAndDataSubset fetches a specific mobile device application by its name and a specified data subset from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByNameAndDataSubset(name string, subset string) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriMobileDeviceApplications, name, subset)

	var app ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "mobile device application and data subset", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// CreateMobileDeviceApplication creates a new mobile device application on the Jamf Pro server.
func (c *Client) CreateMobileDeviceApplication(app *ResourceMobileDeviceApplication) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceApplications)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_application"`
		*ResourceMobileDeviceApplication
	}{
		ResourceMobileDeviceApplication: app,
	}

	var responseApp ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseApp)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "mobile device application", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseApp, nil
}

// UpdateMobileDeviceApplicationByID updates a mobile device application by its ID on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceApplicationByID(id int, app *ResourceMobileDeviceApplication) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceApplications, id)

	// Wrap the application with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_application"`
		*ResourceMobileDeviceApplication
	}{
		ResourceMobileDeviceApplication: app,
	}

	var responseApp ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseApp)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "mobile device application", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseApp, nil
}

// UpdateMobileDeviceApplicationByName updates a mobile device application by its name on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceApplicationByName(name string, app *ResourceMobileDeviceApplication) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceApplications, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_application"`
		*ResourceMobileDeviceApplication
	}{
		ResourceMobileDeviceApplication: app,
	}

	var responseApp ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseApp)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "mobile device application", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseApp, nil
}

// UpdateMobileDeviceApplicationByApplicationBundleID updates a mobile device application by its bundle ID on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceApplicationByApplicationBundleID(id string, app *ResourceMobileDeviceApplication) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/bundleid/%s", uriMobileDeviceApplications, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_application"`
		*ResourceMobileDeviceApplication
	}{
		ResourceMobileDeviceApplication: app,
	}

	var responseApp ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseApp)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "mobile device application (app bundle id)", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseApp, nil
}

// UpdateMobileDeviceApplicationByIDAndAppVersion updates a mobile device application by its ID and application version on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceApplicationByIDAndAppVersion(id int, version string, app *ResourceMobileDeviceApplication) (*ResourceMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/id/%d/version/%s", uriMobileDeviceApplications, id, version)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_application"`
		*ResourceMobileDeviceApplication
	}{
		ResourceMobileDeviceApplication: app,
	}

	var responseApp ResourceMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseApp)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "mobile device application and app version", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseApp, nil
}

// DeleteMobileDeviceApplicationpByID deletes a mobile device application by its ID from the Jamf Pro server.
func (c *Client) DeleteMobileDeviceApplicationpByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceApplications, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device application", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceApplicationByName deletes a mobile device application by its name from the Jamf Pro server.
func (c *Client) DeleteMobileDeviceApplicationByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceApplications, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "mobile device application", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceApplicationByBundleID deletes a mobile device application by its bundle ID from the Jamf Pro server.
func (c *Client) DeleteMobileDeviceApplicationByBundleID(id string) error {
	endpoint := fmt.Sprintf("%s/bundleid/%s", uriMobileDeviceApplications, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device application (bundle id)", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceApplicationByBundleIDAndVersion deletes a mobile device application by its bundle ID and version from the Jamf Pro server.
func (c *Client) DeleteMobileDeviceApplicationByBundleIDAndVersion(id string, version string) error {
	endpoint := fmt.Sprintf("%s/bundleid/%s/version/%s", uriMobileDeviceApplications, id, version)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "mobile device application (bundle id and version)", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
