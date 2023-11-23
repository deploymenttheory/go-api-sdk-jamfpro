// jamfproapi_categories.go
// Jamf Pro Api - Categories
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"
	"strconv"
)

const uriCategories = "/api/v1/categories"

type ResponseCategoriesList struct {
	TotalCount *int           `json:"totalCount,omitempty"`
	Results    []CategoryItem `json:"results,omitempty"`
}

type CategoryItem struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

type ResponseCategories struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

// GetCategories retrieves all categories from the Jamf Pro API, handling pagination automatically.
// This function makes multiple API calls to fetch each page of category data and aggregates the results.
// It uses a loop to iterate through all available pages of categories.
// The default response contains information for 100 resources, this function is set to the maximum number of 2000.
// Parameters:
// - sort: A string specifying the sorting order of the returned categories.
// - filter: A string to filter the categories based on certain criteria.
func (c *Client) GetCategories(sort, filter string) (*ResponseCategoriesList, error) {
	const maxPageSize = 2000
	var allCategories []CategoryItem

	page := 0
	for {
		// Construct the endpoint with query parameters for the current page
		endpointWithParams := fmt.Sprintf("%s?%s", uriCategories, url.Values{
			"page":      []string{strconv.Itoa(page)},
			"page-size": []string{strconv.Itoa(maxPageSize)},
			"sort":      []string{sort},
			"filter":    []string{filter},
		}.Encode())

		// Fetch the categories for the current page
		var responseCategories ResponseCategoriesList
		resp, err := c.HTTP.DoRequest("GET", endpointWithParams, nil, &responseCategories)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch categories: %v", err)
		}

		if resp != nil && resp.Body != nil {
			defer resp.Body.Close()
		}

		// Add the fetched categories to the total list
		allCategories = append(allCategories, responseCategories.Results...)

		// Check if all categories have been fetched
		if responseCategories.TotalCount == nil || len(allCategories) >= *responseCategories.TotalCount {
			break
		}

		// Increment page number for the next iteration
		page++
	}

	// Return the combined list of all categories
	return &ResponseCategoriesList{
		TotalCount: &[]int{len(allCategories)}[0],
		Results:    allCategories,
	}, nil
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
	categoriesList, err := c.GetCategories("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all categories: %v", err)
	}

	// Search for the category with the given name
	for _, category := range categoriesList.Results {
		if category.Name == name {
			return c.GetCategoryByID(category.Id)
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

// UpdateCategoryByID updates an existing category by its ID
func (c *Client) UpdateCategoryByID(id int, updatedCategory *ResponseCategories) (*ResponseCategories, error) {
	endpoint := fmt.Sprintf("%s/%d", uriCategories, id)

	var response ResponseCategories
	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedCategory, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update category: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateCategoryByNameByID updates a category by its name and then updates its details using its ID.
func (c *Client) UpdateCategoryByNameByID(name string, updatedCategory *ResponseCategories) (*ResponseCategories, error) {
	// Fetch all categories
	categoriesList, err := c.GetCategories("", "") // Adjusted call to match new signature of GetCategories
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all categories: %v", err)
	}

	// Search for the category with the given name
	for _, category := range categoriesList.Results {
		if category.Name == name {
			// Parse the ID from string to int
			id, err := strconv.Atoi(category.Id)
			if err != nil {
				return nil, fmt.Errorf("failed to parse category ID: %v", err)
			}
			// Update the category using its ID
			return c.UpdateCategoryByID(id, updatedCategory)
		}
	}

	return nil, fmt.Errorf("no category found with the name %s", name)
}

// DeleteCategoryByID deletes a category by its ID
func (c *Client) DeleteCategoryByID(id int) error {
	endpoint := fmt.Sprintf("%s/%d", uriCategories, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete category: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteCategoryByNameByID deletes a category by its name after inferring its ID.
func (c *Client) DeleteCategoryByNameByID(name string) error {
	// Fetch all categories
	categoriesList, err := c.GetCategories("", "") // Call updated to match new signature of GetCategories
	if err != nil {
		return fmt.Errorf("failed to fetch all categories: %v", err)
	}

	// Search for the category with the given name
	for _, category := range categoriesList.Results {
		if category.Name == name {
			// Parse the ID from string to int
			id, err := strconv.Atoi(category.Id)
			if err != nil {
				return fmt.Errorf("failed to parse category ID: %v", err)
			}
			// Delete the category using its ID
			return c.DeleteCategoryByID(id)
		}
	}

	return fmt.Errorf("no category found with the name %s", name)
}

// DeleteMultipleCategoriesByID deletes multiple categories by their IDs
func (c *Client) DeleteMultipleCategoriesByID(ids []string) error {
	endpoint := fmt.Sprintf("%s/delete-multiple", uriCategories)

	// Construct the request payload
	payload := struct {
		IDs []string `json:"ids"`
	}{
		IDs: ids,
	}

	resp, err := c.HTTP.DoRequest("POST", endpoint, payload, nil)
	if err != nil {
		return fmt.Errorf("failed to delete multiple categories: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
