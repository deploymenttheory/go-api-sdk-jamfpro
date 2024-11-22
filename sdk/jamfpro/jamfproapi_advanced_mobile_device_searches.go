/*
Filename: jamfproapi_advanced_mobile_searches.go
- Jamf Pro API
- Resource: Advanced Mobile Searches
- API reference: https://developer.jamf.com/jamf-pro/reference/get_v1-advanced-mobile-device-searches
- Data Structure: XML

Shared data structure resources in this endpoint:
- SharedResourceSite
- SharedContainerCriteria
- SharedAdvancedSearchSubsetDisplayField
*/

package jamfpro

import (
	"fmt"
)

const uriAdvancedMobileSearches = "/api/v1/advanced-mobile-device-searches"

// ResponseAdvancedMobileDeviceSearchesList represents the response for Advanced Mobile Device Searches
type ResponseAdvancedMobileDeviceSearchesList struct {
	TotalCount int                                  `json:"totalCount"`
	Results    []ResourceAdvancedMobileDeviceSearch `json:"results"`
}

// ResourceAdvancedMobileSearch represents a single Advanced Mobile Device Search
type ResourceAdvancedMobileDeviceSearch struct {
	ID            string                                     `json:"id"`
	Name          string                                     `json:"name"`
	Criteria      []AdvancedMobileDeviceSearchSubsetCriteria `json:"criteria"`
	DisplayFields []string                                   `json:"displayFields"`
	SiteId        *string                                    `json:"siteId"`
}

// Subsets & Containers

// AdvancedMobileDeviceSearchSubsetCriteria represents the criteria for an Advanced Mobile Device Search
type AdvancedMobileDeviceSearchSubsetCriteria struct {
	Name         string `json:"name"`
	Priority     int    `json:"priority"`
	AndOr        string `json:"andOr"`
	SearchType   string `json:"searchType"`
	Value        string `json:"value"`
	OpeningParen bool   `json:"openingParen"`
	ClosingParen bool   `json:"closingParen"`
}

// Responses

// ResponseAdvancedMobileDeviceSearchChoices represents the response for Advanced Mobile Device Search Choices
type ResponseAdvancedMobileDeviceSearchChoices struct {
	Choices []string `json:"choices"`
}

// GetAdvancedMobileDeviceSearches retrieves all Advanced Mobile Device Searches
func (c *Client) GetAdvancedMobileDeviceSearches() (*ResponseAdvancedMobileDeviceSearchesList, error) {
	endpoint := uriAdvancedMobileSearches

	var response ResponseAdvancedMobileDeviceSearchesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "advanced mobile searches", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// GetAdvancedMobileDeviceSearchByID retrieves a specific Advanced Mobile Device Search by ID
func (c *Client) GetAdvancedMobileDeviceSearchByID(id string) (*ResourceAdvancedMobileDeviceSearch, error) {
	endpoint := fmt.Sprintf("%s/%s", uriAdvancedMobileSearches, id)

	var search ResourceAdvancedMobileDeviceSearch
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &search)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "advanced mobile device search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &search, nil
}

// GetAdvancedMobileDeviceSearchChoices retrieves criteria choices for Advanced Mobile Device Searches
func (c *Client) GetAdvancedMobileDeviceSearchChoices(criteria, site, contains string) (*ResponseAdvancedMobileDeviceSearchChoices, error) {
	endpoint := fmt.Sprintf("%s/choices?criteria=%s&site=%s&contains=%s", uriAdvancedMobileSearches, criteria, site, contains)

	var choices ResponseAdvancedMobileDeviceSearchChoices
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &choices)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "advanced mobile device search choices", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &choices, nil
}

// CreateAdvancedMobileDeviceSearch creates a new Advanced Mobile Device Search
func (c *Client) CreateAdvancedMobileDeviceSearch(search ResourceAdvancedMobileDeviceSearch) (*ResourceAdvancedMobileDeviceSearch, error) {
	endpoint := uriAdvancedMobileSearches

	var response ResourceAdvancedMobileDeviceSearch
	resp, err := c.HTTP.DoRequest("POST", endpoint, search, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "advanced mobile device search", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateAdvancedMobileDeviceSearchByID updates an existing Advanced Mobile Device Search by ID
func (c *Client) UpdateAdvancedMobileDeviceSearchByID(id string, search ResourceAdvancedMobileDeviceSearch) (*ResourceAdvancedMobileDeviceSearch, error) {
	endpoint := fmt.Sprintf("%s/%s", uriAdvancedMobileSearches, id)

	var response ResourceAdvancedMobileDeviceSearch
	resp, err := c.HTTP.DoRequest("PUT", endpoint, search, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "advanced mobile device search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteAdvancedMobileDeviceSearchByID deletes an Advanced Mobile Device Search by ID
func (c *Client) DeleteAdvancedMobileDeviceSearchByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriAdvancedMobileSearches, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "advanced mobile device search", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
