/* Filename: classicapi_advanced_mobile_searches.go
- Jamf Pro Classic API
- Resource: Advanced Mobile Searches
- API reference: https://developer.jamf.com/jamf-pro/reference/advancedmobiledevicesearches
- Data Structure: XML

Shared data structure resources in this endpoint:
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
	ID            int                                         `xml:"id"`
	Name          string                                      `xml:"name"`
	ViewAs        string                                      `xml:"view_as,omitempty"`
	Sort1         string                                      `xml:"sort_1,omitempty"`
	Sort2         string                                      `xml:"sort_2,omitempty"`
	Sort3         string                                      `xml:"sort_3,omitempty"`
	Criteria      SharedContainerCriteria                     `xml:"criteria,omitempty"`
	DisplayFields []SharedAdvancedSearchContainerDisplayField `xml:"display_fields,omitempty"`
	MobileDevices []AdvancedMobileSearchContainerDevices      `xml:"mobile_devices,omitempty"`
	Site          SharedResourceSite                          `xml:"site,omitempty"`
}

// Responses

type ResponseAdvancedMobileDeviceSearchCreatedAndUpdated struct {
	ID int `json:"id,omitempty" xml:"id,omitempty"`
}

// Subsets & Containers
// Mobile Device

type AdvancedMobileSearchContainerDevices struct {
	Size   int                              `xml:"size,omitempty"`
	Device AdvancedMobileSearchSubsetDevice `xml:"mobile_device,omitempty"`
}

type AdvancedMobileSearchSubsetDevice struct {
	ID          int    `xml:"id,omitempty"`
	Name        string `xml:"name,omitempty"`
	UDID        string `xml:"udid,omitempty"`
	DisplayName string `xml:"Display_Name,omitempty"`
}

// CRUD

/*
Function: GetAdvancedMobileDeviceSearches
Method: GET
Path: /JSSResource/advancedmobiledevicesearches
Description: Gets a list of all Jamf Pro Advanced Mobile Device Search resources.
*/
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

/*
Function: GetAdvancedMobileDeviceSearchByID
Method: GET
Path: /JSSResource/advancedmobiledevicesearches/id/{id}
Description: Gets a Jamf Pro advanced mobile device search resource by its ID.
*/
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

/*
Function: GetAdvancedMobileDeviceSearchByName
Method: GET
Path: /JSSResource/advancedmobiledevicesearches/name/{name}
Description: Gets a Jamf Pro advanced mobile device search resource by its name.
*/
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

/*
Function: CreateAdvancedMobileDeviceSearch
Method: POST
Path: /JSSResource/advancedmobiledevicesearches
Description: Creates a new Jamf Pro advanced mobile device search resource.
*/
func (c *Client) CreateAdvancedMobileDeviceSearch(search *ResourceAdvancedMobileDeviceSearch) (*ResponseAdvancedMobileDeviceSearchCreatedAndUpdated, error) {
	endpoint := uriAPIAdvancedMobileDeviceSearches

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

/*
Function: UpdateAdvancedMobileDeviceSearchByID
Method: PUT
Path: /JSSResource/advancedmobiledevicesearches/id/{id}
Description: Updates an existing Jamf Pro advanced mobile device search resource by its ID.
*/
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

/*
Function: UpdateAdvancedMobileDeviceSearchByName
Method: PUT
Path: /JSSResource/advancedmobiledevicesearches/name/{name}
Description: Updates an existing Jamf Pro advanced mobile device search resource by its name.
*/
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

/*
Function: DeleteAdvancedMobileDeviceSearchByID
Method: DELETE
Path: /JSSResource/advancedmobiledevicesearches/id/{id}
Description: Deletes an existing Jamf Pro advanced mobile device search resource by its ID.
*/
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

/*
Function: DeleteAdvancedMobileDeviceSearchByName
Method: DELETE
Path: /JSSResource/advancedmobiledevicesearches/name/{name}
Description: Deletes an existing Jamf Pro advanced mobile device search resource by its name.
*/
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
