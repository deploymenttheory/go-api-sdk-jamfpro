// classicapi_advanced_user_searches.go
// Jamf Pro Classic Api - Advanced User Searches
// api reference: https://developer.jamf.com/jamf-pro/reference/advancedusersearches
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAdvancedUserSearches = "/JSSResource/advancedusersearches"

// Response structure for the list of advanced user searches
type ResponseAdvancedUserSearchesList struct {
	Size               int                        `xml:"size"`
	AdvancedUserSearch []AdvancedUserSearchDetail `xml:"advanced_user_search"`
}

type AdvancedUserSearchDetail struct {
	XMLName xml.Name `xml:"advanced_user_search"`
	ID      int      `xml:"id"`
	Name    string   `xml:"name"`
}

// Structs for Advanced User Search details by ID
type AdvancedUserSearch struct {
	ID            int                                        `xml:"id"`
	Name          string                                     `xml:"name"`
	Criteria      []AdvancedUserSearchCriteriaDetail         `xml:"criteria"`
	Users         []AdvancedUserSearchSiteUsersDetail        `xml:"users"`
	DisplayFields []AdvancedUserSearchSiteDisplayFieldDetail `xml:"display_fields"`
	Site          AdvancedUserSearchSiteDetail               `xml:"site"`
}

type AdvancedUserSearchCriteriaDetail struct {
	Size      int                               `xml:"size"`
	Criterion AdvancedUserSearchCriterionDetail `xml:"criterion"`
}

type AdvancedUserSearchCriterionDetail struct {
	Name         string `xml:"name"`
	Priority     int    `xml:"priority"`
	AndOr        string `xml:"and_or"`
	SearchType   string `xml:"search_type"`
	Value        string `xml:"value"`
	OpeningParen bool   `xml:"opening_paren"`
	ClosingParen bool   `xml:"closing_paren"`
}

type AdvancedUserSearchSiteUsersDetail struct {
	Size int                              `xml:"size"`
	User AdvancedUserSearchSiteUserDetail `xml:"user"`
}

type AdvancedUserSearchSiteUserDetail struct {
	ID       int    `xml:"id,omitempty"`
	Name     string `xml:"name,omitempty"`
	Username string `xml:"Username,omitempty"`
}

type AdvancedUserSearchSiteDisplayFieldDetail struct {
	Size         int `xml:"size"`
	DisplayField struct {
		Name string `xml:"name"`
	} `xml:"display_field"`
}

type AdvancedUserSearchSiteDetail struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// GetAdvancedUserSearches retrieves all advanced user searches
func (c *Client) GetAdvancedUserSearches() (*ResponseAdvancedUserSearchesList, error) {
	endpoint := uriAPIAdvancedUserSearches

	var advancedUserSearchesList ResponseAdvancedUserSearchesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &advancedUserSearchesList)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch advanced user searches: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &advancedUserSearchesList, nil
}

// GetAdvancedUserSearchByID retrieves an advanced user search by its ID
func (c *Client) GetAdvancedUserSearchByID(id int) (*AdvancedUserSearch, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedUserSearches, id)

	var searchDetail AdvancedUserSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch advanced user search by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchDetail, nil
}

// GetAdvancedUserSearchByName retrieves an advanced user search by its name
func (c *Client) GetAdvancedUserSearchByName(name string) (*AdvancedUserSearch, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedUserSearches, name)

	var searchDetail AdvancedUserSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchDetail)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch advanced user search by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchDetail, nil
}

// CreateAdvancedUserSearch creates a new advanced user search.
func (c *Client) CreateAdvancedUserSearch(search *AdvancedUserSearch) (*AdvancedUserSearch, error) {
	endpoint := uriAPIAdvancedUserSearches

	// Set default values for Site if not provided
	if search.Site.ID == 0 && search.Site.Name == "" {
		search.Site.ID = -1
		search.Site.Name = "None"
	}

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_user_search"`
		*AdvancedUserSearch
	}{
		AdvancedUserSearch: search,
	}

	var createdSearch AdvancedUserSearch
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to create advanced user search: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdSearch, nil
}

// UpdateAdvancedUserSearchByID updates an existing advanced user search by its ID.
func (c *Client) UpdateAdvancedUserSearchByID(id int, search *AdvancedUserSearch) (*AdvancedUserSearch, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedUserSearches, id)

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_user_search"`
		*AdvancedUserSearch
	}{
		AdvancedUserSearch: search,
	}

	var updatedSearch AdvancedUserSearch
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to update advanced user search by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

// UpdateAdvancedUserSearchByName updates an existing advanced user search by its name.
func (c *Client) UpdateAdvancedUserSearchByName(name string, search *AdvancedUserSearch) (*AdvancedUserSearch, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedUserSearches, name)

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_user_search"`
		*AdvancedUserSearch
	}{
		AdvancedUserSearch: search,
	}

	var updatedSearch AdvancedUserSearch
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf("failed to update advanced user search by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

// DeleteAdvancedUserSearchByID deletes an advanced user search by its ID.
func (c *Client) DeleteAdvancedUserSearchByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedUserSearches, id)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete advanced user search by ID: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	return nil
}

// DeleteAdvancedUserSearchByName deletes an advanced user search by its name.
func (c *Client) DeleteAdvancedUserSearchByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedUserSearches, name)
	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete advanced user search by name: %v", err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	return nil
}
