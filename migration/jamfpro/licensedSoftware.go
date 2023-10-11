// licensedSoftware.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
)

const uriLicensedSoftware = "/JSSResource/licensedsoftware"

type ResponseLicensedSoftware struct {
	General             LicensedSoftwareGeneral                `json:"general" xml:"general"`
	SoftwareDefinitions []LicensedSoftwareDataSubsetDefinition `json:"software_definitions" xml:"software_definitions>definition"`
	FontDefinitions     []LicensedSoftwareDataSubsetDefinition `json:"font_definitions" xml:"font_definitions>definition"`
	PluginDefinitions   []LicensedSoftwareDataSubsetDefinition `json:"plugin_definitions" xml:"plugin_definitions>definition"`
	Licenses            []LicensedSoftwareDataSubsetLicense    `json:"licenses" xml:"licenses>license"`
}

type LicensedSoftwareGeneral struct {
	ID                                 int    `json:"id,omitempty" xml:"id,omitempty"`
	Name                               string `json:"name" xml:"name"`
	Publisher                          string `json:"publisher" xml:"publisher"`
	Platform                           string `json:"platform" xml:"platform"`
	SendEmailOnViolation               bool   `json:"send_email_on_violation" xml:"send_email_on_violation"`
	RemoveTitlesFromInventoryReports   bool   `json:"remove_titles_from_inventory_reports" xml:"remove_titles_from_inventory_reports"`
	ExcludeTitlesPurchasedFromAppStore bool   `json:"exclude_titles_purchased_from_app_store" xml:"exclude_titles_purchased_from_app_store"`
	Notes                              string `json:"notes" xml:"notes"`
	Site                               Site   `json:"site" xml:"site"`
}

type LicensedSoftwareDataSubsetDefinition struct {
	CompareType string `json:"compare_type" xml:"compare_type"`
	Name        string `json:"name" xml:"name"`
	Version     int    `json:"version" xml:"version"`
}

type LicensedSoftwareDataSubsetLicense struct {
	SerialNumber1    string                               `json:"serial_number_1" xml:"serial_number_1"`
	SerialNumber2    string                               `json:"serial_number_2" xml:"serial_number_2"`
	OrganizationName string                               `json:"organization_name" xml:"organization_name"`
	RegisteredTo     string                               `json:"registered_to" xml:"registered_to"`
	LicenseType      string                               `json:"license_type" xml:"license_type"`
	LicenseCount     int                                  `json:"license_count" xml:"license_count"`
	Notes            string                               `json:"notes" xml:"notes"`
	Purchasing       LicensedSoftwareDataSubsetPurchasing `json:"purchasing" xml:"purchasing"`
}

type LicensedSoftwareDataSubsetPurchasing struct {
	IsPerpetual       bool   `json:"is_perpetual" xml:"is_perpetual"`
	IsAnnual          bool   `json:"is_annual" xml:"is_annual"`
	PONumber          string `json:"po_number" xml:"po_number"`
	Vendor            string `json:"vendor" xml:"vendor"`
	PurchasePrice     string `json:"purchase_price" xml:"purchase_price"`
	PurchasingAccount string `json:"purchasing_account" xml:"purchasing_account"`
	PODate            string `json:"po_date" xml:"po_date"`
	LicenseExpires    string `json:"license_expires" xml:"license_expires"`
	LifeExpectancy    int    `json:"life_expectancy" xml:"life_expectancy"`
	PurchasingContact string `json:"purchasing_contact" xml:"purchasing_contact"`
}

type ResponseLicensedSoftwareList struct {
	SoftwareItems []LicensedSoftwareListItem `json:"licensed_software" xml:"licensed_software"`
}

type LicensedSoftwareListItem struct {
	ID   int    `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

// GetLicensedSoftwareByID retrieves the Licensed Software by its ID
func (c *Client) GetLicensedSoftwareByID(id int) (*ResponseLicensedSoftware, error) {
	url := fmt.Sprintf("%s/id/%d", uriLicensedSoftware, id)

	var software ResponseLicensedSoftware
	if err := c.DoRequest("GET", url, nil, nil, &software); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &software, nil
}

// GetLicensedSoftwareByName retrieves the Licensed Software by its Name
func (c *Client) GetLicensedSoftwareByName(name string) (*ResponseLicensedSoftware, error) {
	url := fmt.Sprintf("%s/name/%s", uriLicensedSoftware, name)

	var software ResponseLicensedSoftware
	if err := c.DoRequest("GET", url, nil, nil, &software); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &software, nil
}

// GetLicensedSoftwares retrieves all Licensed Softwares
func (c *Client) GetLicensedSoftwares() (*ResponseLicensedSoftwareList, error) {
	url := uriLicensedSoftware

	var softwareList ResponseLicensedSoftwareList
	if err := c.DoRequest("GET", url, nil, nil, &softwareList); err != nil {
		return nil, fmt.Errorf("failed to fetch all Licensed Software: %v", err)
	}

	return &softwareList, nil
}

// CreateLicensedSoftware creates a new Licensed Software
func (c *Client) CreateLicensedSoftware(software *ResponseLicensedSoftware) (*ResponseLicensedSoftware, error) {
	url := fmt.Sprintf("%s/id/0", uriLicensedSoftware)

	// If no value is provided for jamf site id, set it to -1
	if software.General.Site.ID == 0 {
		software.General.Site.ID = -1
		software.General.Site.Name = "None"
	}

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"licensed_software"`
		*ResponseLicensedSoftware
	}{
		ResponseLicensedSoftware: software,
	}

	// Execute the request
	var responseSoftware ResponseLicensedSoftware
	if err := c.DoRequest("POST", url, reqBody, nil, &responseSoftware); err != nil {
		return nil, fmt.Errorf("failed to create Licensed Software: %v", err)
	}

	return &responseSoftware, nil
}

// UpdateLicensedSoftwareById updates an existing Licensed Software by its ID
func (c *Client) UpdateLicensedSoftwareById(id int, software *ResponseLicensedSoftware) (*ResponseLicensedSoftware, error) {
	url := fmt.Sprintf("%s/id/%d", uriLicensedSoftware, id)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"licensed_software"`
		*ResponseLicensedSoftware
	}{
		ResponseLicensedSoftware: software,
	}

	// Execute the request
	var responseSoftware ResponseLicensedSoftware
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseSoftware); err != nil {
		return nil, fmt.Errorf("failed to update Licensed Software by ID: %v", err)
	}

	return &responseSoftware, nil
}

// UpdateLicensedSoftwareByName updates an existing Licensed Software by its name
func (c *Client) UpdateLicensedSoftwareByName(name string, software *ResponseLicensedSoftware) (*ResponseLicensedSoftware, error) {
	url := fmt.Sprintf("%s/name/%s", uriLicensedSoftware, name)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"licensed_software"`
		*ResponseLicensedSoftware
	}{
		ResponseLicensedSoftware: software,
	}

	// Execute the request
	var responseSoftware ResponseLicensedSoftware
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseSoftware); err != nil {
		return nil, fmt.Errorf("failed to update Licensed Software by name: %v", err)
	}

	return &responseSoftware, nil
}

// DeleteLicensedSoftwareById deletes an existing Licensed Software by its ID
func (c *Client) DeleteLicensedSoftwareById(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriLicensedSoftware, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Licensed Software by ID: %v", err)
	}

	return nil
}

// DeleteLicensedSoftwareByName deletes an existing Licensed Software by its name
func (c *Client) DeleteLicensedSoftwareByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriLicensedSoftware, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Licensed Software by name: %v", err)
	}

	return nil
}
