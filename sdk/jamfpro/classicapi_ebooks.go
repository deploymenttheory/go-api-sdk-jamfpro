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
	Size  int   `xml:"size"`
	Ebook Ebook `xml:"ebook"`
}

type Ebook struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResponseEbooks represents the detailed structure of an Ebook response.
type ResponseEbooks struct {
	General     EbooksDataSubsetGeneral     `xml:"general"`
	Scope       EbooksDataSubsetScope       `xml:"scope"`
	SelfService EbooksDataSubsetSelfService `xml:"self_service"`
}

type EbooksDataSubsetGeneral struct {
	ID              int                             `xml:"id"`
	Name            string                          `xml:"name"`
	Author          string                          `xml:"author"`
	Version         string                          `xml:"version"`
	Free            bool                            `xml:"free"`
	URL             string                          `xml:"url"`
	DeploymentType  string                          `xml:"deployment_type"`
	FileType        string                          `xml:"file_type"`
	DeployAsManaged bool                            `xml:"deploy_as_managed"`
	Category        EbooksDataSubsetCategory        `xml:"category"`
	SelfServiceIcon EbooksDataSubsetSelfServiceIcon `xml:"self_service_icon"`
	Site            EbooksDataSubsetSite            `xml:"site"`
}

type EbooksDataSubsetCategory struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type EbooksDataSubsetSelfServiceIcon struct {
	ID   int    `xml:"id"`
	URI  string `xml:"uri"`
	Data string `xml:"data"`
}

type EbooksDataSubsetSite struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

type EbooksDataSubsetScope struct {
	AllComputers       bool                            `xml:"all_computers"`
	AllMobileDevices   bool                            `xml:"all_mobile_devices"`
	AllJSSUsers        bool                            `xml:"all_jss_users"`
	Computers          []EbooksDataSubsetComputer      `xml:"computers>computer"`
	ComputerGroups     []EbooksDataSubsetComputerGroup `xml:"computer_groups>computer_group"`
	MobileDevices      []EbooksDataSubsetMobileDevice  `xml:"mobile_devices>mobile_device"`
	MobileDeviceGroups []MobileDeviceGroup             `xml:"mobile_device_groups>mobile_device_group"`
	Buildings          []EbooksDataSubsetBuilding      `xml:"buildings>building"`
	Departments        []EbooksDataSubsetDepartment    `xml:"departments>department"`
	JSSUsers           []EbooksDataSubsetUser          `xml:"jss_users>user"`
	JSSUserGroups      []EbooksDataSubsetUserGroup     `xml:"jss_user_groups>user_group"`
	Classes            []EbooksDataSubsetClass         `xml:"classes>class"`
	Limitations        EbooksDataSubsetLimitations     `xml:"limitations"`
	Exclusions         EbooksDataSubsetExclusions      `xml:"exclusions"`
}

// Computer represents a single computer within the scope.
type EbooksDataSubsetComputer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	UDID string `xml:"udid"`
}

// ComputerGroup represents a group of computers within the scope.
type EbooksDataSubsetComputerGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// MobileDevice represents a single mobile device within the scope.
type EbooksDataSubsetMobileDevice struct {
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
type EbooksDataSubsetBuilding struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Department represents a department within the scope.
type EbooksDataSubsetDepartment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// User represents a user within the scope.
type EbooksDataSubsetUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// UserGroup represents a group of users within the scope.
type EbooksDataSubsetUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Class represents a class within the scope.
type EbooksDataSubsetClass struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Exclusions represent any exclusions within the scope.
type EbooksDataSubsetExclusions struct {
	Computers          []EbooksDataSubsetComputer          `xml:"computers>computer"`
	ComputerGroups     []EbooksDataSubsetComputerGroup     `xml:"computer_groups>computer_group"`
	MobileDevices      []EbooksDataSubsetMobileDevice      `xml:"mobile_devices>mobile_device"`
	MobileDeviceGroups []EbooksDataSubsetMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group"`
	Buildings          []EbooksDataSubsetBuilding          `xml:"buildings>building"`
	Departments        []EbooksDataSubsetDepartment        `xml:"departments>department"`
	JSSUsers           []EbooksDataSubsetUser              `xml:"jss_users>user"`
	JSSUserGroups      []EbooksDataSubsetUserGroup         `xml:"jss_user_groups>user_group"`
}

// SelfServiceCategories represent the categories within SelfService.
type EbooksDataSubsetSelfServiceCategories struct {
	Category []EbooksDataSubsetCategory `xml:"category"`
}

type EbooksDataSubsetSelfService struct {
	SelfServiceDisplayName      string                                `xml:"self_service_display_name"`
	InstallButtonText           string                                `xml:"install_button_text"`
	SelfServiceDescription      string                                `xml:"self_service_description"`
	ForceUsersToViewDescription bool                                  `xml:"force_users_to_view_description"`
	SelfServiceIcon             EbooksDataSubsetSelfServiceIcon       `xml:"self_service_icon"`
	FeatureOnMainPage           bool                                  `xml:"feature_on_main_page"`
	SelfServiceCategories       EbooksDataSubsetSelfServiceCategories `xml:"self_service_categories"`
	Notification                bool                                  `xml:"notification"`
	NotificationSubject         string                                `xml:"notification_subject"`
	NotificationMessage         string                                `xml:"notification_message"`
}

// EbooksDataSubsetMobileDeviceGroup represents a group of mobile devices within the scope.
type EbooksDataSubsetMobileDeviceGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// EbooksDataSubsetNetworkSegment represents a network segment within the limitations.
type EbooksDataSubsetNetworkSegment struct {
	ID   int    `xml:"id"`
	UID  string `xml:"uid,omitempty"`
	Name string `xml:"name"`
}

// EbooksDataSubsetLimitations represents any limitations within the scope.
type EbooksDataSubsetLimitations struct {
	NetworkSegments []EbooksDataSubsetNetworkSegment `xml:"network_segments>network_segment"`
	Users           []EbooksDataSubsetUser           `xml:"users>user"`
	UserGroups      []EbooksDataSubsetUserGroup      `xml:"user_groups>user_group"`
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
func (c *Client) GetEbooksByID(id int) (*ResponseEbooks, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriEbooks, id)

	var ebook ResponseEbooks
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
func (c *Client) GetEbooksByName(name string) (*ResponseEbooks, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriEbooks, name)

	var ebook ResponseEbooks
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
func (c *Client) GetEbooksByNameAndDataSubset(name, subset string) (*ResponseEbooks, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriEbooks, name, subset)

	var ebook ResponseEbooks
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
func (c *Client) CreateEbook(ebook ResponseEbooks) (*ResponseEbooks, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriEbooks) // '0' typically used for creation in APIs

	// Handle default values, especially for the Site ID if not provided
	if ebook.General.Site.ID == 0 && ebook.General.Site.Name == "" {
		ebook.General.Site = EbooksDataSubsetSite{
			ID:   -1,
			Name: "None",
		}
	}

	// The requestBody struct should mirror the ResponseEbooks struct, including all nested structs
	requestBody := struct {
		XMLName xml.Name `xml:"ebook"`
		ResponseEbooks
	}{
		ResponseEbooks: ebook,
	}

	var response ResponseEbooks
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
func (c *Client) UpdateEbookByID(id int, ebook ResponseEbooks) (*ResponseEbooks, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriEbooks, id)

	requestBody := struct {
		XMLName xml.Name `xml:"ebook"`
		ResponseEbooks
	}{
		ResponseEbooks: ebook,
	}

	var updatedEbook ResponseEbooks
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
func (c *Client) UpdateEbookByName(name string, ebook ResponseEbooks) (*ResponseEbooks, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriEbooks, name)

	requestBody := struct {
		XMLName xml.Name `xml:"ebook"`
		ResponseEbooks
	}{
		ResponseEbooks: ebook,
	}

	var updatedEbook ResponseEbooks
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
