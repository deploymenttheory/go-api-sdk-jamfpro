// classicapi_advanced_mobile_searches.go
// Jamf Pro Classic Api - Advanced Mobile Searches
// api reference: https://developer.jamf.com/jamf-pro/reference/advancedmobiledevicesearches
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAdvancedMobileDeviceSearches = "/JSSResource/advancedmobiledevicesearches"

// ResponseAdvancedMobileDeviceSearchesList represents the structure for multiple advanced mobile device searches.
type ResponseAdvancedMobileDeviceSearchesList struct {
	Size                         int                                `xml:"size"`
	AdvancedMobileDeviceSearches []AdvancedMobileDeviceSearchDetail `xml:"advanced_mobile_device_search"`
}

// AdvancedMobileDeviceSearchDetail represents the details of an advanced mobile device search.
type AdvancedMobileDeviceSearchDetail struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// ResponseAdvancedMobileDeviceSearches represents the structure of the response for an advanced mobile device search.
type ResponseAdvancedMobileDeviceSearches struct {
	ID            int                                        `xml:"id"`                       // Unique identifier for the search
	Name          string                                     `xml:"name"`                     // Name of the search
	ViewAs        string                                     `xml:"view_as,omitempty"`        // The format in which the search results are viewed
	Sort1         string                                     `xml:"sort_1,omitempty"`         // First sorting criteria
	Sort2         string                                     `xml:"sort_2,omitempty"`         // Second sorting criteria
	Sort3         string                                     `xml:"sort_3,omitempty"`         // Third sorting criteria
	Criteria      []AdvancedMobileDeviceSearchesCriteria     `xml:"criteria,omitempty"`       // List of search criteria
	DisplayFields []AdvancedMobileDeviceSearchesDisplayField `xml:"display_fields,omitempty"` // Fields to display in search results
	MobileDevices []AdvancedMobileDeviceSearchesMobileDevice `xml:"mobile_devices,omitempty"` // List of mobile devices that match the search
	Site          AdvancedMobileDeviceSearchesSite           `xml:"site,omitempty"`           // Information about the site associated with the search
}

// CriteriaDetail represents a single search criterion.
type AdvancedMobileDeviceSearchesCriteria struct {
	Size      int       `xml:"size"`      // Number of criteria
	Criterion Criterion `xml:"criterion"` // Detailed criterion
}

// Criterion contains the details of a single criterion in the search.
type Criterion struct {
	Name         string `xml:"name"`                    // Name of the criterion
	Priority     int    `xml:"priority"`                // Priority of the criterion
	AndOr        string `xml:"and_or"`                  // Logical operator to combine criteria
	SearchType   string `xml:"search_type"`             // Type of search being performed
	Value        int    `xml:"value"`                   // Value for the criterion
	OpeningParen bool   `xml:"opening_paren,omitempty"` // Indicates if there is an opening parenthesis for grouping
	ClosingParen bool   `xml:"closing_paren,omitempty"` // Indicates if there is a closing parenthesis for grouping
}

// DisplayFieldDetailWrapper wraps a display field with its size.
type AdvancedMobileDeviceSearchesDisplayField struct {
	Size         int                                          `xml:"size"`          // Number of display fields
	DisplayField AdvancedMobileDeviceSearchesDisplayFieldItem `xml:"display_field"` // Detailed display field
}

// DisplayField represents a field to display in the search results.
type AdvancedMobileDeviceSearchesDisplayFieldItem struct {
	Name string `xml:"name"` // Name of the display field
}

// MobileDeviceDetailWrapper wraps a mobile device with its size.
type AdvancedMobileDeviceSearchesMobileDevice struct {
	Size         int              `xml:"size"`          // Number of mobile devices
	MobileDevice MobileDeviceItem `xml:"mobile_device"` // Detailed mobile device
}

// MobileDevice contains details about a single mobile device.
type MobileDeviceItem struct {
	ID          int    `xml:"id"`           // Unique identifier for the mobile device
	Name        string `xml:"name"`         // Name of the mobile device
	UDID        string `xml:"udid"`         // Unique Device Identifier for the mobile device
	DisplayName string `xml:"Display_Name"` // Display name of the mobile device
}

// SiteDetail represents the details of a site associated with the search.
type AdvancedMobileDeviceSearchesSite struct {
	ID   int    `xml:"id"`   // Unique identifier for the site
	Name string `xml:"name"` // Name of the site
}

// GetAdvancedMobileDeviceSearches retrieves all advanced mobile device searches.
func (c *Client) GetAdvancedMobileDeviceSearches() (*ResponseAdvancedMobileDeviceSearchesList, error) {
	endpoint := uriAPIAdvancedMobileDeviceSearches

	var searchesList ResponseAdvancedMobileDeviceSearchesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch advanced mobile device searches: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchesList, nil
}

// GetAdvancedMobileDeviceSearchByID retrieves an advanced mobile device search by its ID.
func (c *Client) GetAdvancedMobileDeviceSearchByID(id int) (*ResponseAdvancedMobileDeviceSearches, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedMobileDeviceSearches, id)

	var searchDetail ResponseAdvancedMobileDeviceSearches
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch advanced mobile device search by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchDetail, nil
}

// GetAdvancedMobileDeviceSearchByName retrieves an advanced mobile device search by its name.
func (c *Client) GetAdvancedMobileDeviceSearchByName(name string) (*ResponseAdvancedMobileDeviceSearches, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedMobileDeviceSearches, name)

	var searchDetail ResponseAdvancedMobileDeviceSearches
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch advanced mobile device search by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchDetail, nil
}

// CreateAdvancedMobileDeviceSearchByID creates a new advanced mobile device search with the given ID.
func (c *Client) CreateAdvancedMobileDeviceSearchByID(id int, search *ResponseAdvancedMobileDeviceSearches) (*ResponseAdvancedMobileDeviceSearches, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedMobileDeviceSearches, id)

	// Check if the Site field in the search struct is not provided and set default values if needed
	if search.Site.ID == 0 && search.Site.Name == "" {
		search.Site = AdvancedMobileDeviceSearchesSite{
			ID:   -1,
			Name: "None",
		}
	}

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_mobile_device_search"`
		*ResponseAdvancedMobileDeviceSearches
	}{
		ResponseAdvancedMobileDeviceSearches: search,
	}

	var createdSearch ResponseAdvancedMobileDeviceSearches
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to create advanced mobile device search by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdSearch, nil
}

// UpdateAdvancedMobileDeviceSearchByID updates an existing advanced mobile device search by its ID.
func (c *Client) UpdateAdvancedMobileDeviceSearchByID(id int, search *ResponseAdvancedMobileDeviceSearches) (*ResponseAdvancedMobileDeviceSearches, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedMobileDeviceSearches, id)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_mobile_device_search"`
		*ResponseAdvancedMobileDeviceSearches
	}{
		ResponseAdvancedMobileDeviceSearches: search,
	}

	var updatedSearch ResponseAdvancedMobileDeviceSearches
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to update advanced mobile device search by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

// UpdateAdvancedMobileDeviceSearchByName updates an existing advanced mobile device search by its name.
func (c *Client) UpdateAdvancedMobileDeviceSearchByName(name string, search *ResponseAdvancedMobileDeviceSearches) (*ResponseAdvancedMobileDeviceSearches, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedMobileDeviceSearches, name)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_mobile_device_search"`
		*ResponseAdvancedMobileDeviceSearches
	}{
		ResponseAdvancedMobileDeviceSearches: search,
	}

	var updatedSearch ResponseAdvancedMobileDeviceSearches
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to update advanced mobile device search by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

// DeleteAdvancedMobileDeviceSearchByID deletes an existing advanced mobile device search by its ID.
func (c *Client) DeleteAdvancedMobileDeviceSearchByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedMobileDeviceSearches, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete advanced mobile device search by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAdvancedMobileDeviceSearchByName deletes an existing advanced mobile device search by its name.
func (c *Client) DeleteAdvancedMobileDeviceSearchByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedMobileDeviceSearches, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete advanced mobile device search by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
