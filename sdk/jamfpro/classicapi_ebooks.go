// classicapi_ebooks.go
// Jamf Pro Classic Api - Ebooks
// api reference: https://developer.jamf.com/jamf-pro/reference/ebooks
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// URI for Ebooks in Jamf Pro API
const uriEbooks = "/JSSResource/ebooks"

// Struct to capture the XML response for ebooks list
type ResponseEbooksList struct {
	Size  int `xml:"size"`
	Ebook struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	} `xml:"ebook"`
}

// ResourceEbooks represents the detailed structure of an Ebook response.
type ResourceEbooks struct {
	General struct {
		ID              int    `xml:"id"`
		Name            string `xml:"name"`
		Author          string `xml:"author"`
		Version         string `xml:"version"`
		Free            bool   `xml:"free"`
		URL             string `xml:"url"`
		DeploymentType  string `xml:"deployment_type"`
		FileType        string `xml:"file_type"`
		DeployAsManaged bool   `xml:"deploy_as_managed"`
		Category        struct {
			ID   int    `xml:"id"`
			Name string `xml:"name"`
		} `xml:"category"`
		SelfServiceIcon struct {
			ID   int    `xml:"id"`
			URI  string `xml:"uri"`
			Data string `xml:"data"`
		} `xml:"self_service_icon"`
		Site struct {
			ID   int    `xml:"id"`
			Name string `xml:"name"`
		} `xml:"site"`
	} `xml:"general"`
	Scope       EbooksSubsetScope       `xml:"scope"`
	SelfService EbooksSubsetSelfService `xml:"self_service"`
}

type EbooksSubsetScope struct {
	AllComputers       bool                        `xml:"all_computers"`
	AllMobileDevices   bool                        `xml:"all_mobile_devices"`
	AllJSSUsers        bool                        `xml:"all_jss_users"`
	Computers          []EbooksSubsetComputer      `xml:"computers>computer"`
	ComputerGroups     []EbooksSubsetComputerGroup `xml:"computer_groups>computer_group"`
	MobileDevices      []EbooksSubsetMobileDevice  `xml:"mobile_devices>mobile_device"`
	MobileDeviceGroups []MobileDeviceGroup         `xml:"mobile_device_groups>mobile_device_group"`
	Buildings          []EbooksSubsetBuilding      `xml:"buildings>building"`
	Departments        []EbooksSubsetDepartment    `xml:"departments>department"`
	JSSUsers           []EbooksSubsetUser          `xml:"jss_users>user"`
	JSSUserGroups      []EbooksSubsetUserGroup     `xml:"jss_user_groups>user_group"`
	Classes            []EbooksSubsetClass         `xml:"classes>class"`
	Limitations        EbooksSubsetLimitations     `xml:"limitations"`
	Exclusions         EbooksSubsetExclusions      `xml:"exclusions"`
}

// Exclusions represent any exclusions within the scope.
type EbooksSubsetExclusions struct {
	Computers          []EbooksSubsetComputer          `xml:"computers>computer"`
	ComputerGroups     []EbooksSubsetComputerGroup     `xml:"computer_groups>computer_group"`
	MobileDevices      []EbooksSubsetMobileDevice      `xml:"mobile_devices>mobile_device"`
	MobileDeviceGroups []EbooksSubsetMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group"`
	Buildings          []EbooksSubsetBuilding          `xml:"buildings>building"`
	Departments        []EbooksSubsetDepartment        `xml:"departments>department"`
	JSSUsers           []EbooksSubsetUser              `xml:"jss_users>user"`
	JSSUserGroups      []EbooksSubsetUserGroup         `xml:"jss_user_groups>user_group"`
}

type EbooksSubsetSelfService struct {
	SelfServiceDisplayName      string `xml:"self_service_display_name"`
	InstallButtonText           string `xml:"install_button_text"`
	SelfServiceDescription      string `xml:"self_service_description"`
	ForceUsersToViewDescription bool   `xml:"force_users_to_view_description"`
	SelfServiceIcon             struct {
		ID   int    `xml:"id"`
		URI  string `xml:"uri"`
		Data string `xml:"data"`
	} `xml:"self_service_icon"`
	FeatureOnMainPage     bool                              `xml:"feature_on_main_page"`
	SelfServiceCategories EbooksSubsetSelfServiceCategories `xml:"self_service_categories"`
	Notification          bool                              `xml:"notification"`
	NotificationSubject   string                            `xml:"notification_subject"`
	NotificationMessage   string                            `xml:"notification_message"`
}

// EbooksSubsetLimitations represents any limitations within the scope.
type EbooksSubsetLimitations struct {
	NetworkSegments []struct {
		ID   int    `xml:"id"`
		UID  string `xml:"uid,omitempty"`
		Name string `xml:"name"`
	} `xml:"network_segments>network_segment"`
	Users      []EbooksSubsetUser      `xml:"users>user"`
	UserGroups []EbooksSubsetUserGroup `xml:"user_groups>user_group"`
}

// Computer represents a single computer within the scope.
type EbooksSubsetComputer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	UDID string `xml:"udid"`
}

// ComputerGroup represents a group of computers within the scope.
type EbooksSubsetComputerGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// MobileDevice represents a single mobile device within the scope.
type EbooksSubsetMobileDevice struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name"`
	UDID           string `xml:"udid"`
	WiFiMacAddress string `xml:"wifi_mac_address"`
}

// MobileDeviceGroup represents a group of mobile devices within the scope.
type MobileDeviceGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Building represents a building within the scope.
type EbooksSubsetBuilding struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Department represents a department within the scope.
type EbooksSubsetDepartment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// User represents a user within the scope.
type EbooksSubsetUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// UserGroup represents a group of users within the scope.
type EbooksSubsetUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Class represents a class within the scope.
type EbooksSubsetClass struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// SelfServiceCategories represent the categories within SelfService.
type EbooksSubsetSelfServiceCategories struct {
	Category []struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	} `xml:"category"`
}

// EbooksSubsetMobileDeviceGroup represents a group of mobile devices within the scope.
type EbooksSubsetMobileDeviceGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// GetEbooks retrieves a serialized list of ebooks.
func (c *Client) GetEbooks() (*ResponseEbooksList, error) {
	endpoint := uriEbooks

	var ebooks ResponseEbooksList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ebooks)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Ebooks: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ebooks, nil
}

// GetEbooksByID retrieves a single ebook by its ID.
func (c *Client) GetEbooksByID(id int) (*ResourceEbooks, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriEbooks, id)

	var ebook ResourceEbooks
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ebook)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Ebook by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ebook, nil
}

// GetEbooksByName retrieves a single ebook by its name.
func (c *Client) GetEbooksByName(name string) (*ResourceEbooks, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriEbooks, name)

	var ebook ResourceEbooks
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ebook)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Ebook by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ebook, nil
}

// GetEbooksByNameAndDataSubset retrieves a specific subset of an ebook by its name.
func (c *Client) GetEbooksByNameAndDataSubset(name, subset string) (*ResourceEbooks, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriEbooks, name, subset)

	var ebook ResourceEbooks
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ebook)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch Ebook by Name and Subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ebook, nil
}

// CreateEbook creates a new ebook.
func (c *Client) CreateEbook(ebook ResourceEbooks) (*ResourceEbooks, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriEbooks) // '0' typically used for creation in APIs

	// Handle default values, especially for the Site ID if not provided
	if ebook.General.Site.ID == 0 && ebook.General.Site.Name == "" {
		ebook.General.Site.ID = -1
		ebook.General.Site.Name = "none"
	}

	// The requestBody struct should mirror the ResourceEbooks struct, including all nested structs
	requestBody := struct {
		XMLName xml.Name `xml:"ebook"`
		ResourceEbooks
	}{
		ResourceEbooks: ebook,
	}

	var response ResourceEbooks
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create ebook: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateEbookByID updates an existing ebook by its ID.
func (c *Client) UpdateEbookByID(id int, ebook ResourceEbooks) (*ResourceEbooks, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriEbooks, id)

	requestBody := struct {
		XMLName xml.Name `xml:"ebook"`
		ResourceEbooks
	}{
		ResourceEbooks: ebook,
	}

	var updatedEbook ResourceEbooks
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedEbook)
	if err != nil {
		return nil, fmt.Errorf("failed to update Ebook by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedEbook, nil
}

// UpdateEbookByName updates an existing ebook by its name.
func (c *Client) UpdateEbookByName(name string, ebook ResourceEbooks) (*ResourceEbooks, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriEbooks, name)

	requestBody := struct {
		XMLName xml.Name `xml:"ebook"`
		ResourceEbooks
	}{
		ResourceEbooks: ebook,
	}

	var updatedEbook ResourceEbooks
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedEbook)
	if err != nil {
		return nil, fmt.Errorf("failed to update Ebook by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedEbook, nil
}

// DeleteEbookByID deletes a ebook by its ID.
func (c *Client) DeleteEbookByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriEbooks, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Ebook Item by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteEbookByName deletes a ebook by its name.
func (c *Client) DeleteEbookByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriEbooks, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete Ebook Item by Name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
