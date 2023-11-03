// classicapi_advanced_mobile_searches.go
// Jamf Pro Classic Api - Advanced Mobile Searches
// api reference: https://developer.jamf.com/jamf-pro/reference/advancedmobiledevicesearches
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
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
	ID            int                                        `json:"id"`             // Unique identifier for the search
	Name          string                                     `json:"name"`           // Name of the search
	ViewAs        string                                     `json:"view_as"`        // The format in which the search results are viewed
	Sort1         string                                     `json:"sort_1"`         // First sorting criteria
	Sort2         string                                     `json:"sort_2"`         // Second sorting criteria
	Sort3         string                                     `json:"sort_3"`         // Third sorting criteria
	Criteria      []AdvancedMobileDeviceSearchesCriteria     `json:"criteria"`       // List of search criteria
	DisplayFields []AdvancedMobileDeviceSearchesDisplayField `json:"display_fields"` // Fields to display in search results
	MobileDevices []AdvancedMobileDeviceSearchesMobileDevice `json:"mobile_devices"` // List of mobile devices that match the search
	Site          AdvancedMobileDeviceSearchesSite           `json:"site"`           // Information about the site associated with the search
}

// CriteriaDetail represents a single search criterion.
type AdvancedMobileDeviceSearchesCriteria struct {
	Size      int       `json:"size"`      // Number of criteria
	Criterion Criterion `json:"criterion"` // Detailed criterion
}

// Criterion contains the details of a single criterion in the search.
type Criterion struct {
	Name         string `json:"name"`          // Name of the criterion
	Priority     int    `json:"priority"`      // Priority of the criterion
	AndOr        string `json:"and_or"`        // Logical operator to combine criteria
	SearchType   string `json:"search_type"`   // Type of search being performed
	Value        int    `json:"value"`         // Value for the criterion
	OpeningParen bool   `json:"opening_paren"` // Indicates if there is an opening parenthesis for grouping
	ClosingParen bool   `json:"closing_paren"` // Indicates if there is a closing parenthesis for grouping
}

// DisplayFieldDetailWrapper wraps a display field with its size.
type AdvancedMobileDeviceSearchesDisplayField struct {
	Size         int                                          `json:"size"`          // Number of display fields
	DisplayField AdvancedMobileDeviceSearchesDisplayFieldItem `json:"display_field"` // Detailed display field
}

// DisplayField represents a field to display in the search results.
type AdvancedMobileDeviceSearchesDisplayFieldItem struct {
	Name string `json:"name"` // Name of the display field
}

// MobileDeviceDetailWrapper wraps a mobile device with its size.
type AdvancedMobileDeviceSearchesMobileDevice struct {
	Size         int              `json:"size"`          // Number of mobile devices
	MobileDevice MobileDeviceItem `json:"mobile_device"` // Detailed mobile device
}

// MobileDevice contains details about a single mobile device.
type MobileDeviceItem struct {
	ID          int    `json:"id"`           // Unique identifier for the mobile device
	Name        string `json:"name"`         // Name of the mobile device
	UDID        string `json:"udid"`         // Unique Device Identifier for the mobile device
	DisplayName string `json:"Display_Name"` // Display name of the mobile device
}

// SiteDetail represents the details of a site associated with the search.
type AdvancedMobileDeviceSearchesSite struct {
	ID   int    `json:"id"`   // Unique identifier for the site
	Name string `json:"name"` // Name of the site
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
