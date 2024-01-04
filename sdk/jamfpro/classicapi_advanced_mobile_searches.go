// classicapi_advanced_mobile_searches.go
// Jamf Pro Classic Api - Advanced Mobile Searches
// api reference: https://developer.jamf.com/jamf-pro/reference/advancedmobiledevicesearches
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
- SharedContainerCriteria
- SharedAdvancedSearchSubsetDisplayField
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriAPIAdvancedMobileDeviceSearches = "/JSSResource/advancedmobiledevicesearches"

// List

// ResourceAdvancedMobileDeviceSearchList represents the structure for multiple advanced mobile device searches.
type ResponseAdvancedMobileDeviceSearchesList struct {
	Size                         int                            `xml:"size"`
	AdvancedMobileDeviceSearches []AdvancedMobileSearchListItem `xml:"advanced_mobile_device_search"`
}

type AdvancedMobileSearchListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// ResourceAdvancedMobileDeviceSearch represents the structure of the response for an advanced mobile device search.
type ResourceAdvancedMobileDeviceSearch struct {
	ID            int                                      `xml:"id"`
	Name          string                                   `xml:"name"`
	ViewAs        string                                   `xml:"view_as,omitempty"`
	Sort1         string                                   `xml:"sort_1,omitempty"`
	Sort2         string                                   `xml:"sort_2,omitempty"`
	Sort3         string                                   `xml:"sort_3,omitempty"`
	Criteria      SharedContainerCriteria                  `xml:"criteria,omitempty"`
	DisplayFields []SharedAdvancedSearchSubsetDisplayField `xml:"display_fields"`
	MobileDevices []AdvancedMobileSearchContainerDevices   `xml:"mobile_devices,omitempty"`
	Site          SharedResourceSite                       `xml:"site"`
}

// Responses

type ResponseAdvancedMobileDeviceSearchCreatedAndUpdated struct {
	ID int `json:"id,omitempty" xml:"id,omitempty"`
}

// Subsets & Containers
// Mobile Device

type AdvancedMobileSearchContainerDevices struct {
	Size   int                              `xml:"size"`
	Device AdvancedMobileSearchSubsetDevice `xml:"mobile_device"`
}

type AdvancedMobileSearchSubsetDevice struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	UDID        string `xml:"udid"`
	DisplayName string `xml:"Display_Name"`
}

// CRUD

// GetAdvancedMobileDeviceSearches retrieves all advanced mobile device searches.
func (c *Client) GetAdvancedMobileDeviceSearches() (*ResponseAdvancedMobileDeviceSearchesList, error) {
	endpoint := uriAPIAdvancedMobileDeviceSearches

	var searchesList ResponseAdvancedMobileDeviceSearchesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchesList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "advanced mobile device searches", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchesList, nil
}

// GetAdvancedMobileDeviceSearchByID retrieves an advanced mobile device search by its ID.
func (c *Client) GetAdvancedMobileDeviceSearchByID(id int) (*ResourceAdvancedMobileDeviceSearch, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedMobileDeviceSearches, id)

	var searchDetail ResourceAdvancedMobileDeviceSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchDetail)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "advanced mobile device search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchDetail, nil
}

// GetAdvancedMobileDeviceSearchByName retrieves an advanced mobile device search by its name.
func (c *Client) GetAdvancedMobileDeviceSearchByName(name string) (*ResourceAdvancedMobileDeviceSearch, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedMobileDeviceSearches, name)

	var searchDetail ResourceAdvancedMobileDeviceSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchDetail)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "advanced mobile device search", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchDetail, nil
}

// CreateAdvancedMobileDeviceSearch creates a new advanced mobile device search with the given ID.
func (c *Client) CreateAdvancedMobileDeviceSearch(search *ResourceAdvancedMobileDeviceSearch) (*ResponseAdvancedMobileDeviceSearchCreatedAndUpdated, error) {
	endpoint := uriAPIAdvancedMobileDeviceSearches

	if search.Site.ID == 0 && search.Site.Name == "" {
		search.Site.ID = -1
		search.Site.Name = "none"
	}

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_mobile_device_search"`
		*ResourceAdvancedMobileDeviceSearch
	}{
		ResourceAdvancedMobileDeviceSearch: search,
	}

	var createdSearch ResponseAdvancedMobileDeviceSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "advanced mobile device search", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdSearch, nil
}

// UpdateAdvancedMobileDeviceSearchByID updates an existing advanced mobile device search by its ID.
func (c *Client) UpdateAdvancedMobileDeviceSearchByID(id int, search *ResourceAdvancedMobileDeviceSearch) (*ResponseAdvancedMobileDeviceSearchCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedMobileDeviceSearches, id)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_mobile_device_search"`
		*ResourceAdvancedMobileDeviceSearch
	}{
		ResourceAdvancedMobileDeviceSearch: search,
	}

	var updatedSearch ResponseAdvancedMobileDeviceSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "advanced mobile device search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

// UpdateAdvancedMobileDeviceSearchByName updates an existing advanced mobile device search by its name.
func (c *Client) UpdateAdvancedMobileDeviceSearchByName(name string, search *ResourceAdvancedMobileDeviceSearch) (*ResponseAdvancedMobileDeviceSearchCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedMobileDeviceSearches, name)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_mobile_device_search"`
		*ResourceAdvancedMobileDeviceSearch
	}{
		ResourceAdvancedMobileDeviceSearch: search,
	}

	var updatedSearch ResponseAdvancedMobileDeviceSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "advanced mobile device search", name, err)
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
		return fmt.Errorf(errMsgFailedDeleteByID, "advanced mobile device search", id, err)
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
		return fmt.Errorf(errMsgFailedDeleteByName, "advanced mobile device search", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
