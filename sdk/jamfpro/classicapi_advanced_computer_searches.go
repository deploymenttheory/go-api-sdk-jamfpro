/*
Filename: classicapi_advanced_computer_searches.go
- Jamf Pro Classic API
- Resource: Advanced Computer Searches
- API reference: https://developer.jamf.com/jamf-pro/reference/advancedcomputersearches
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

const uriAPIAdvancedComputerSearches = "/JSSResource/advancedcomputersearches"

// List

// ResponseAdvancedComputerSearchesList represents the structure for multiple advanced computer searches.
type ResponseAdvancedComputerSearchesList struct {
	Size                     int                              `xml:"size"`
	AdvancedComputerSearches []AdvancedComputerSearchListItem `xml:"advanced_computer_search"`
}

type AdvancedComputerSearchListItem struct {
	ID   int    `json:"id,omitempty" xml:"id"`
	Name string `json:"name,omitempty" xml:"name"`
}

// Resource

// ResourceAdvancedComputerSearch represents the structure of the response for an advanced computer search.
type ResourceAdvancedComputerSearch struct {
	ID            int                                        `xml:"id"`
	Name          string                                     `xml:"name"`
	ViewAs        string                                     `xml:"view_as,omitempty"`
	Sort1         string                                     `xml:"sort_1,omitempty"`
	Sort2         string                                     `xml:"sort_2,omitempty"`
	Sort3         string                                     `xml:"sort_3,omitempty"`
	Criteria      SharedContainerCriteria                    `xml:"criteria,omitempty"`
	DisplayFields []DisplayField                             `xml:"display_fields>display_field,omitempty"`
	Computers     []AdvancedComputerSearchContainerComputers `xml:"computer,omitempty"`
	Site          *SharedResourceSite                        `xml:"site,omitempty"`
}

// Responses

type ResponseAdvancedComputerSearchCreatedAndUpdated struct {
	ID int `json:"id,omitempty" xml:"id,omitempty"`
}

// Subsets & Containers

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

// CRUD

/*
Function: GetAdvancedComputerSearches
Method: GET
Path: /JSSResource/advancedcomputersearches
Description: Gets a list of all Jamf Pro Advanced Computer Search resources.
Parameters: None
Returns: ResponseAdvancedComputerSearchesList - A list of advanced computer searches.
Errors: Returns an error if the request fails.
Example:

	searches, err := client.GetAdvancedComputerSearches()
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(searches)
*/
func (c *Client) GetAdvancedComputerSearches() (*ResponseAdvancedComputerSearchesList, error) {
	endpoint := uriAPIAdvancedComputerSearches

	var searchesList ResponseAdvancedComputerSearchesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &searchesList)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "advance computer searches", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &searchesList, nil
}

/*
Function: GetAdvancedComputerSearchByID
Method: GET
Path: /JSSResource/advancedcomputersearches/id/{id}
Description: Gets a Jamf Pro advanced computer search resource by its ID.
Parameters:
  - id (int): The ID of the advanced computer search.

Returns: ResourceAdvancedComputerSearch - The advanced computer search resource.
Example:

	search, err := client.GetAdvancedComputerSearchByID(123)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(search)

Errors: Returns an error if the request fails or if the ID is not found.
*/
func (c *Client) GetAdvancedComputerSearchByID(id int) (*ResourceAdvancedComputerSearch, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedComputerSearches, id)

	var search ResourceAdvancedComputerSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &search)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "advance computer search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &search, nil
}

/*
Function: GetAdvancedComputerSearchByName
Method: GET
Path: /JSSResource/advancedcomputersearches/name/{name}
Description: Gets a Jamf Pro advanced computer search resource by its name.
Parameters:
  - name (string): The name of the advanced computer search.

Returns: ResourceAdvancedComputerSearch - The advanced computer search resource.
Example:

	search, err := client.GetAdvancedComputerSearchByName("SearchName")
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(search)

Errors: Returns an error if the request fails or if the name is not found.
*/
func (c *Client) GetAdvancedComputerSearchByName(name string) (*ResourceAdvancedComputerSearch, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedComputerSearches, name)

	var search ResourceAdvancedComputerSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &search)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "advance computer search", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &search, nil
}

/*
Function: CreateAdvancedComputerSearch
Method: POST
Path: /JSSResource/advancedcomputersearches
Description: Creates a new Jamf Pro advanced computer search resource.
Parameters:
  - search (*ResourceAdvancedComputerSearch): The advanced computer search resource to create.

Returns: ResponseAdvancedComputerSearchCreatedAndUpdated - The ID of the created advanced computer search resource.
Example:

	newSearch := &jamfpro.ResourceAdvancedComputerSearch{
	    Name: "New Search",
	    // Other fields...
	}
	created, err := client.CreateAdvancedComputerSearch(newSearch)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(created)

Errors: Returns an error if the request fails or if the resource cannot be created.
*/
func (c *Client) CreateAdvancedComputerSearch(search *ResourceAdvancedComputerSearch) (*ResponseAdvancedComputerSearchCreatedAndUpdated, error) {
	endpoint := uriAPIAdvancedComputerSearches

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_computer_search"`
		*ResourceAdvancedComputerSearch
	}{
		ResourceAdvancedComputerSearch: search,
	}

	var createdSearch ResponseAdvancedComputerSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "advance computer search", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdSearch, nil
}

/*
Function: UpdateAdvancedComputerSearchByID
Method: PUT
Path: /JSSResource/advancedcomputersearches/id/{id}
Description: Updates an existing Jamf Pro advanced computer search resource by its ID.
Parameters:
  - id (int): The ID of the advanced computer search.
  - search (*ResourceAdvancedComputerSearch): The updated advanced computer search resource.

Returns: ResponseAdvancedComputerSearchCreatedAndUpdated - The ID of the updated advanced computer search resource.
Example:

	updatedSearch := &jamfpro.ResourceAdvancedComputerSearch{
	    Name: "Updated Search",
	    // Other fields...
	}
	updated, err := client.UpdateAdvancedComputerSearchByID(123, updatedSearch)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(updated)

Errors: Returns an error if the request fails or if the resource cannot be updated.
*/
func (c *Client) UpdateAdvancedComputerSearchByID(id int, search *ResourceAdvancedComputerSearch) (*ResponseAdvancedComputerSearchCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedComputerSearches, id)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_computer_search"`
		*ResourceAdvancedComputerSearch
	}{
		ResourceAdvancedComputerSearch: search,
	}

	var updatedSearch ResponseAdvancedComputerSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "advance computer search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

/*
Function: UpdateAdvancedComputerSearchByName
Method: PUT
Path: /JSSResource/advancedcomputersearches/name/{name}
Description: Updates an existing Jamf Pro advanced computer search resource by its name.
Parameters:
  - name (string): The name of the advanced computer search.
  - search (*ResourceAdvancedComputerSearch): The updated advanced computer search resource.

Returns: ResponseAdvancedComputerSearchCreatedAndUpdated - The ID of the updated advanced computer search resource.
Example:

	updatedSearch := &jamfpro.ResourceAdvancedComputerSearch{
	    Name: "Updated Search",
	    // Other fields...
	}
	updated, err := client.UpdateAdvancedComputerSearchByName("SearchName", updatedSearch)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(updated)

Errors: Returns an error if the request fails or if the resource cannot be updated.
*/
func (c *Client) UpdateAdvancedComputerSearchByName(name string, search *ResourceAdvancedComputerSearch) (*ResponseAdvancedComputerSearchCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedComputerSearches, name)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_computer_search"`
		*ResourceAdvancedComputerSearch
	}{
		ResourceAdvancedComputerSearch: search,
	}

	var updatedSearch ResponseAdvancedComputerSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "advance computer search", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

/*
Function: DeleteAdvancedComputerSearchByID
Method: DELETE
Path: /JSSResource/advancedcomputersearches/id/{id}
Description: Deletes an existing Jamf Pro advanced computer search resource by its ID.
Parameters:
  - id (int): The ID of the advanced computer search.

Returns: None
Example:

	err := client.DeleteAdvancedComputerSearchByID(123)
	if err != nil {
	    log.Fatal(err)
	}

Errors: Returns an error if the request fails or if the resource cannot be deleted.
*/
func (c *Client) DeleteAdvancedComputerSearchByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedComputerSearches, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "advance computer search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

/*
Function: DeleteAdvancedComputerSearchByName
Method: DELETE
Path: /JSSResource/advancedcomputersearches/name/{name}
Description: Deletes an existing Jamf Pro advanced computer search resource by its name.
Parameters:
  - name (string): The name of the advanced computer search.

Returns: None
Example:

	err := client.DeleteAdvancedComputerSearchByName("SearchName")
	if err != nil {
	    log.Fatal(err)
	}

Errors: Returns an error if the request fails or if the resource cannot be deleted.
*/
func (c *Client) DeleteAdvancedComputerSearchByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedComputerSearches, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "advance computer search", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
