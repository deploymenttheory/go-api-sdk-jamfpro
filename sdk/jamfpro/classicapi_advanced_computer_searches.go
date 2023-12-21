// classicapi_advanced_computer_searches.go
// Jamf Pro Classic Api - Advanced Computer Searches
// api reference: https://developer.jamf.com/jamf-pro/reference/advancedcomputersearches
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint
SharedResourceSite
SharedAdvancedSearchContainerCriteria
SharedAdvancedSearchSubsetDisplayField
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAdvancedComputerSearches = "/JSSResource/advancedcomputersearches"

/// List

// ResponseAdvancedComputerSearchesList represents the structure for multiple advanced computer searches.
type ResponseAdvancedComputerSearchesList struct {
	Size                     int                              `xml:"size"`
	AdvancedComputerSearches []AdvancedComputerSearchListItem `xml:"advanced_computer_search"`
}

type AdvancedComputerSearchListItem struct {
	ID   int    `json:"id,omitempty" xml:"id"`
	Name string `json:"name,omitempty" xml:"name"`
}

/// Resource

// ResourceAdvancedComputerSearch represents the structure of the response for an advanced computer search.
type ResourceAdvancedComputerSearch struct {
	ID            int                                        `xml:"id"`
	Name          string                                     `xml:"name"`
	ViewAs        string                                     `xml:"view_as,omitempty"`
	Sort1         string                                     `xml:"sort_1,omitempty"`
	Sort2         string                                     `xml:"sort_2,omitempty"`
	Sort3         string                                     `xml:"sort_3,omitempty"`
	Criteria      []SharedAdvancedSearchContainerCriteria    `xml:"criteria"`
	DisplayFields []SharedAdvancedSearchSubsetDisplayField   `xml:"display_fields"`
	Computers     []AdvancedComputerSearchContainerComputers `xml:"computer"`
	Site          SharedResourceSite                         `xml:"site"`
}

/// Subsets & Containers
// Computer

type AdvancedComputerSearchContainerComputers struct {
	Size     int
	Computer AdvancedComputerSearchSubsetComputer `xml:"computers"`
}

type AdvancedComputerSearchSubsetComputer struct {
	ID           int    `xml:"id"`
	Name         string `xml:"name"`
	UDID         string `xml:"udid,omitempty"`
	ComputerName string `xml:"Computer_Name,omitempty"`
}

// GetAdvancedComputerSearches retrieves all advanced computer searches.
func (c *Client) GetAdvancedComputerSearches() (*ResponseAdvancedComputerSearchesList, error) {
	endpoint := uriAPIAdvancedComputerSearches

	var searchesList ResponseAdvancedComputerSearchesList
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
func (c *Client) GetAdvancedComputerSearchByID(id int) (*ResourceAdvancedComputerSearch, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedComputerSearches, id)

	var search ResourceAdvancedComputerSearch
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
func (c *Client) GetAdvancedComputerSearchesByName(name string) (*ResourceAdvancedComputerSearch, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedComputerSearches, name)

	var search ResourceAdvancedComputerSearch
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
func (c *Client) CreateAdvancedComputerSearch(search *ResourceAdvancedComputerSearch) (*ResourceAdvancedComputerSearch, error) {
	endpoint := uriAPIAdvancedComputerSearches

	// Set default values for Site if not provided
	if search.Site.ID == 0 && search.Site.Name == "" {
		search.Site.ID = -1
		search.Site.Name = "None"
	}

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_computer_search"`
		*ResourceAdvancedComputerSearch
	}{
		ResourceAdvancedComputerSearch: search,
	}

	var createdSearch ResourceAdvancedComputerSearch
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
func (c *Client) UpdateAdvancedComputerSearchByID(id int, search *ResourceAdvancedComputerSearch) (*ResourceAdvancedComputerSearch, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedComputerSearches, id)

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_computer_search"`
		*ResourceAdvancedComputerSearch
	}{
		ResourceAdvancedComputerSearch: search,
	}

	var updatedSearch ResourceAdvancedComputerSearch
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
func (c *Client) UpdateAdvancedComputerSearchByName(name string, search *ResourceAdvancedComputerSearch) (*ResourceAdvancedComputerSearch, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedComputerSearches, name)

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_computer_search"`
		*ResourceAdvancedComputerSearch
	}{
		ResourceAdvancedComputerSearch: search,
	}

	var updatedSearch ResourceAdvancedComputerSearch
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
