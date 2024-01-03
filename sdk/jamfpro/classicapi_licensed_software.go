// classicapi_licensed_software.go
// Jamf Pro Classic Api - Licensed Software
// api reference: https://developer.jamf.com/jamf-pro/reference/licensedsoftware
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriLicensedSoftware = "/JSSResource/licensedsoftware"

// List

// ResponseLicensedSoftwareList represents the response for a list of licensed software.
type ResponseLicensedSoftwareList struct {
	LicensedSoftware []LicensedSoftwareListItem `xml:"licensed_software"`
}

type LicensedSoftwareListItem struct {
}

// Resource

// ResourceLicensedSoftware represents the structure of a single licensed software item.
type ResourceLicensedSoftware struct {
	General             LicensedSoftwareSubsetGeneral               `xml:"general"`
	SoftwareDefinitions []LicensedSoftwareSubsetSoftwareDefinitions `xml:"software_definitions>definition"`
	FontDefinitions     []LicensedSoftwareSubsetFontDefinitions     `xml:"font_definitions>definition"`
	PluginDefinitions   []LicensedSoftwareSubsetPluginDefinitions   `xml:"plugin_definitions>definition"`
	Licenses            []LicensedSoftwareSubsetLicenses            `xml:"licenses>license"`
}

// Subsets & Containers

type LicensedSoftwareSubsetGeneral struct {
	ID                                 int                `xml:"id"`
	Name                               string             `xml:"name"`
	Publisher                          string             `xml:"publisher"`
	Platform                           string             `xml:"platform"`
	SendEmailOnViolation               bool               `xml:"send_email_on_violation"`
	RemoveTitlesFromInventoryReports   bool               `xml:"remove_titles_from_inventory_reports"`
	ExcludeTitlesPurchasedFromAppStore bool               `xml:"exclude_titles_purchased_from_app_store"`
	Notes                              string             `xml:"notes"`
	Site                               SharedResourceSite `xml:"site"`
}

type LicensedSoftwareSubsetSoftwareDefinitions struct {
	CompareType string `xml:"compare_type"`
	Name        string `xml:"name"`
	Version     int    `xml:"version"`
}

type LicensedSoftwareSubsetFontDefinitions struct {
	CompareType string `xml:"compare_type"`
	Name        string `xml:"name"`
	Version     int    `xml:"version"`
}

type LicensedSoftwareSubsetPluginDefinitions struct {
	CompareType string `xml:"compare_type"`
	Name        string `xml:"name"`
	Version     int    `xml:"version"`
}

type LicensedSoftwareSubsetLicenses struct {
	Size    int                           `xml:"size"`
	License LicensedSoftwareSubsetLicense `xml:"license"`
}

type LicensedSoftwareSubsetLicense struct {
	SerialNumber1    string                                     `xml:"serial_number_1"`
	SerialNumber2    string                                     `xml:"serial_number_2"`
	OrganizationName string                                     `xml:"organization_name"`
	RegisteredTo     string                                     `xml:"registered_to"`
	LicenseType      string                                     `xml:"license_type"`
	LicenseCount     int                                        `xml:"license_count"`
	Notes            string                                     `xml:"notes"`
	Purchasing       LicensedSoftwareSubsetLicensePurchasing    `xml:"purchasing"`
	Attachments      []LicensedSoftwareSubsetLicenseAttachments `xml:"attachments>attachment"`
}

type LicensedSoftwareSubsetLicensePurchasing struct {
	IsPerpetual         bool   `xml:"is_perpetual"`
	IsAnnual            bool   `xml:"is_annual"`
	PONumber            string `xml:"po_number"`
	Vendor              string `xml:"vendor"`
	PurchasePrice       string `xml:"purchase_price"`
	PurchasingAccount   string `xml:"purchasing_account"`
	PODate              string `xml:"po_date"`
	PODateEpoch         int64  `xml:"po_date_epoch"`
	PODateUTC           string `xml:"po_date_utc"`
	LicenseExpires      string `xml:"license_expires"`
	LicenseExpiresEpoch int64  `xml:"license_expires_epoch"`
	LicenseExpiresUTC   string `xml:"license_expires_utc"`
	LifeExpectancy      int    `xml:"life_expectancy"`
	PurchasingContact   string `xml:"purchasing_contact"`
}

type LicensedSoftwareSubsetLicenseAttachments struct {
	ID       int    `xml:"id"`
	Filename string `xml:"filename"`
	URI      string `xml:"uri"`
}

// CRUD

// GetLicensedSoftware retrieves a serialized list of licensed software.
func (c *Client) GetLicensedSoftware() (*ResponseLicensedSoftwareList, error) {
	endpoint := uriLicensedSoftware

	var licensedSoftware ResponseLicensedSoftwareList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &licensedSoftware)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch licensed software: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &licensedSoftware, nil
}

// GetLicensedSoftwareByID retrieves details of a specific licensed software by its ID.
func (c *Client) GetLicensedSoftwareByID(id int) (*ResourceLicensedSoftware, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriLicensedSoftware, id)

	var licensedSoftware ResourceLicensedSoftware
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &licensedSoftware)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch licensed software by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &licensedSoftware, nil
}

// GetLicensedSoftwareByName retrieves details of a specific licensed software by its name.
func (c *Client) GetLicensedSoftwareByName(name string) (*ResourceLicensedSoftware, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriLicensedSoftware, name)

	var licensedSoftware ResourceLicensedSoftware
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &licensedSoftware)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch licensed software by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &licensedSoftware, nil
}

// CreateLicensedSoftware creates a new licensed software item in Jamf Pro.
func (c *Client) CreateLicensedSoftware(licensedSoftware *ResourceLicensedSoftware) (*ResourceLicensedSoftware, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriLicensedSoftware) // '0' typically used for creation in APIs

	// Set default values for site if not included within request
	if licensedSoftware.General.Site.ID == 0 && licensedSoftware.General.Site.Name == "" {
		licensedSoftware.General.Site.ID = -1
		licensedSoftware.General.Site.Name = "none"
	}

	// Wrap licensedSoftware in an anonymous struct to match the expected XML structure
	requestBody := struct {
		XMLName xml.Name `xml:"licensed_software"`
		*ResourceLicensedSoftware
	}{
		ResourceLicensedSoftware: licensedSoftware,
	}

	var ResourceLicensedSoftware ResourceLicensedSoftware
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &ResourceLicensedSoftware)
	if err != nil {
		return nil, fmt.Errorf("failed to create licensed software: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ResourceLicensedSoftware, nil
}

// UpdateLicensedSoftwareByID updates an existing licensed software item by its ID.
func (c *Client) UpdateLicensedSoftwareByID(id int, licensedSoftware *ResourceLicensedSoftware) (*ResourceLicensedSoftware, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriLicensedSoftware, id)

	requestBody := struct {
		XMLName xml.Name `xml:"licensed_software"`
		*ResourceLicensedSoftware
	}{
		ResourceLicensedSoftware: licensedSoftware,
	}

	var ResourceLicensedSoftware ResourceLicensedSoftware
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &ResourceLicensedSoftware)
	if err != nil {
		return nil, fmt.Errorf("failed to update licensed software by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ResourceLicensedSoftware, nil
}

// UpdateLicensedSoftwareByName updates an existing licensed software item by its name.
func (c *Client) UpdateLicensedSoftwareByName(name string, licensedSoftware *ResourceLicensedSoftware) (*ResourceLicensedSoftware, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriLicensedSoftware, name)

	requestBody := struct {
		XMLName xml.Name `xml:"licensed_software"`
		*ResourceLicensedSoftware
	}{
		ResourceLicensedSoftware: licensedSoftware,
	}

	var ResourceLicensedSoftware ResourceLicensedSoftware
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &ResourceLicensedSoftware)
	if err != nil {
		return nil, fmt.Errorf("failed to update licensed software by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &ResourceLicensedSoftware, nil
}

// DeleteLicensedSoftwareByID deletes a licensed software item by its ID.
func (c *Client) DeleteLicensedSoftwareByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriLicensedSoftware, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete licensed software by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteLicensedSoftwareByName deletes a licensed software item by its name.
func (c *Client) DeleteLicensedSoftwareByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriLicensedSoftware, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete licensed software by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
