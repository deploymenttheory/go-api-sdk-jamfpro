// classicapi_ebooks.go
// Jamf Pro Classic Api - Ebooks
// api reference: https://developer.jamf.com/jamf-pro/reference/ebooks
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
- SharedResourceCategory
- SharedResourceSelfServiceIcon
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// URI for Ebooks in Jamf Pro API
const uriEbooks = "/JSSResource/ebooks"

// List

// Struct to capture the XML response for ebooks list
type ResponseEbooksList struct {
	Size   int           `xml:"size"`
	Ebooks EBookListItem `xml:"ebook"`
}

type EBookListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// ResourceEbooks represents the detailed structure of an Ebook response.
type ResourceEbooks struct {
	General     EbookSubsetGeneral     `xml:"general"`
	Scope       EbookSubsetScope       `xml:"scope"`
	SelfService EbookSubsetSelfService `xml:"self_service"`
}

// Subsets & Containers

// General

type EbookSubsetGeneral struct {
	ID              int                           `xml:"id"`
	Name            string                        `xml:"name"`
	Author          string                        `xml:"author"`
	Version         string                        `xml:"version"`
	Free            bool                          `xml:"free"`
	URL             string                        `xml:"url"`
	DeploymentType  string                        `xml:"deployment_type"`
	FileType        string                        `xml:"file_type"`
	DeployAsManaged bool                          `xml:"deploy_as_managed"`
	Category        *SharedResourceCategory       `xml:"category"`
	SelfServiceIcon SharedResourceSelfServiceIcon `xml:"self_service_icon"`
	Site            SharedResourceSite            `xml:"site"`
}

// Scope

type EbookSubsetScope struct {
	AllComputers       bool                                `xml:"all_computers"`
	AllMobileDevices   bool                                `xml:"all_mobile_devices"`
	AllJSSUsers        bool                                `xml:"all_jss_users"`
	Computers          []EbookSubsetScopeComputer          `xml:"computers>computer"`
	ComputerGroups     []EbookSubsetScopeComputerGroup     `xml:"computer_groups>computer_group"`
	MobileDevices      []EbookSubsetScopeMobileDevice      `xml:"mobile_devices>mobile_device"`
	MobileDeviceGroups []EbookSubsetScopeMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group"`
	Buildings          []EbookSubsetScopeBuilding          `xml:"buildings>building"`
	Departments        []EbookSubsetScopeDepartment        `xml:"departments>department"`
	JSSUsers           []EbookSubsetScopeUser              `xml:"jss_users>user"`
	JSSUserGroups      []EbookSubsetScopeUserGroup         `xml:"jss_user_groups>user_group"`
	Classes            []EbooksSubsetScopeClass            `xml:"classes>class"`
	Limitations        EbookSubsetScopeLimitations         `xml:"limitations"`
	Exclusions         EbookSubsetScopeExclusions          `xml:"exclusions"`
}

// EbooksSubsetLimitations represents any limitations within the scope.
type EbookSubsetScopeLimitations struct {
	NetworkSegments []struct {
		ID   int    `xml:"id"`
		UID  string `xml:"uid,omitempty"`
		Name string `xml:"name"`
	} `xml:"network_segments>network_segment"`
	Users      []EbookSubsetScopeUser      `xml:"users>user"`
	UserGroups []EbookSubsetScopeUserGroup `xml:"user_groups>user_group"`
}

// Exclusions represent any exclusions within the scope.
type EbookSubsetScopeExclusions struct {
	Computers          []EbookSubsetScopeComputer          `xml:"computers>computer"`
	ComputerGroups     []EbookSubsetScopeComputerGroup     `xml:"computer_groups>computer_group"`
	MobileDevices      []EbookSubsetScopeMobileDevice      `xml:"mobile_devices>mobile_device"`
	MobileDeviceGroups []EbookSubsetScopeMobileDeviceGroup `xml:"mobile_device_groups>mobile_device_group"`
	Buildings          []EbookSubsetScopeBuilding          `xml:"buildings>building"`
	Departments        []EbookSubsetScopeDepartment        `xml:"departments>department"`
	JSSUsers           []EbookSubsetScopeUser              `xml:"jss_users>user"`
	JSSUserGroups      []EbookSubsetScopeUserGroup         `xml:"jss_user_groups>user_group"`
}

// Class represents a class within the scope.
type EbooksSubsetScopeClass struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Self Service

type EbookSubsetSelfService struct {
	SelfServiceDisplayName      string                           `xml:"self_service_display_name"`
	InstallButtonText           string                           `xml:"install_button_text"`
	SelfServiceDescription      string                           `xml:"self_service_description"`
	ForceUsersToViewDescription bool                             `xml:"force_users_to_view_description"`
	SelfServiceIcon             SharedResourceSelfServiceIcon    `xml:"self_service_icon"`
	FeatureOnMainPage           bool                             `xml:"feature_on_main_page"`
	SelfServiceCategories       EbookSubsetSelfServiceCategories `xml:"self_service_categories"`
	Notification                bool                             `xml:"notification"`
	NotificationSubject         string                           `xml:"notification_subject"`
	NotificationMessage         string                           `xml:"notification_message"`
}

// SelfServiceCategories represent the categories within SelfService.
type EbookSubsetSelfServiceCategories struct {
	Category []struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	} `xml:"category"`
}

/// Shared In Resource

// Computer represents a single computer within the scope.
type EbookSubsetScopeComputer struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
	UDID string `xml:"udid"`
}

// ComputerGroup represents a group of computers within the scope.
type EbookSubsetScopeComputerGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// MobileDevice represents a single mobile device within the scope.
type EbookSubsetScopeMobileDevice struct {
	ID             int    `xml:"id"`
	Name           string `xml:"name"`
	UDID           string `xml:"udid"`
	WiFiMacAddress string `xml:"wifi_mac_address"`
}

// EbooksSubsetMobileDeviceGroup represents a group of mobile devices within the scope.
type EbookSubsetScopeMobileDeviceGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Building represents a building within the scope.
type EbookSubsetScopeBuilding struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Department represents a department within the scope.
type EbookSubsetScopeDepartment struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// User represents a user within the scope.
type EbookSubsetScopeUser struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// UserGroup represents a group of users within the scope.
type EbookSubsetScopeUserGroup struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// CRUD

// GetEbooks retrieves a serialized list of ebooks.
func (c *Client) GetEbooks() (*ResponseEbooksList, error) {
	endpoint := uriEbooks

	var ebooks ResponseEbooksList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ebooks)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "ebooks", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ebooks, nil
}

// GetEbooksByID retrieves a single ebook by its ID.
func (c *Client) GetEbookByID(id int) (*ResourceEbooks, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriEbooks, id)

	var ebook ResourceEbooks
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ebook)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "ebook", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ebook, nil
}

// GetEbooksByName retrieves a single ebook by its name.
func (c *Client) GetEbookByName(name string) (*ResourceEbooks, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriEbooks, name)

	var ebook ResourceEbooks
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ebook)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "ebook", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ebook, nil
}

// GetEbooksByNameAndDataSubset retrieves a specific subset of an ebook by its name.
func (c *Client) GetEbookByNameAndDataSubset(name, subset string) (*ResourceEbooks, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriEbooks, name, subset)

	var ebook ResourceEbooks
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &ebook)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "ebook", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ebook, nil
}

// CreateEbook creates a new ebook.
func (c *Client) CreateEbook(ebook ResourceEbooks) (*ResourceEbooks, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriEbooks)

	requestBody := struct {
		XMLName xml.Name `xml:"ebook"`
		ResourceEbooks
	}{
		ResourceEbooks: ebook,
	}

	var response ResourceEbooks
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "ebook", err)
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
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "ebook", id, err)
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
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "ebook", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedEbook, nil
}

// DeleteEbookByID deletes a ebook by its ID.
func (c *Client) DeleteEbookByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriEbooks, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "ebook", id, err)
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
		return fmt.Errorf(errMsgFailedDeleteByName, "ebook", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
