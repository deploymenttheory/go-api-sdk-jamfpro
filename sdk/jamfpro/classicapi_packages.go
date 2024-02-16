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
	Size    int               `xml:"size"`    // The size attribute
	Package []PackageListItem `xml:"package"` // The package element
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
	Category                   string `xml:"category,omitempty"`
	Filename                   string `xml:"filename,omitempty"`
	Info                       string `xml:"info,omitempty"`
	Notes                      string `xml:"notes,omitempty"`
	Priority                   int    `xml:"priority,omitempty"`
	RebootRequired             bool   `xml:"reboot_required,omitempty"`
	FillUserTemplate           bool   `xml:"fill_user_template,omitempty"`
	FillExistingUsers          bool   `xml:"fill_existing_users,omitempty"`
	BootVolumeRequired         bool   `xml:"boot_volume_required,omitempty"`
	AllowUninstalled           bool   `xml:"allow_uninstalled,omitempty"`
	OSRequirements             string `xml:"os_requirements,omitempty"`
	RequiredProcessor          string `xml:"required_processor,omitempty"`
	SwitchWithPackage          string `xml:"switch_with_package,omitempty"`
	InstallIfReportedAvailable bool   `xml:"install_if_reported_available,omitempty"`
	ReinstallOption            string `xml:"reinstall_option,omitempty"`
	TriggeringFiles            string `xml:"triggering_files,omitempty"`
	SendNotification           bool   `xml:"send_notification,omitempty"`
}

// CRUD

// GetPackages retrieves a list of packages.
func (c *Client) GetPackages() (*ResponsePackagesList, error) {
	endpoint := uriPackages

	var response ResponsePackagesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "package", err)
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
		return nil, fmt.Errorf(errMsgFailedGetByID, "package", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetPackageByName retrieves details of a specific package by its name.
func (c *Client) GetPackageByName(name string) (*ResourcePackage, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPackages, name)

	var response ResourcePackage
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "package", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// CreatePackage creates a new package in Jamf Pro
func (c *Client) CreatePackage(pkg ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPackages, pkg.ID)

	requestBody := struct {
		XMLName xml.Name `xml:"package"`
		ResourcePackage
	}{
		ResourcePackage: pkg,
	}

	var response ResponsePackageCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "package", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdatePackageByID updates an existing package by its ID on the Jamf Pro server
// and returns the response with the ID of the updated package.
func (c *Client) UpdatePackageByID(id int, pkg *ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriPackages, id)

	requestBody := struct {
		XMLName xml.Name `xml:"package"`
		*ResourcePackage
	}{
		ResourcePackage: pkg,
	}

	var response ResponsePackageCreatedAndUpdated

	// Use PUT method for updating the package
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "package", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdatePackageByName updates an existing package by its ID on the Jamf Pro server
// and returns the response with the ID of the updated package.
func (c *Client) UpdatePackageByName(name string, pkg *ResourcePackage) (*ResponsePackageCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriPackages, name)

	requestBody := struct {
		XMLName xml.Name `xml:"package"`
		*ResourcePackage
	}{
		ResourcePackage: pkg,
	}

	var response ResponsePackageCreatedAndUpdated

	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "package", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeletePackageByID deletes a package by its ID from the Jamf Pro server.
func (c *Client) DeletePackageByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriPackages, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "package", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeletePackageByName deletes a package by its name from the Jamf Pro server.
func (c *Client) DeletePackageByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriPackages, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "package", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
