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

// ResourceAdvancedMobileDeviceSearchesList represents the structure for multiple advanced mobile device searches.
type ResponseAdvancedMobileDeviceSearchesList struct {
	Size                         int `xml:"size"`
	AdvancedMobileDeviceSearches []struct {
		ID   int    `xml:"id"`
		Name string `xml:"name"`
	} `xml:"advanced_mobile_device_search"`
}

// ResourceAdvancedMobileDeviceSearches represents the structure of the response for an advanced mobile device search.
type ResourceAdvancedMobileDeviceSearches struct {
	ID       int    `xml:"id"`                // Unique identifier for the search
	Name     string `xml:"name"`              // Name of the search
	ViewAs   string `xml:"view_as,omitempty"` // The format in which the search results are viewed
	Sort1    string `xml:"sort_1,omitempty"`  // First sorting criteria
	Sort2    string `xml:"sort_2,omitempty"`  // Second sorting criteria
	Sort3    string `xml:"sort_3,omitempty"`  // Third sorting criteria
	Criteria []struct {
		Size      int `xml:"size"` // Number of criteria
		Criterion struct {
			Name         string `xml:"name"`                    // Name of the criterion
			Priority     int    `xml:"priority"`                // Priority of the criterion
			AndOr        string `xml:"and_or"`                  // Logical operator to combine criteria
			SearchType   string `xml:"search_type"`             // Type of search being performed
			Value        int    `xml:"value"`                   // Value for the criterion
			OpeningParen bool   `xml:"opening_paren,omitempty"` // Indicates if there is an opening parenthesis for grouping
			ClosingParen bool   `xml:"closing_paren,omitempty"` // Indicates if there is a closing parenthesis for grouping
		} `xml:"criterion"` // Detailed criterion
	} `xml:"criteria,omitempty"` // List of search criteria
	DisplayFields []struct {
		Size         int `xml:"size"` // Number of display fields
		DisplayField struct {
			Name string `xml:"name"` // Name of the display field
		} `xml:"display_field"` // Detailed display field
	} `xml:"display_fields,omitempty"` // Fields to display in search results
	MobileDevices []struct {
		Size         int `xml:"size"` // Number of mobile devices
		MobileDevice struct {
			ID          int    `xml:"id"`           // Unique identifier for the mobile device
			Name        string `xml:"name"`         // Name of the mobile device
			UDID        string `xml:"udid"`         // Unique Device Identifier for the mobile device
			DisplayName string `xml:"Display_Name"` // Display name of the mobile device
		} `xml:"mobile_device"` // Detailed mobile device
	} `xml:"mobile_devices,omitempty"` // List of mobile devices that match the search
	Site struct {
		ID   int    `xml:"id"`   // Unique identifier for the site
		Name string `xml:"name"` // Name of the site
	} `xml:"site,omitempty"` // Information about the site associated with the search
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
func (c *Client) GetAdvancedMobileDeviceSearchByID(id int) (*ResourceAdvancedMobileDeviceSearches, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedMobileDeviceSearches, id)

	var searchDetail ResourceAdvancedMobileDeviceSearches
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
func (c *Client) GetAdvancedMobileDeviceSearchByName(name string) (*ResourceAdvancedMobileDeviceSearches, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedMobileDeviceSearches, name)

	var searchDetail ResourceAdvancedMobileDeviceSearches
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
func (c *Client) CreateAdvancedMobileDeviceSearchByID(id int, search *ResourceAdvancedMobileDeviceSearches) (*ResourceAdvancedMobileDeviceSearches, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedMobileDeviceSearches, id)

	// Check if the Site field in the search struct is not provided and set default values if needed
	if search.Site.ID == 0 && search.Site.Name == "" {
		search.Site.ID = -1
		search.Site.Name = "none"
	}

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_mobile_device_search"`
		*ResourceAdvancedMobileDeviceSearches
	}{
		ResourceAdvancedMobileDeviceSearches: search,
	}

	var createdSearch ResourceAdvancedMobileDeviceSearches
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
func (c *Client) UpdateAdvancedMobileDeviceSearchByID(id int, search *ResourceAdvancedMobileDeviceSearches) (*ResourceAdvancedMobileDeviceSearches, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedMobileDeviceSearches, id)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_mobile_device_search"`
		*ResourceAdvancedMobileDeviceSearches
	}{
		ResourceAdvancedMobileDeviceSearches: search,
	}

	var updatedSearch ResourceAdvancedMobileDeviceSearches
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
func (c *Client) UpdateAdvancedMobileDeviceSearchByName(name string, search *ResourceAdvancedMobileDeviceSearches) (*ResourceAdvancedMobileDeviceSearches, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedMobileDeviceSearches, name)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_mobile_device_search"`
		*ResourceAdvancedMobileDeviceSearches
	}{
		ResourceAdvancedMobileDeviceSearches: search,
	}

	var updatedSearch ResourceAdvancedMobileDeviceSearches
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
