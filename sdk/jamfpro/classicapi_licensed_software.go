// classicapi_licensed_software.go
// Jamf Pro Classic Api - Licensed Software
// api reference: https://developer.jamf.com/jamf-pro/reference/licensedsoftware
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriLicensedSoftware = "/JSSResource/licensedsoftware"

// ResponseLicensedSoftwareList represents the response for a list of licensed software.
type ResponseLicensedSoftwareList struct {
	LicensedSoftware []LicensedSoftwareItem `xml:"licensed_software"`
}

// LicensedSoftwareItem represents a single licensed software item.
type LicensedSoftwareItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResponseLicensedSoftware represents the structure of a single licensed software item.
type ResponseLicensedSoftware struct {
	General             LicensedSoftwareGeneral            `xml:"general"`
	SoftwareDefinitions []SoftwareDefinition               `xml:"software_definitions>definition"`
	FontDefinitions     []LicensedSoftwareFontDefinition   `xml:"font_definitions>definition"`
	PluginDefinitions   []LicensedSoftwarePluginDefinition `xml:"plugin_definitions>definition"`
	Licenses            []LicensedSoftwareLicense          `xml:"licenses>license"`
}

type LicensedSoftwareGeneral struct {
	ID                                 int                  `xml:"id"`
	Name                               string               `xml:"name"`
	Publisher                          string               `xml:"publisher"`
	Platform                           string               `xml:"platform"`
	SendEmailOnViolation               bool                 `xml:"send_email_on_violation"`
	RemoveTitlesFromInventoryReports   bool                 `xml:"remove_titles_from_inventory_reports"`
	ExcludeTitlesPurchasedFromAppStore bool                 `xml:"exclude_titles_purchased_from_app_store"`
	Notes                              string               `xml:"notes"`
	Site                               LicensedSoftwareSite `xml:"site"`
}

type SoftwareDefinition struct {
	CompareType string `xml:"compare_type"`
	Name        string `xml:"name"`
	Version     int    `xml:"version"`
}

type LicensedSoftwareFontDefinition struct {
	CompareType string `xml:"compare_type"`
	Name        string `xml:"name"`
	Version     int    `xml:"version"`
}

type LicensedSoftwarePluginDefinition struct {
	CompareType string `xml:"compare_type"`
	Name        string `xml:"name"`
	Version     int    `xml:"version"`
}

type LicensedSoftwareLicense struct {
	Size    int           `xml:"size"`
	License LicenseDetail `xml:"license"`
}

type LicenseDetail struct {
	SerialNumber1    string                       `xml:"serial_number_1"`
	SerialNumber2    string                       `xml:"serial_number_2"`
	OrganizationName string                       `xml:"organization_name"`
	RegisteredTo     string                       `xml:"registered_to"`
	LicenseType      string                       `xml:"license_type"`
	LicenseCount     int                          `xml:"license_count"`
	Notes            string                       `xml:"notes"`
	Purchasing       PurchasingDetail             `xml:"purchasing"`
	Attachments      []LicensedSoftwareAttachment `xml:"attachments>attachment"`
}

type PurchasingDetail struct {
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
	// Other relevant fields
}

type LicensedSoftwareAttachment struct {
	ID       int    `xml:"id"`
	Filename string `xml:"filename"`
	URI      string `xml:"uri"`
}

type LicensedSoftwareSite struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

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
func (c *Client) GetLicensedSoftwareByID(id int) (*ResponseLicensedSoftware, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriLicensedSoftware, id)

	var licensedSoftware ResponseLicensedSoftware
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
func (c *Client) GetLicensedSoftwareByName(name string) (*ResponseLicensedSoftware, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriLicensedSoftware, name)

	var licensedSoftware ResponseLicensedSoftware
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
func (c *Client) CreateLicensedSoftware(licensedSoftware *ResponseLicensedSoftware) (*ResponseLicensedSoftware, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriLicensedSoftware) // '0' typically used for creation in APIs

	// Set default values for site if not included within request
	if licensedSoftware.General.Site.ID == 0 && licensedSoftware.General.Site.Name == "" {
		licensedSoftware.General.Site = LicensedSoftwareSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Wrap licensedSoftware in an anonymous struct to match the expected XML structure
	requestBody := struct {
		XMLName xml.Name `xml:"licensed_software"`
		*ResponseLicensedSoftware
	}{
		ResponseLicensedSoftware: licensedSoftware,
	}

	var responseLicensedSoftware ResponseLicensedSoftware
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseLicensedSoftware)
	if err != nil {
		return nil, fmt.Errorf("failed to create licensed software: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseLicensedSoftware, nil
}

// UpdateLicensedSoftwareByID updates an existing licensed software item by its ID.
func (c *Client) UpdateLicensedSoftwareByID(id int, licensedSoftware *ResponseLicensedSoftware) (*ResponseLicensedSoftware, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriLicensedSoftware, id)

	requestBody := struct {
		XMLName xml.Name `xml:"licensed_software"`
		*ResponseLicensedSoftware
	}{
		ResponseLicensedSoftware: licensedSoftware,
	}

	var responseLicensedSoftware ResponseLicensedSoftware
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseLicensedSoftware)
	if err != nil {
		return nil, fmt.Errorf("failed to update licensed software by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseLicensedSoftware, nil
}

// UpdateLicensedSoftwareByName updates an existing licensed software item by its name.
func (c *Client) UpdateLicensedSoftwareByName(name string, licensedSoftware *ResponseLicensedSoftware) (*ResponseLicensedSoftware, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriLicensedSoftware, name)

	requestBody := struct {
		XMLName xml.Name `xml:"licensed_software"`
		*ResponseLicensedSoftware
	}{
		ResponseLicensedSoftware: licensedSoftware,
	}

	var responseLicensedSoftware ResponseLicensedSoftware
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseLicensedSoftware)
	if err != nil {
		return nil, fmt.Errorf("failed to update licensed software by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseLicensedSoftware, nil
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
