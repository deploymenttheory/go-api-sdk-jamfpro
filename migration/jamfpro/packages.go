// packages.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIPackages = "/JSSResource/packages"

type ResponsePackage struct {
	ID                         int    `json:"id,omitempty" xml:"id,omitempty"`
	Name                       string `json:"name" xml:"name"`
	Category                   string `json:"category,omitempty" xml:"category,omitempty"`
	Filename                   string `json:"filename,omitempty" xml:"filename,omitempty"`
	Info                       string `json:"info,omitempty" xml:"info,omitempty"`
	Notes                      string `json:"notes,omitempty" xml:"notes,omitempty"`
	Priority                   int    `json:"priority,omitempty" xml:"priority,omitempty"`
	RebootRequired             bool   `json:"reboot_required,omitempty" xml:"reboot_required,omitempty"`
	FillUserTemplate           bool   `json:"fill_user_template,omitempty" xml:"fill_user_template,omitempty"`
	FillExistingUsers          bool   `json:"fill_existing_users,omitempty" xml:"fill_existing_users,omitempty"`
	AllowUninstalled           bool   `json:"allow_uninstalled,omitempty" xml:"allow_uninstalled,omitempty"`
	OSRequirements             string `json:"os_requirements,omitempty" xml:"os_requirements,omitempty"`
	RequiredProcessor          string `json:"required_processor,omitempty" xml:"required_processor,omitempty"`
	SwitchWithPackage          string `json:"switch_with_package,omitempty" xml:"switch_with_package,omitempty"`
	InstallIfReportedAvailable bool   `json:"install_if_reported_available,omitempty" xml:"install_if_reported_available,omitempty"`
	ReinstallOption            string `json:"reinstall_option,omitempty" xml:"reinstall_option,omitempty"`
	TriggeringFiles            string `json:"triggering_files,omitempty" xml:"triggering_files,omitempty"`
	SendNotification           bool   `json:"send_notification,omitempty" xml:"send_notification,omitempty"`
}

type ResponsePackagesList struct {
	Size     int             `json:"size" xml:"size"`
	Packages []PackageDetail `json:"package" xml:"package"`
}

type PackageDetail struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// GetPackageByID gets a package by it's id
func (c *Client) GetPackageByID(id int) (*ResponsePackage, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIPackages, id)

	var pkg ResponsePackage
	if err := c.DoRequest("GET", url, nil, nil, &pkg); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &pkg, nil
}

// GetPackageByName gets a package by it's name
func (c *Client) GetPackageByName(name string) (*ResponsePackage, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIPackages, name)

	var pkg ResponsePackage
	if err := c.DoRequest("GET", url, nil, nil, &pkg); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &pkg, nil
}

// GetPackages gets a list of all packages
func (c *Client) GetPackages() ([]ResponsePackagesList, error) {
	url := uriAPIPackages

	var packagesList []ResponsePackagesList
	if err := c.DoRequest("GET", url, nil, nil, &packagesList); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return packagesList, nil
}

// GetPackageNameByID retrieves the name of a package by its ID.
func (c *Client) GetPackageNameByID(id int) (string, error) {
	packagesList, err := c.GetPackages()
	if err != nil {
		return "", fmt.Errorf("failed to retrieve packages: %v", err)
	}

	for _, pkgList := range packagesList {
		for _, pkg := range pkgList.Packages {
			if pkg.ID == id {
				return pkg.Name, nil
			}
		}
	}

	return "", fmt.Errorf("package with ID %d not found", id)
}

// GetPackageIDByName retrieves the ID of a package by its name.
func (c *Client) GetPackageIDByName(name string) (int, error) {
	packagesList, err := c.GetPackages()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve packages: %v", err)
	}

	for _, pkgList := range packagesList {
		for _, pkg := range pkgList.Packages {
			if pkg.Name == name {
				return pkg.ID, nil
			}
		}
	}

	return 0, fmt.Errorf("package with name %s not found", name)
}

// CreatePackage creates a new Jamf Pro Package.
func (c *Client) CreatePackage(pkg *ResponsePackage) (*ResponsePackage, error) {
	url := fmt.Sprintf("%s/id/0", uriAPIPackages) // ID 0 is typically used for creation in many APIs

	reqBody := &struct {
		XMLName xml.Name `xml:"package"`
		*ResponsePackage
	}{
		ResponsePackage: pkg,
	}

	var responsePackage ResponsePackage
	if err := c.DoRequest("POST", url, reqBody, nil, &responsePackage); err != nil {
		return nil, fmt.Errorf("failed to create package: %v", err)
	}

	return &responsePackage, nil
}

// UpdatePackageByID updates an existing Jamf Pro Package by ID.
func (c *Client) UpdatePackageByID(id int, pkg *ResponsePackage) (*ResponsePackage, error) {
	url := fmt.Sprintf("%s/id/%d", uriAPIPackages, id)

	reqBody := &struct {
		XMLName xml.Name `xml:"package"`
		*ResponsePackage
	}{
		ResponsePackage: pkg,
	}

	var responsePackage ResponsePackage
	if err := c.DoRequest("PUT", url, reqBody, nil, &responsePackage); err != nil {
		return nil, fmt.Errorf("failed to update package by ID: %v", err)
	}

	return &responsePackage, nil
}

// UpdatePackageByName updates an existing Jamf Pro Package by Name.
func (c *Client) UpdatePackageByName(name string, pkg *ResponsePackage) (*ResponsePackage, error) {
	url := fmt.Sprintf("%s/name/%s", uriAPIPackages, name)

	reqBody := &struct {
		XMLName xml.Name `xml:"package"`
		*ResponsePackage
	}{
		ResponsePackage: pkg,
	}

	var responsePackage ResponsePackage
	if err := c.DoRequest("PUT", url, reqBody, nil, &responsePackage); err != nil {
		return nil, fmt.Errorf("failed to update package by Name: %v", err)
	}

	return &responsePackage, nil
}

// DeletePackageByID deletes an existing Jamf Pro Package by ID.
func (c *Client) DeletePackageByID(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriAPIPackages, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil, c.HTTP.Logger); err != nil {
		return fmt.Errorf("failed to delete package by ID: %v", err)
	}

	return nil
}

// DeletePackageByName deletes an existing Jamf Pro Package by Name.
func (c *Client) DeletePackageByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriAPIPackages, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil, c.HTTP.Logger); err != nil {
		return fmt.Errorf("failed to delete package by Name: %v", err)
	}

	return nil
}
