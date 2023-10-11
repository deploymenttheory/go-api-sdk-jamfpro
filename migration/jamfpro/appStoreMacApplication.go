// appStoreMacApplication.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAppStoreMacApplication = "/JSSResource/macapplications"

// Top-level struct
type ResponseAppStoreMacApplication struct {
	General     General     `xml:"general"`
	Scope       Scope       `xml:"scope"`
	SelfService SelfService `xml:"self_service"`
}

// Tier 2 - General Section
type General struct {
	ID       int                                    `xml:"id,omitempty"`
	Name     string                                 `xml:"name"`
	Version  string                                 `xml:"version"`
	IsFree   bool                                   `xml:"is_free,omitempty"`
	BundleID string                                 `xml:"bundle_id"`
	URL      string                                 `xml:"url"`
	Category AppStoreMacApplicationDataSubsetIDName `xml:"category,omitempty"`
	Site     AppStoreMacApplicationDataSubsetIDName `xml:"site,omitempty"`
}

// Tier 2 - Scope section
type Scope struct {
	AllComputers   bool                                            `xml:"all_computers,omitempty"`
	AllJSSUsers    bool                                            `xml:"all_jss_users,omitempty"`
	Buildings      []AppStoreMacApplicationDataSubsetBuilding      `xml:"buildings,omitempty"`
	Departments    []AppStoreMacApplicationDataSubsetDepartment    `xml:"departments,omitempty"`
	Computers      []AppStoreMacApplicationDataSubsetComputer      `xml:"computers,omitempty"`
	ComputerGroups []AppStoreMacApplicationDataSubsetComputerGroup `xml:"computer_groups,omitempty"`
	JSSUsers       []AppStoreMacApplicationDataSubsetJSSUser       `xml:"jss_users,omitempty"`
	JSSUserGroups  []AppStoreMacApplicationDataSubsetJSSUserGroup  `xml:"jss_user_groups,omitempty"`
	Limitations    struct {
		Users           []AppStoreMacApplicationDataSubsetJSSUser      `xml:"users,omitempty"`
		UserGroups      []AppStoreMacApplicationDataSubsetJSSUserGroup `xml:"user_groups,omitempty"`
		NetworkSegments []struct {
			NetworkSegment AppStoreMacApplicationDataSubsetIDName `xml:"network_segment,omitempty"`
		} `xml:"network_segments"`
	} `xml:"limitations,omitempty"`
	Exclusions struct {
		Buildings       []AppStoreMacApplicationDataSubsetBuilding     `xml:"buildings,omitempty"`
		Departments     []AppStoreMacApplicationDataSubsetDepartment   `xml:"departments,omitempty"`
		Users           []AppStoreMacApplicationDataSubsetJSSUser      `xml:"users,omitempty"`
		UserGroups      []AppStoreMacApplicationDataSubsetJSSUserGroup `xml:"user_groups,omitempty"`
		NetworkSegments []struct {
			NetworkSegment struct {
				ID   int    `xml:"id,omitempty"`
				UID  string `xml:"uid,omitempty"`
				Name string `xml:"name"`
			} `xml:"network_segment,omitempty"`
		} `xml:"network_segments"`
		Computers      []AppStoreMacApplicationDataSubsetComputer      `xml:"computers,omitempty"`
		ComputerGroups []AppStoreMacApplicationDataSubsetComputerGroup `xml:"computer_groups,omitempty"`
		JSSUsers       []AppStoreMacApplicationDataSubsetJSSUser       `xml:"jss_users,omitempty"`
		JSSUserGroups  []AppStoreMacApplicationDataSubsetJSSUserGroup  `xml:"jss_user_groups,omitempty"`
	} `xml:"exclusions,omitempty"`
}

// Tier 2 - SelfService section
type SelfService struct {
	InstallButtonText           string                                                `xml:"install_button_text,omitempty"`
	SelfServiceDescription      string                                                `xml:"self_service_description,omitempty"`
	ForceUsersToViewDescription bool                                                  `xml:"force_users_to_view_description,omitempty"`
	SelfServiceIcon             AppStoreMacApplicationDataSubsetSelfServiceIcon       `xml:"self_service_icon,omitempty"`
	FeatureOnMainPage           bool                                                  `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories       []AppStoreMacApplicationDataSubsetSelfServiceCategory `xml:"self_service_categories,omitempty"`
	Notification                string                                                `xml:"notification,omitempty"`
	NotificationSubject         string                                                `xml:"notification_subject,omitempty"`
	NotificationMessage         string                                                `xml:"notification_message,omitempty"`
	VPP                         AppStoreMacApplicationDataSubsetVPP                   `xml:"vpp,omitempty"`
}

// Tier 3 - Scope Section

// Shared inner structs for reusability
type AppStoreMacApplicationDataSubsetIDName struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type AppStoreMacApplicationDataSubsetBuilding struct {
	Building AppStoreMacApplicationDataSubsetIDName `xml:"building,omitempty"`
}

type AppStoreMacApplicationDataSubsetDepartment struct {
	Department AppStoreMacApplicationDataSubsetIDName `xml:"department,omitempty"`
}

type AppStoreMacApplicationDataSubsetComputer struct {
	Computer struct {
		ID   int    `xml:"id,omitempty"`
		Name string `xml:"name"`
		UDID string `xml:"udid,omitempty"`
	} `xml:"computer"`
}

type AppStoreMacApplicationDataSubsetComputerGroup struct {
	ComputerGroup AppStoreMacApplicationDataSubsetIDName `xml:"computer_group,omitempty"`
}

type AppStoreMacApplicationDataSubsetJSSUser struct {
	User AppStoreMacApplicationDataSubsetIDName `xml:"user,omitempty"`
}

type AppStoreMacApplicationDataSubsetJSSUserGroup struct {
	UserGroup AppStoreMacApplicationDataSubsetIDName `xml:"user_group,omitempty"`
}

// Tier 3 - Self Service section
type AppStoreMacApplicationDataSubsetSelfServiceIcon struct {
	ID   int    `xml:"id,omitempty"`
	URI  string `xml:"uri,omitempty"`
	Data string `xml:"data,omitempty"`
}

type AppStoreMacApplicationDataSubsetSelfServiceCategory struct {
	Category struct {
		ID        int    `xml:"id,omitempty"`
		Name      string `xml:"name"`
		DisplayIn bool   `xml:"display_in,omitempty"`
		FeatureIn bool   `xml:"feature_in,omitempty"`
	} `xml:"category"`
}

// Tier 3 - VPP section
type AppStoreMacApplicationDataSubsetVPP struct {
	AssignVPPDeviceBasedLicenses bool `xml:"assign_vpp_device_based_licenses,omitempty"`
	VPPAdminAccountID            int  `xml:"vpp_admin_account_id,omitempty"`
}

// List all app store mac apps

type ResponseAppStoreMacApplicationList struct {
	MacApplications []MacApplication `xml:"mac_application"`
}

type MacApplication struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Create / Update - App store mac application

type AppStoreMacApplication struct {
	XMLName     xml.Name    `xml:"mac_application"`
	General     General     `xml:"general"`
	Scope       Scope       `xml:"scope"`
	SelfService SelfService `xml:"self_service"`
}

//--- appStoreMacApplication CRUD Functions ---//

// GetAppStoreMacApplicationByID retrieves the App Store Mac Application by its ID
func (c *Client) GetAppStoreMacApplicationByID(id int) (*ResponseAppStoreMacApplication, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIAppStoreMacApplication, id)

	var app ResponseAppStoreMacApplication
	if err := c.DoRequest("GET", url, nil, nil, &app); err != nil {
		// The raw response logging is already handled within DoRequestDebug, so we don't need to explicitly log it here.
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &app, nil
}

// GetAppStoreMacApplications retrieves a list of all App Store Mac Applications
func (c *Client) GetAppStoreMacApplications() ([]ResponseAppStoreMacApplicationList, error) {
	url := uriAPIAppStoreMacApplication

	// Define a slice of the adjusted struct to hold the response
	var appList []ResponseAppStoreMacApplicationList

	if err := c.DoRequest("GET", url, nil, nil, &appList); err != nil {
		// Handle error
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return appList, nil
}

// GetAppStoreMacApplicationByName retrieves the App Store Mac Application by its name
func (c *Client) GetAppStoreMacApplicationByName(appName string) (*ResponseAppStoreMacApplication, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIAppStoreMacApplication, appName)

	var app ResponseAppStoreMacApplication
	if err := c.DoRequest("GET", url, nil, nil, &app); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &app, nil
}

// CreateAppStoreMacApplication creates a new App Store Mac Application
func (c *Client) CreateAppStoreMacApplication(app *AppStoreMacApplication) (*ResponseAppStoreMacApplication, error) {
	// URL for creating a new App Store Mac Application
	url := uriAPIAppStoreMacApplication

	// Execute the request
	var responseApp ResponseAppStoreMacApplication
	if err := c.DoRequest("POST", url, app, nil, &responseApp); err != nil {
		return nil, fmt.Errorf("failed to create App Store Mac Application: %v", err)
	}

	return &responseApp, nil
}

// UpdateAppStoreMacApplication updates an existing App Store Mac Application
func (c *Client) UpdateAppStoreMacApplication(id int, app *AppStoreMacApplication) (*ResponseAppStoreMacApplication, error) {
	// URL for updating the App Store Mac Application with the specified ID
	url := fmt.Sprintf("%s/id/%d", uriAPIAppStoreMacApplication, id)

	// Execute the request
	var responseApp ResponseAppStoreMacApplication
	if err := c.DoRequest("PUT", url, app, nil, &responseApp); err != nil {
		return nil, fmt.Errorf("failed to update App Store Mac Application: %v", err)
	}

	return &responseApp, nil
}

// DeleteAppStoreMacApplicationByID deletes an App Store Mac Application by its ID
func (c *Client) DeleteAppStoreMacApplicationByID(id int) error {
	// URL for deleting the App Store Mac Application with the specified ID
	url := fmt.Sprintf("%s/id/%d", uriAPIAppStoreMacApplication, id)

	// Execute the request
	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete App Store Mac Application: %v", err)
	}

	return nil
}

// DeleteAppStoreMacApplicationByName deletes an App Store Mac Application by its name
func (c *Client) DeleteAppStoreMacApplicationByName(appName string) error {
	// URL for deleting the App Store Mac Application with the specified name
	url := fmt.Sprintf("%s/name/%s", uriAPIAppStoreMacApplication, appName)

	// Execute the request
	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete App Store Mac Application named '%s': %v", appName, err)
	}

	return nil
}
