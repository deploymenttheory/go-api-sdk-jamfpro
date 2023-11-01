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

type ResponseCategories struct {
	Id       *string `json:"id"`
	Name     *string `json:"name"`
	Priority *int    `json:"priority"`
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

// GetCategoryByID retrieves a category by its ID
func (c *Client) GetCategoryByID(id string) (*ResponseCategories, error) {
	endpoint := fmt.Sprintf("%s/%s", uriCategories, id)

	var category ResponseCategories
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &category)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch category by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &category, nil
}

// GetCategoryNameByID retrieves a category by its name and then retrieves its details using its ID
func (c *Client) GetCategoryNameByID(name string) (*ResponseCategories, error) {
	// Fetch all categories
	categoriesList, err := c.GetCategories(0, 100, "", "") // You may adjust page, pageSize, sort, and filter as needed
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all categories: %v", err)
	}

	// Search for the category with the given name
	for _, category := range categoriesList.Results {
		if category.Name != nil && *category.Name == name {
			return c.GetCategoryByID(*category.Id)
		}
	}

	return nil, fmt.Errorf("no category found with the name %s", name)
}

// CreateCategory creates a new category
func (c *Client) CreateCategory(category *ResponseCategories) (*ResponseCategories, error) {
	endpoint := uriCategories

	var response ResponseCategories
	resp, err := c.HTTP.DoRequest("POST", endpoint, category, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create category: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}
