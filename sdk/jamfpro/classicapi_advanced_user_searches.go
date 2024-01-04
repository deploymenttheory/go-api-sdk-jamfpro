// classicapi_advanced_user_searches.go
// Jamf Pro Classic Api - Advanced User Searches
// api reference: https://developer.jamf.com/jamf-pro/reference/advancedusersearches
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint
- SharedResourceSite
- SharedContainerCriteria
- SharedAdvancedSearchSubsetDisplayField
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAdvancedUserSearches = "/JSSResource/advancedusersearches"

// List

// Response structure for the list of advanced user searches
type ResponseAdvancedUserSearchesList struct {
	Size               int                          `xml:"size"`
	AdvancedUserSearch []AdvancedUserSearchListItem `xml:"advanced_user_search"`
}

type AdvancedUserSearchListItem struct {
	XMLName xml.Name `xml:"advanced_user_search"`
	ID      int      `xml:"id"`
	Name    string   `xml:"name"`
}

// Resource

// Structs for Advanced User Search details by ID
type ResourceAdvancedUserSearch struct {
	ID            int                                       `xml:"id"`
	Name          string                                    `xml:"name"`
	Criteria      SharedContainerCriteria                   `xml:"criteria"`
	Users         []AdvancedUserSearchContainerUsers        `xml:"users"`
	DisplayFields SharedAdvancedSearchContainerDisplayField `xml:"display_fields"`
	Site          SharedResourceSite                        `xml:"site"`
}

// Responses

type ResourceAdvancedUserSearchCreatedAndUpdated struct {
	ID int `json:"id,omitempty" xml:"id,omitempty"`
}

// Subsets & Containers

type AdvancedUserSearchContainerUsers struct {
	Size int                          `xml:"size"`
	User AdvancedUserSearchSubsetUser `xml:"user"`
}

type AdvancedUserSearchSubsetUser struct {
	ID       int    `xml:"id,omitempty"`
	Name     string `xml:"name,omitempty"`
	Username string `xml:"Username,omitempty"`
}

// CRUD

// GetAdvancedUserSearches retrieves all advanced user searches
func (c *Client) GetAdvancedUserSearches() (*ResponseAdvancedUserSearchesList, error) {
	endpoint := uriAPIAdvancedUserSearches

	var advancedUserSearchesList ResponseAdvancedUserSearchesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &advancedUserSearchesList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "advanced user searches", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &advancedUserSearchesList, nil
}

// GetAdvancedUserSearchByID retrieves an advanced user search by its ID
func (c *Client) GetAdvancedUserSearchByID(id int) (*ResourceAdvancedUserSearch, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedUserSearches, id)

	var searchDetail ResourceAdvancedUserSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchDetail)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "advanced user search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchDetail, nil
}

// GetAdvancedUserSearchByName retrieves an advanced user search by its name
func (c *Client) GetAdvancedUserSearchByName(name string) (*ResourceAdvancedUserSearch, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedUserSearches, name)

	var searchDetail ResourceAdvancedUserSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchDetail)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "advanced user search", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchDetail, nil
}

// CreateAdvancedUserSearch creates a new advanced user search.
func (c *Client) CreateAdvancedUserSearch(search *ResourceAdvancedUserSearch) (*ResourceAdvancedUserSearchCreatedAndUpdated, error) {
	endpoint := uriAPIAdvancedUserSearches

	// Set default values for Site if not provided
	if search.Site.ID == 0 && search.Site.Name == "" {
		search.Site.ID = -1
		search.Site.Name = "None"
	}

	// Wrap the search request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"advanced_user_search"`
		*ResourceAdvancedUserSearch
	}{
		ResourceAdvancedUserSearch: search,
	}

	var createdSearch ResourceAdvancedUserSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "advanced user search", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdSearch, nil
}

// UpdateAdvancedUserSearchByID updates an existing advanced user search by its ID.
func (c *Client) UpdateAdvancedUserSearchByID(id int, search *ResourceAdvancedUserSearch) (*ResourceAdvancedUserSearchCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedUserSearches, id)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_user_search"`
		*ResourceAdvancedUserSearch
	}{
		ResourceAdvancedUserSearch: search,
	}

	var updatedSearch ResourceAdvancedUserSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "advanced user search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

// UpdateAdvancedUserSearchByName updates an existing advanced user search by its name.
func (c *Client) UpdateAdvancedUserSearchByName(name string, search *ResourceAdvancedUserSearch) (*ResourceAdvancedUserSearchCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedUserSearches, name)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_user_search"`
		*ResourceAdvancedUserSearch
	}{
		ResourceAdvancedUserSearch: search,
	}

	var updatedSearch ResourceAdvancedUserSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "advanced user search", name, err)
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
		return fmt.Errorf(errMsgFailedDeleteByID, "advanced user search", id, err)
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
		return fmt.Errorf(errMsgFailedDeleteByName, "advanced user search", name, err)
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	return nil
}
