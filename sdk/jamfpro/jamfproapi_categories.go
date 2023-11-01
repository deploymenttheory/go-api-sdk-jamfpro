// jamfproapi_categories.go
// Jamf Pro Api - osx configuration profiles
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
)

const uriCategories = "/api/v1/categories"

type ResponseCategoriesList struct {
	TotalCount *int           `json:"totalCount,omitempty"`
	Results    []CategoryItem `json:"results,omitempty"`
}

type CategoryItem struct {
	Id       *string `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	Priority *int    `json:"priority,omitempty"`
}

// GetCategories retrieves categories based on query parameters
func (c *Client) GetCategories(page, pageSize int, sort, filter string) (*ResponseCategoriesList, error) {
	endpoint := uriCategories

	// Construct the query parameters
	params := url.Values{}
	if page >= 0 {
		params.Add("page", fmt.Sprintf("%d", page))
	}
	if pageSize > 0 {
		params.Add("page-size", fmt.Sprintf("%d", pageSize))
	}
	if sort != "" {
		params.Add("sort", sort)
	}
	if filter != "" {
		params.Add("filter", filter)
	}

	// Append query parameters to the endpoint
	endpointWithParams := fmt.Sprintf("%s?%s", endpoint, params.Encode())

	var responseCategories ResponseCategoriesList
	resp, err := c.HTTP.DoRequest("GET", endpointWithParams, nil, &responseCategories)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseCategories, nil
}
