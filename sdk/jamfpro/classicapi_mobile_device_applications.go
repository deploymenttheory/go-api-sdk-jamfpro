// classicapi_mobile_device_applications.go
// Jamf Pro Classic Api - Mobile Device Applications
// api reference: https://developer.jamf.com/jamf-pro/reference/mobiledeviceapplications
// Jamf Pro Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMobileDeviceApplications = "/JSSResource/mobiledeviceapplications"

// ResponseMobileDeviceApplicationsList represents the response for a list of mobile device applications.
type ResponseMobileDeviceApplicationsList struct {
	MobileDeviceApplications []MobileDeviceApplicationItem `xml:"mobile_device_application"`
}

// MobileDeviceApplicationItem represents a single mobile device application item.
type MobileDeviceApplicationItem struct {
	ID          int     `xml:"id"`
	Name        string  `xml:"name"`
	DisplayName string  `xml:"display_name"`
	BundleID    string  `xml:"bundle_id"`
	Version     float64 `xml:"version"`
	InternalApp bool    `xml:"internal_app"`
}

// ResponseMobileDeviceApplication represents the detailed structure of a single mobile device application.
type ResponseMobileDeviceApplication struct {
	General          MobileDeviceApplicationGeneral       `xml:"general"`
	Scope            MobileDeviceApplicationScope         `xml:"scope"`
	SelfService      MobileDeviceApplicationSelfService   `xml:"self_service"`
	VPP              MobileDeviceApplicationVPP           `xml:"vpp,omitempty"`
	AppConfiguration MobileDeviceApplicationConfiguration `xml:"app_configuration,omitempty"`
}

type MobileDeviceApplicationGeneral struct {
	ID                               int                             `xml:"id,omitempty"`
	Name                             string                          `xml:"name"`
	DisplayName                      string                          `xml:"display_name"`
	Description                      string                          `xml:"description,omitempty"`
	BundleID                         string                          `xml:"bundle_id"`
	Version                          string                          `xml:"version"`
	InternalApp                      bool                            `xml:"internal_app,omitempty"`
	OsType                           string                          `xml:"os_type,omitempty"`
	Category                         MobileDeviceApplicationCategory `xml:"category"`
	IPA                              MobileDeviceApplicationIPA      `xml:"ipa,omitempty"`
	Icon                             MobileDeviceApplicationIcon     `xml:"icon"`
	ProvisioningProfile              int                             `xml:"mobile_device_provisioning_profile,omitempty"`
	ITunesStoreURL                   string                          `xml:"itunes_store_url,omitempty"`
	MakeAvailableAfterInstall        bool                            `xml:"make_available_after_install,omitempty"`
	ITunesCountryRegion              string                          `xml:"itunes_country_region,omitempty"`
	ITunesSyncTime                   int                             `xml:"itunes_sync_time,omitempty"`
	DeploymentType                   string                          `xml:"deployment_type,omitempty"`
	DeployAutomatically              bool                            `xml:"deploy_automatically,omitempty"`
	DeployAsManagedApp               bool                            `xml:"deploy_as_managed_app,omitempty"`
	RemoveAppWhenMDMProfileIsRemoved bool                            `xml:"remove_app_when_mdm_profile_is_removed,omitempty"`
	PreventBackupOfAppData           bool                            `xml:"prevent_backup_of_app_data,omitempty"`
	KeepDescriptionAndIconUpToDate   bool                            `xml:"keep_description_and_icon_up_to_date,omitempty"`
	Free                             bool                            `xml:"free,omitempty"`
	TakeOverManagement               bool                            `xml:"take_over_management,omitempty"`
	HostExternally                   bool                            `xml:"host_externally,omitempty"`
	ExternalURL                      string                          `xml:"external_url,omitempty"`
	Site                             MobileDeviceApplicationSite     `xml:"site"`
}

type MobileDeviceApplicationCategory struct {
	ID   int    `xml:"id,omitempty"` // ID is optional
	Name string `xml:"name"`         // Name is required
}

type MobileDeviceApplicationIPA struct {
	Name string `xml:"name,omitempty"` // Optional fields
	URI  string `xml:"uri,omitempty"`
	Data string `xml:"data,omitempty"`
}

type MobileDeviceApplicationIcon struct {
	ID   int    `xml:"id,omitempty"` // ID is optional
	Name string `xml:"name"`         // Name is required
	URI  string `xml:"uri,omitempty"`
	Data string `xml:"data,omitempty"`
}

type MobileDeviceApplicationSite struct {
	ID   int    `xml:"id,omitempty"` // ID is optional
	Name string `xml:"name"`         // Name is required
}

type MobileDeviceApplicationScope struct {
	AllMobileDevices   bool                                       `xml:"all_mobile_devices,omitempty"`
	AllJSSUsers        bool                                       `xml:"all_jss_users,omitempty"`
	MobileDevices      []MobileDeviceApplicationMobileDevice      `xml:"mobile_devices>mobile_device,omitempty"`
	Buildings          []MobileDeviceApplicationBuilding          `xml:"buildings>building,omitempty"`
	Departments        []MobileDeviceApplicationDepartment        `xml:"departments>department,omitempty"`
	MobileDeviceGroups []MobileDeviceApplicationMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group,omitempty"`
	JSSUsers           []MobileDeviceApplicationJSSUser           `xml:"jss_users>user,omitempty"`
	JSSUserGroups      []MobileDeviceApplicationJSSUserGroup      `xml:"jss_user_groups>user_group,omitempty"`
	Limitations        MobileDeviceApplicationLimitation          `xml:"limitations,omitempty"`
	Exclusions         MobileDeviceApplicationExclusion           `xml:"exclusions,omitempty"`
}

type MobileDeviceApplicationMobileDevice struct {
	ID             int    `xml:"id,omitempty"`               // ID is optional
	Name           string `xml:"name,omitempty"`             // Name is optional
	UDID           string `xml:"udid,omitempty"`             // UDID is optional
	WifiMacAddress string `xml:"wifi_mac_address,omitempty"` // WifiMacAddress is optional
}

type MobileDeviceApplicationBuilding struct {
	ID   int    `xml:"id,omitempty"`   // ID is optional
	Name string `xml:"name,omitempty"` // Name is optional
}

type MobileDeviceApplicationDepartment struct {
	ID   int    `xml:"id,omitempty"`   // ID is optional
	Name string `xml:"name,omitempty"` // Name is optional
}

type MobileDeviceApplicationMobileDeviceGroup struct {
	ID   int    `xml:"id,omitempty"`   // ID is optional
	Name string `xml:"name,omitempty"` // Name is optional
}

type MobileDeviceApplicationJSSUser struct {
	ID   int    `xml:"id,omitempty"`   // ID is optional
	Name string `xml:"name,omitempty"` // Name is optional
}

type MobileDeviceApplicationJSSUserGroup struct {
	ID   int    `xml:"id,omitempty"`   // ID is optional
	Name string `xml:"name,omitempty"` // Name is optional
}

type MobileDeviceApplicationLimitation struct {
	Users           []MobileDeviceApplicationUser           `xml:"users>user,omitempty"`                       // Optional
	UserGroups      []MobileDeviceApplicationUserGroup      `xml:"user_groups>user_group,omitempty"`           // Optional
	NetworkSegments []MobileDeviceApplicationNetworkSegment `xml:"network_segments>network_segment,omitempty"` // Optional
}

type MobileDeviceApplicationUser struct {
	ID   int    `xml:"id,omitempty"`   // ID is optional
	Name string `xml:"name,omitempty"` // Name is optional
}

type MobileDeviceApplicationUserGroup struct {
	ID   int    `xml:"id,omitempty"`   // ID is optional
	Name string `xml:"name,omitempty"` // Name is optional
}

type MobileDeviceApplicationNetworkSegment struct {
	ID   int    `xml:"id,omitempty"`   // ID is optional
	UID  string `xml:"uid,omitempty"`  // UID is optional
	Name string `xml:"name,omitempty"` // Name is optional
}

type MobileDeviceApplicationExclusion struct {
	MobileDevices      []MobileDeviceApplicationMobileDevice      `xml:"mobile_devices>mobile_device,omitempty"`             // Optional
	Buildings          []MobileDeviceApplicationBuilding          `xml:"buildings>building,omitempty"`                       // Optional
	Departments        []MobileDeviceApplicationDepartment        `xml:"departments>department,omitempty"`                   // Optional
	MobileDeviceGroups []MobileDeviceApplicationMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group,omitempty"` // Optional
	JSSUsers           []MobileDeviceApplicationJSSUser           `xml:"jss_users>user,omitempty"`                           // Optional
	JSSUserGroups      []MobileDeviceApplicationJSSUserGroup      `xml:"jss_user_groups>user_group,omitempty"`               // Optional
}

type MobileDeviceApplicationSelfService struct {
	SelfServiceDescription string                      `xml:"self_service_description,omitempty"`         // Optional
	SelfServiceIcon        MobileDeviceApplicationIcon `xml:"self_service_icon,omitempty"`                // Optional
	FeatureOnMainPage      bool                        `xml:"feature_on_main_page,omitempty"`             // Optional
	SelfServiceCategories  []SelfServiceCategoryItem   `xml:"self_service_categories>category,omitempty"` // Optional
	Notification           bool                        `xml:"notification,omitempty"`                     // Optional
	NotificationSubject    string                      `xml:"notification_subject,omitempty"`             // Optional
	NotificationMessage    string                      `xml:"notification_message,omitempty"`             // Optional
}

type SelfServiceCategoryItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MobileDeviceApplicationVPP struct {
	AssignVPPDeviceBasedLicenses bool `xml:"assign_vpp_device_based_licenses,omitempty"` // Optional as per the documentation.
	VPPAdminAccountID            int  `xml:"vpp_admin_account_id,omitempty"`             // Optional as per the documentation.
}

type MobileDeviceApplicationConfiguration struct {
	Preferences string `xml:"preferences,omitempty"` // Optional as per the documentation.
}

// GetMobileDeviceApplications retrieves a serialized list of mobile device applications.
func (c *Client) GetMobileDeviceApplications() (*ResponseMobileDeviceApplicationsList, error) {
	endpoint := uriMobileDeviceApplications

	var mobileDeviceApps ResponseMobileDeviceApplicationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &mobileDeviceApps)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device applications: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &mobileDeviceApps, nil
}

// GetMobileDeviceApplicationByID fetches a specific mobile device application by its ID from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByID(id int) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceApplications, id)

	var app ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device application by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// GetMobileDeviceApplicationByName fetches a specific mobile device application by its name from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByName(name string) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceApplications, name)

	var app ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device application by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// GetMobileDeviceApplicationByAppBundleID fetches a specific mobile device application by its bundle ID from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByAppBundleID(bundleID string) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/bundleid/%s", uriMobileDeviceApplications, bundleID)

	var app ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device application by bundle ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// GetMobileDeviceApplicationByAppBundleIDAndVersion fetches a specific mobile device application by its bundle ID and version from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByAppBundleIDAndVersion(bundleID string, version string) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/bundleid/%s/version/%s", uriMobileDeviceApplications, bundleID, version)

	var app ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device application by bundle ID and version: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// GetMobileDeviceApplicationByIDAndDataSubset fetches a specific mobile device application by its ID and a specified data subset from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByIDAndDataSubset(id int, subset string) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriMobileDeviceApplications, id, subset)

	var app ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device application by ID and data subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// GetMobileDeviceApplicationByNameAndDataSubset fetches a specific mobile device application by its name and a specified data subset from the Jamf Pro server.
func (c *Client) GetMobileDeviceApplicationByNameAndDataSubset(name string, subset string) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriMobileDeviceApplications, name, subset)

	var app ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &app)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device application by name and data subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &app, nil
}

// CreateMobileDeviceApplication creates a new mobile device application on the Jamf Pro server.
func (c *Client) CreateMobileDeviceApplication(app *ResponseMobileDeviceApplication) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceApplications)

	// Set default values for site and category if not included within request
	if app.General.Site.ID == 0 && app.General.Site.Name == "" {
		app.General.Site = MobileDeviceApplicationSite{
			ID:   -1,
			Name: "None",
		}
	}
	if app.General.Category.ID == 0 && app.General.Category.Name == "" {
		app.General.Category = MobileDeviceApplicationCategory{
			ID:   -1,
			Name: "No category assigned",
		}
	}

	// Wrap the application with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_application"`
		*ResponseMobileDeviceApplication
	}{
		ResponseMobileDeviceApplication: app,
	}

	var responseApp ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseApp)
	if err != nil {
		return nil, fmt.Errorf("failed to create mobile device application: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseApp, nil
}

// UpdateMobileDeviceApplicationByID updates a mobile device application by its ID on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceApplicationByID(id int, app *ResponseMobileDeviceApplication) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceApplications, id)

	// Wrap the application with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_application"`
		*ResponseMobileDeviceApplication
	}{
		ResponseMobileDeviceApplication: app,
	}

	var responseApp ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseApp)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device application by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseApp, nil
}

// UpdateMobileDeviceApplicationByName updates a mobile device application by its name on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceApplicationByName(name string, app *ResponseMobileDeviceApplication) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceApplications, name)

	// Wrap the application with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_application"`
		*ResponseMobileDeviceApplication
	}{
		ResponseMobileDeviceApplication: app,
	}

	var responseApp ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseApp)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device application by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseApp, nil
}

// UpdateMobileDeviceApplicationByApplicationBundleID updates a mobile device application by its bundle ID on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceApplicationByApplicationBundleID(bundleID string, app *ResponseMobileDeviceApplication) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/bundleid/%s", uriMobileDeviceApplications, bundleID)

	// Wrap the application with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_application"`
		*ResponseMobileDeviceApplication
	}{
		ResponseMobileDeviceApplication: app,
	}

	var responseApp ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseApp)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device application by bundle ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseApp, nil
}

// UpdateMobileDeviceApplicationByIDAndAppVersion updates a mobile device application by its ID and application version on the Jamf Pro server.
func (c *Client) UpdateMobileDeviceApplicationByIDAndAppVersion(id int, version string, app *ResponseMobileDeviceApplication) (*ResponseMobileDeviceApplication, error) {
	endpoint := fmt.Sprintf("%s/id/%d/version/%s", uriMobileDeviceApplications, id, version)

	// Wrap the application with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_application"`
		*ResponseMobileDeviceApplication
	}{
		ResponseMobileDeviceApplication: app,
	}

	var responseApp ResponseMobileDeviceApplication
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseApp)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device application by ID and version: %v", err)
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
		return fmt.Errorf("failed to delete mobile device application by ID: %v", err)
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
		return fmt.Errorf("failed to delete mobile device application by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceApplicationByBundleID deletes a mobile device application by its bundle ID from the Jamf Pro server.
func (c *Client) DeleteMobileDeviceApplicationByBundleID(bundleID string) error {
	endpoint := fmt.Sprintf("%s/bundleid/%s", uriMobileDeviceApplications, bundleID)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device application by bundle ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceApplicationByBundleIDAndVersion deletes a mobile device application by its bundle ID and version from the Jamf Pro server.
func (c *Client) DeleteMobileDeviceApplicationByBundleIDAndVersion(bundleID string, version string) error {
	endpoint := fmt.Sprintf("%s/bundleid/%s/version/%s", uriMobileDeviceApplications, bundleID, version)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device application by bundle ID and version: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
