// classicapi_mac_applications.go
// Jamf Pro Classic Api - VPP Mac Applications
// api reference: https://developer.jamf.com/jamf-pro/reference/macapplications
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriVPPMacApplications = "/JSSResource/macapplications"

type ResponseMacApplicationsList struct {
	MacApplications []struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	} `xml:"mac_application"`
}

// ResourceMacApplications represents the detailed structure of a Mac Application response.
type ResourceMacApplications struct {
	General struct {
		ID       int    `xml:"id"`
		Name     string `xml:"name"`
		Version  string `xml:"version"`
		IsFree   bool   `xml:"is_free"`
		BundleID string `xml:"bundle_id"`
		URL      string `xml:"url"`
		Category struct {
			ID   int    `xml:"id"`
			Name string `xml:"name"`
		} `xml:"category"`
		Site struct {
			ID   int    `xml:"id"`
			Name string `xml:"name"`
		} `xml:"site"`
	} `xml:"general"`
	Scope struct {
		AllComputers   bool                             `xml:"all_computers"`
		AllJSSUsers    bool                             `xml:"all_jss_users"`
		Buildings      []MacAppSubsetScopeBuilding      `xml:"buildings>building"`
		Departments    []MacAppSubsetScopeDepartment    `xml:"departments>department"`
		Computers      []MacAppSubsetScopeComputer      `xml:"computers>computer"`
		ComputerGroups []MacAppSubsetScopeComputerGroup `xml:"computer_groups>computer_group"`
		JSSUsers       []MacAppSubsetScopeUser          `xml:"jss_users>user"`
		JSSUserGroups  []MacAppSubsetScopeUserGroup     `xml:"jss_user_groups>user_group"`
		Limitations    MacAppScopeLimitations           `xml:"limitations"`
		Exclusions     MacAppScopeExclusions            `xml:"exclusions"`
	} `xml:"scope"`
	SelfService MacAppSubsetSelfService `xml:"self_service"`
}

type MacAppScopeLimitations struct {
	Users           []MacAppSubsetScopeUser           `xml:"users>user"`
	UserGroups      []MacAppSubsetScopeUserGroup      `xml:"user_groups>user_group"`
	NetworkSegments []MacAppSubsetScopeNetworkSegment `xml:"network_segments>network_segment"`
}

type MacAppScopeExclusions struct {
	Buildings       []MacAppSubsetScopeBuilding       `xml:"buildings>building"`
	Departments     []MacAppSubsetScopeDepartment     `xml:"departments>department"`
	Users           []MacAppSubsetScopeUser           `xml:"users>user"`
	UserGroups      []MacAppSubsetScopeUserGroup      `xml:"user_groups>user_group"`
	NetworkSegments []MacAppSubsetScopeNetworkSegment `xml:"network_segments>network_segment"`
	Computers       []MacAppSubsetScopeComputer       `xml:"computers>computer"`
	ComputerGroups  []MacAppSubsetScopeComputerGroup  `xml:"computer_groups>computer_group"`
	JSSUsers        []MacAppSubsetScopeUser           `xml:"jss_users>user"`
	JSSUserGroups   []MacAppSubsetScopeUserGroup      `xml:"jss_user_groups>user_group"`
}

// Define structs for each scope component (Building, Department, Computer, etc.)
type MacAppSubsetScopeBuilding struct {
	Building MacAppSubsetBuilding `xml:"building"`
}

type MacAppSubsetBuilding struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Struct definitions for Department, Computer, ComputerGroup, User, UserGroup

type MacAppSubsetScopeDepartment struct {
	Department MacAppSubsetDepartment `xml:"department"`
}

type MacAppSubsetDepartment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MacAppSubsetScopeComputer struct {
	Computer MacAppSubsetComputer `xml:"computer"`
}

type MacAppSubsetComputer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	UDID string `xml:"udid"`
}

type MacAppSubsetScopeComputerGroup struct {
	ComputerGroup MacAppSubsetComputerGroup `xml:"computer_group"`
}

type MacAppSubsetComputerGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MacAppSubsetScopeUser struct {
	User MacAppSubsetUser `xml:"user"`
}

type MacAppSubsetUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MacAppSubsetScopeUserGroup struct {
	UserGroup MacAppSubsetUserGroup `xml:"user_group"`
}

type MacAppSubsetUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type MacAppSubsetScopeNetworkSegment struct {
	ID   int    `xml:"id"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name"`
}

type MacAppSubsetSelfService struct {
	InstallButtonText           string `xml:"install_button_text"`
	SelfServiceDescription      string `xml:"self_service_description"`
	ForceUsersToViewDescription bool   `xml:"force_users_to_view_description"`
	SelfServiceIcon             struct {
		ID   int    `xml:"id"`
		URI  string `xml:"uri"`
		Data string `xml:"data"`
	} `xml:"self_service_icon"`
	FeatureOnMainPage     bool `xml:"feature_on_main_page"`
	SelfServiceCategories []struct {
		ID        int    `xml:"id"`
		Name      string `xml:"name"`
		DisplayIn bool   `xml:"display_in"`
		FeatureIn bool   `xml:"feature_in"`
	} `xml:"self_service_categories>category"`
	Notification        string `xml:"notification"`
	NotificationSubject string `xml:"notification_subject"`
	NotificationMessage string `xml:"notification_message"`
	VPP                 struct {
		AssignVPPDeviceBasedLicenses bool `xml:"assign_vpp_device_based_licenses"`
		VPPAdminAccountID            int  `xml:"vpp_admin_account_id"`
	} `xml:"vpp"`
}

// GetDockItems retrieves a serialized list of vpp mac applications.
func (c *Client) GetMacApplications() (*ResponseMacApplicationsList, error) {
	endpoint := uriVPPMacApplications

	var macApps ResponseMacApplicationsList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macApps)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mac Applications: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macApps, nil
}

// GetMacApplicationByID retrieves a single Mac application by its ID.
func (c *Client) GetMacApplicationByID(id int) (*ResourceMacApplications, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPMacApplications, id)

	var macApp ResourceMacApplications
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macApp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mac Application by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macApp, nil
}

// GetMacApplicationByName retrieves a single Mac application by its name.
func (c *Client) GetMacApplicationByName(name string) (*ResourceMacApplications, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriVPPMacApplications, name)

	var macApp ResourceMacApplications
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macApp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mac Application by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macApp, nil
}

// GetMacApplicationByNameAndDataSubset retrieves a specific Mac Application by its ID and filters by a specific data subset.
// Subset values can be General, Scope, SelfService, VPPCodes and VPP.
func (c *Client) GetMacApplicationByIDAndDataSubset(id int, subset string) (*ResourceMacApplications, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriVPPMacApplications, id, subset)

	var macApp ResourceMacApplications
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macApp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mac Application by Name and Subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macApp, nil
}

// GetMacApplicationByNameAndDataSubset retrieves a specific Mac Application by its name and filters by a specific data subset.
// Subset values can be General, Scope, SelfService, VPPCodes and VPP.
func (c *Client) GetMacApplicationByNameAndDataSubset(name, subset string) (*ResourceMacApplications, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriVPPMacApplications, name, subset)

	var macApp ResourceMacApplications
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &macApp)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Mac Application by Name and Subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &macApp, nil
}

// CreateMacApplication creates a new Mac Application.
func (c *Client) CreateMacApplication(macApp ResourceMacApplications) (*ResourceMacApplications, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriVPPMacApplications) // '0' typically used for creation in APIs

	// Set default values for site if not included within request
	if macApp.General.Site.ID == 0 && macApp.General.Site.Name == "" {
		macApp.General.Site.ID = -1
		macApp.General.Site.Name = "none"
	}

	// The requestBody struct should mirror the ResourceMacApplications struct, including all nested structs
	requestBody := struct {
		XMLName xml.Name `xml:"mac_application"`
		ResourceMacApplications
	}{
		ResourceMacApplications: macApp,
	}

	var response ResourceMacApplications
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create Mac Application: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateMacApplicationByID updates an existing Mac Application by its ID.
func (c *Client) UpdateMacApplicationByID(id int, macApp ResourceMacApplications) (*ResourceMacApplications, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPMacApplications, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mac_application"`
		ResourceMacApplications
	}{
		ResourceMacApplications: macApp,
	}

	var response ResourceMacApplications
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update Mac Application by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateMacApplicationByName updates an existing Mac Application by its name.
func (c *Client) UpdateMacApplicationByName(name string, macApp ResourceMacApplications) (*ResourceMacApplications, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriVPPMacApplications, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mac_application"`
		ResourceMacApplications
	}{
		ResourceMacApplications: macApp,
	}

	var response ResourceMacApplications
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update Mac Application by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteMacApplicationByID deletes a MacApplication by its ID.
func (c *Client) DeleteMacApplicationByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriVPPMacApplications, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete VPP Mac Application Item by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMacApplicationByName deletes a MacApplication by its name.
func (c *Client) DeleteMacApplicationByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriVPPMacApplications, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete VPP Mac Application Item by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
