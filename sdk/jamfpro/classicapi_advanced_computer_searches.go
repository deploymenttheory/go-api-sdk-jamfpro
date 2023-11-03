// classicapi_advanced_computer_searches.go
// Jamf Pro Classic Api - Advanced Computer Searches
// api reference: https://developer.jamf.com/jamf-pro/reference/advancedcomputersearches
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAdvancedComputerSearches = "/JSSResource/advancedcomputersearches"

// ResponseAdvancedComputerSearches represents the structure for multiple advanced computer searches.
type ResponseAdvancedComputerSearches struct {
	Size                     int                            `xml:"size"`
	AdvancedComputerSearches []AdvancedComputerSearchDetail `xml:"advanced_computer_search"`
}

// AdvancedComputerSearchDetail represents the details of an advanced computer search.
type AdvancedComputerSearchDetail struct {
	ID   int    `json:"id,omitempty" xml:"id"`
	Name string `json:"name,omitempty" xml:"name"`
}

// ResponseAdvancedComputerSearch represents the structure of the response for an advanced computer search.
type ResponseAdvancedComputerSearch struct {
	ID            int                                        `xml:"id"`
	Name          string                                     `xml:"name"`
	ViewAs        string                                     `xml:"view_as,omitempty"`
	Sort1         string                                     `xml:"sort_1,omitempty"`
	Sort2         string                                     `xml:"sort_2,omitempty"`
	Sort3         string                                     `xml:"sort_3,omitempty"`
	Criteria      []AdvancedComputerSearchesCriteria         `xml:"criteria"`
	DisplayFields []AdvancedComputerSearchesDisplayField     `xml:"display_fields"`
	Computers     []AdvancedComputerSearchDataSubsetComputer `xml:"computers"`
	Site          AdvancedComputerSearchesSiteDetail         `xml:"site"`
}

// Criteria represents a criterion with its details.
type AdvancedComputerSearchesCriteria struct {
	Size      int             `xml:"size"`
	Criterion CriterionDetail `xml:"criterion"`
}

// CriterionDetail represents the details of a criterion in a search.
type CriterionDetail struct {
	Name         string `xml:"name"`
	Priority     int    `xml:"priority,omitempty"`
	AndOr        string `xml:"and_or,omitempty"`
	SearchType   string `xml:"search_type,omitempty"`
	Value        string `xml:"value,omitempty"`
	OpeningParen bool   `xml:"opening_paren,omitempty"`
	ClosingParen bool   `xml:"closing_paren,omitempty"`
}

// DisplayField represents a display field with its details.
type AdvancedComputerSearchesDisplayField struct {
	Size         int                `xml:"size"`
	DisplayField DisplayFieldDetail `xml:"display_field"`
}

// DisplayFieldDetail represents the details of a display field.
type DisplayFieldDetail struct {
	Name string `xml:"name"`
}

// Computer represents a computer with its details.
type AdvancedComputerSearchDataSubsetComputer struct {
	Size     int            `xml:"size"`
	Computer ComputerDetail `xml:"computer"`
}

// ComputerDetail represents the details of a computer.
type ComputerDetail struct {
	ID           int    `xml:"id"`
	Name         string `xml:"name"`
	UDID         string `xml:"udid,omitempty"`
	ComputerName string `xml:"Computer_Name,omitempty"`
}

// SiteDetail represents the details of a site.
type AdvancedComputerSearchesSiteDetail struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// GetAdvancedComputerSearches retrieves all advanced computer searches.
func (c *Client) GetAdvancedComputerSearches() (*ResponseAdvancedComputerSearches, error) {
	endpoint := uriAPIAdvancedComputerSearches

	var searchesList ResponseAdvancedComputerSearches
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch advanced computer searches: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchesList, nil
}

// GetAdvancedComputerSearchByID retrieves an advanced computer search by its ID
func (c *Client) GetAdvancedComputerSearchByID(id int) (*ResponseAdvancedComputerSearch, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedComputerSearches, id)

	var search ResponseAdvancedComputerSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &search)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch advanced computer search by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &search, nil
}

// GetAdvancedComputerSearchesByName retrieves advanced computer searches by their name
func (c *Client) GetAdvancedComputerSearchesByName(name string) (*ResponseAdvancedComputerSearch, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedComputerSearches, name)

	var search ResponseAdvancedComputerSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &search)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch advanced computer search by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &search, nil
}

// CreateAdvancedComputerSearch creates a new advanced computer search.
func (c *Client) CreateAdvancedComputerSearch(search *ResponseAdvancedComputerSearch) (*ResponseAdvancedComputerSearch, error) {
	endpoint := uriAPIAdvancedComputerSearches

	// Set default values for Site if not provided
	if search.Site.ID == 0 && search.Site.Name == "" {
		search.Site.ID = -1
		search.Site.Name = "None"
	}

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_computer_search"`
		*ResponseAdvancedComputerSearch
	}{
		ResponseAdvancedComputerSearch: search,
	}

	var createdSearch ResponseAdvancedComputerSearch
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to create advanced computer search: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdSearch, nil
}

// UpdateAdvancedComputerSearchByID updates an existing advanced computer search by its ID.
func (c *Client) UpdateAdvancedComputerSearchByID(id int, search *ResponseAdvancedComputerSearch) (*ResponseAdvancedComputerSearch, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedComputerSearches, id)

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_computer_search"`
		*ResponseAdvancedComputerSearch
	}{
		ResponseAdvancedComputerSearch: search,
	}

	var updatedSearch ResponseAdvancedComputerSearch
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to update advanced computer search by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

// UpdateAdvancedComputerSearchByName updates an existing advanced computer search by its name.
func (c *Client) UpdateAdvancedComputerSearchByName(name string, search *ResponseAdvancedComputerSearch) (*ResponseAdvancedComputerSearch, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedComputerSearches, name)

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_computer_search"`
		*ResponseAdvancedComputerSearch
	}{
		ResponseAdvancedComputerSearch: search,
	}

	var updatedSearch ResponseAdvancedComputerSearch
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to update advanced computer search by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

// DeleteAdvancedComputerSearchByID deletes an advanced computer search by its ID.
func (c *Client) DeleteAdvancedComputerSearchByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedComputerSearches, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete advanced computer search by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteAdvancedComputerSearchByName deletes an advanced computer search by its name.
func (c *Client) DeleteAdvancedComputerSearchByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedComputerSearches, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete advanced computer search by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
