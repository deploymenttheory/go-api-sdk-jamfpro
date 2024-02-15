// classicapi_packages.go
// Jamf Pro Classic Api  - Packages
// api reference: https://developer.jamf.com/jamf-pro/reference/packages
// Jamf Pro Classic Api requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSelfServiceIcon
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

// URI for Packages in the Jamf Pro Classic API
const uriPackages = "/JSSResource/packages"

// ResponsePackagesList struct to capture the XML response for packages list
type ResponsePackagesList struct {
	Size    int             `xml:"size"`    // The size attribute
	Package PackageListItem `xml:"package"` // The package element
}

// PackageListItem struct to capture individual package items in the list
type PackageListItem struct {
	ID   int    `xml:"id"`   // The ID element
	Name string `xml:"name"` // The Name element
}

// Response

type ResponsePackageCreatedAndUpdated struct {
	ID int `json:"id,omitempty" xml:"id,omitempty"`
}

// ResourcePackage struct to capture the XML response for a single package detail
type ResourcePackage struct {
	ID                         int    `xml:"id"`
	Name                       string `xml:"name"`
	Category                   string `xml:"category"`
	Filename                   string `xml:"filename"`
	Info                       string `xml:"info"`
	Notes                      string `xml:"notes"`
	Priority                   int    `xml:"priority"`
	RebootRequired             bool   `xml:"reboot_required"`
	FillUserTemplate           bool   `xml:"fill_user_template"`
	FillExistingUsers          bool   `xml:"fill_existing_users"`
	BootVolumeRequired         bool   `xml:"boot_volume_required"`
	AllowUninstalled           bool   `xml:"allow_uninstalled"`
	OSRequirements             string `xml:"os_requirements"`
	RequiredProcessor          string `xml:"required_processor"`
	SwitchWithPackage          string `xml:"switch_with_package"`
	InstallIfReportedAvailable bool   `xml:"install_if_reported_available"`
	ReinstallOption            string `xml:"reinstall_option"`
	TriggeringFiles            string `xml:"triggering_files"`
	SendNotification           bool   `xml:"send_notification"`
}

// CRUD

// GetPackages retrieves a list of packages.
func (c *Client) GetPackages() (*ResponsePackagesList, error) {
	endpoint := uriPackages

	var response ResponsePackagesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "packages", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetPackageByID retrieves details of a specific package by its ID.
func (c *Client) GetPackageByID(id int) (*ResourcePackage, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPackages, id)

	var response ResourcePackage
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "packages", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetPackageByID retrieves details of a specific package by its ID.
func (c *Client) GetPackageByName(name string) (*ResourcePackage, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPackages, name)

	var response ResourcePackage
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "packages", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// CreatePackage creates a new package in Jamf Pro
func (c *Client) CreatePackage(pkg ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriPackages)

	requestBody := struct {
		XMLName xml.Name `xml:"package"`
		ResourcePackage
	}{
		ResourcePackage: pkg,
	}

	var response ResponsePackageCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create package: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
