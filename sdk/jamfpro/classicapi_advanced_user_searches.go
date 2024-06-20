/*
Filename: classicapi_advanced_user_searches.go
- Jamf Pro Classic API
- Resource: Advanced User Searches
- API reference: https://developer.jamf.com/jamf-pro/reference/advancedusersearches
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
	ID            int                                         `xml:"id"`
	Name          string                                      `xml:"name,omitempty"`
	Criteria      SharedContainerCriteria                     `xml:"criteria,omitempty"`
	Users         []AdvancedUserSearchContainerUsers          `xml:"users,omitempty"`
	DisplayFields []SharedAdvancedSearchContainerDisplayField `xml:"display_fields,omitempty"`
	Site          SharedResourceSite                          `xml:"site,omitempty"`
}

// Responses

type ResponseAdvancedUserSearchCreatedAndUpdated struct {
	ID int `json:"id,omitempty" xml:"id,omitempty"`
}

// Subsets & Containers

type AdvancedUserSearchContainerUsers struct {
	Size int                            `xml:"size,omitempty"`
	User []AdvancedUserSearchSubsetUser `xml:"user,omitempty"`
}

type AdvancedUserSearchSubsetUser struct {
	ID       int    `xml:"id,omitempty"`
	Name     string `xml:"name,omitempty"`
	Username string `xml:"Username,omitempty"`
}

// CRUD

/*
Function: GetAdvancedUserSearches
Method: GET
Path: /JSSResource/advancedusersearches
Description: Gets a list of all Jamf Pro Advanced User Search resources.
Parameters: None
Returns: ResponseAdvancedUserSearchesList - A list of advanced user searches.
Example:

	searches, err := client.GetAdvancedUserSearches()
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(searches)

Errors: Returns an error if the request fails.
*/
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

/*
Function: GetAdvancedUserSearchByID
Method: GET
Path: /JSSResource/advancedusersearches/id/{id}
Description: Gets a Jamf Pro advanced user search resource by its ID.
Parameters:
  - id (int): The ID of the advanced user search.

Returns: ResourceAdvancedUserSearch - The advanced user search resource.
Example:

	search, err := client.GetAdvancedUserSearchByID(123)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(search)

Errors: Returns an error if the request fails or if the ID is not found.
*/
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

/*
Function: GetAdvancedUserSearchByName
Method: GET
Path: /JSSResource/advancedusersearches/name/{name}
Description: Gets a Jamf Pro advanced user search resource by its name.
Parameters:
  - name (string): The name of the advanced user search.

Returns: ResourceAdvancedUserSearch - The advanced user search resource.
Example:

	search, err := client.GetAdvancedUserSearchByName("SearchName")
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(search)

Errors: Returns an error if the request fails or if the name is not found.
*/
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

/*
Function: CreateAdvancedUserSearch
Method: POST
Path: /JSSResource/advancedusersearches
Description: Creates a new Jamf Pro advanced user search resource.
Parameters:
  - search (*ResourceAdvancedUserSearch): The advanced user search resource to create.

Returns: ResponseAdvancedUserSearchCreatedAndUpdated - The ID of the created advanced user search resource.
Example:

	newSearch := &jamfpro.ResourceAdvancedUserSearch{
	    Name: "New Search",
	    // Other fields...
	}
	created, err := client.CreateAdvancedUserSearch(newSearch)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(created)

Errors: Returns an error if the request fails or if the resource cannot be created.
*/
func (c *Client) CreateAdvancedUserSearch(search *ResourceAdvancedUserSearch) (*ResponseAdvancedUserSearchCreatedAndUpdated, error) {
	endpoint := uriAPIAdvancedUserSearches

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_user_search"`
		*ResourceAdvancedUserSearch
	}{
		ResourceAdvancedUserSearch: search,
	}

	var createdSearch ResponseAdvancedUserSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "advanced user search", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdSearch, nil
}

/*
Function: UpdateAdvancedUserSearchByID
Method: PUT
Path: /JSSResource/advancedusersearches/id/{id}
Description: Updates an existing Jamf Pro advanced user search resource by its ID.
Parameters:
  - id (int): The ID of the advanced user search.
  - search (*ResourceAdvancedUserSearch): The updated advanced user search resource.

Returns: ResponseAdvancedUserSearchCreatedAndUpdated - The ID of the updated advanced user search resource.
Example:

	updatedSearch := &jamfpro.ResourceAdvancedUserSearch{
	    Name: "Updated Search",
	    // Other fields...
	}
	updated, err := client.UpdateAdvancedUserSearchByID(123, updatedSearch)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(updated)

Errors: Returns an error if the request fails or if the resource cannot be updated.
*/
func (c *Client) UpdateAdvancedUserSearchByID(id int, search *ResourceAdvancedUserSearch) (*ResponseAdvancedUserSearchCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriAPIAdvancedUserSearches, id)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_user_search"`
		*ResourceAdvancedUserSearch
	}{
		ResourceAdvancedUserSearch: search,
	}

	var updatedSearch ResponseAdvancedUserSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "advanced user search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

/*
Function: UpdateAdvancedUserSearchByName
Method: PUT
Path: /JSSResource/advancedusersearches/name/{name}
Description: Updates an existing Jamf Pro advanced user search resource by its name.
Parameters:
  - name (string): The name of the advanced user search.
  - search (*ResourceAdvancedUserSearch): The updated advanced user search resource.

Returns: ResponseAdvancedUserSearchCreatedAndUpdated - The ID of the updated advanced user search resource.
Example:

	updatedSearch := &jamfpro.ResourceAdvancedUserSearch{
	    Name: "Updated Search",
	    // Other fields...
	}
	updated, err := client.UpdateAdvancedUserSearchByName("SearchName", updatedSearch)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println(updated)

Errors: Returns an error if the request fails or if the resource cannot be updated.
*/
func (c *Client) UpdateAdvancedUserSearchByName(name string, search *ResourceAdvancedUserSearch) (*ResponseAdvancedUserSearchCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriAPIAdvancedUserSearches, name)

	requestBody := struct {
		XMLName xml.Name `xml:"advanced_user_search"`
		*ResourceAdvancedUserSearch
	}{
		ResourceAdvancedUserSearch: search,
	}

	var updatedSearch ResponseAdvancedUserSearchCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedSearch)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "advanced user search", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedSearch, nil
}

/*
Function: DeleteAdvancedUserSearchByID
Method: DELETE
Path: /JSSResource/advancedusersearches/id/{id}
Description: Deletes an existing Jamf Pro advanced user search resource by its ID.
Parameters:
  - id (int): The ID of the advanced user search.

Returns: None
Example:

	err := client.DeleteAdvancedUserSearchByID(123)
	if err != nil {
	    log.Fatal(err)
	}

Errors: Returns an error if the request fails or if the resource cannot be deleted.
*/
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

/*
Function: DeleteAdvancedUserSearchByName
Method: DELETE
Path: /JSSResource/advancedusersearches/name/{name}
Description: Deletes an existing Jamf Pro advanced user search resource by its name.
Parameters:
  - name (string): The name of the advanced user search.

Returns: None
Example:

	err := client.DeleteAdvancedUserSearchByName("SearchName")
	if err != nil {
	    log.Fatal(err)
	}

Errors: Returns an error if the request fails or if the resource cannot be deleted.
*/
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
