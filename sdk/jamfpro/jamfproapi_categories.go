// jamfproapi_categories.go
// Jamf Pro Api - Categories
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriCategories = "/api/v1/categories"

type ResponseCategoriesList struct {
	TotalCount int                `json:"totalCount,omitempty"`
	Results    []ResourceCategory `json:"results,omitempty"`
}

type ResourceCategory struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Priority int    `json:"priority,omitempty"`
}

// GetCategories retrieves all categories from the Jamf Pro API, handling pagination automatically.
// This function makes multiple API calls to fetch each page of category data and aggregates the results.
// It uses a loop to iterate through all available pages of categories.
// The default response contains information for 100 resources, this function is set to the maximum number of 2000.
// Parameters:
// - sort: A string specifying the sorting order of the returned categories.
// - filter: A string to filter the categories based on certain criteria.
func (c *Client) GetCategories(sort_filter string) (*ResponseCategoriesList, error) {
	resp, err := c.DoPaginatedGet(uriCategories, standardPageSize, startingPageNumber, sort_filter)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "categories", err)
	}

	var out ResponseCategoriesList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceCategory
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "category", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetCategoryByID retrieves a category by its ID
func (c *Client) GetCategoryByID(id string) (*ResourceCategory, error) {
	endpoint := fmt.Sprintf("%s/%s", uriCategories, id)

	var category ResourceCategory
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &category)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch category by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &category, nil
}

// // GetCategoryNameByID retrieves a category by its name and then retrieves its details using its ID
// func (c *Client) GetCategoryNameByID(name string) (*ResponseCategories, error) {
// 	// Fetch all categories
// 	categoriesList, err := c.GetCategories("", "")
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch all categories: %v", err)
// 	}

// 	// Search for the category with the given name
// 	for _, category := range categoriesList.Results {
// 		if category.Name == name {
// 			return c.GetCategoryByID(category.Id)
// 		}
// 	}

// 	return nil, fmt.Errorf("no category found with the name %s", name)
// }

// // CreateCategory creates a new category
// func (c *Client) CreateCategory(category *ResponseCategories) (*ResponseCategories, error) {
// 	endpoint := uriCategories

// 	var response ResponseCategories
// 	resp, err := c.HTTP.DoRequest("POST", endpoint, category, &response)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create category: %v", err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return &response, nil
// }

// // UpdateCategoryByID updates an existing category by its ID
// func (c *Client) UpdateCategoryByID(id int, updatedCategory *ResponseCategories) (*ResponseCategories, error) {
// 	endpoint := fmt.Sprintf("%s/%d", uriCategories, id)

// 	var response ResponseCategories
// 	resp, err := c.HTTP.DoRequest("PUT", endpoint, updatedCategory, &response)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to update category: %v", err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return &response, nil
// }

// // UpdateCategoryByNameByID updates a category by its name and then updates its details using its ID.
// func (c *Client) UpdateCategoryByNameByID(name string, updatedCategory *ResponseCategories) (*ResponseCategories, error) {
// 	// Fetch all categories
// 	categoriesList, err := c.GetCategories("", "") // Adjusted call to match new signature of GetCategories
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch all categories: %v", err)
// 	}

// 	// Search for the category with the given name
// 	for _, category := range categoriesList.Results {
// 		if category.Name == name {
// 			// Parse the ID from string to int
// 			id, err := strconv.Atoi(category.Id)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to parse category ID: %v", err)
// 			}
// 			// Update the category using its ID
// 			return c.UpdateCategoryByID(id, updatedCategory)
// 		}
// 	}

// 	return nil, fmt.Errorf("no category found with the name %s", name)
// }

// // DeleteCategoryByID deletes a category by its ID
// func (c *Client) DeleteCategoryByID(id int) error {
// 	endpoint := fmt.Sprintf("%s/%d", uriCategories, id)

// 	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
// 	if err != nil {
// 		return fmt.Errorf("failed to delete category: %v", err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return nil
// }

// // DeleteCategoryByNameByID deletes a category by its name after inferring its ID.
// func (c *Client) DeleteCategoryByNameByID(name string) error {
// 	// Fetch all categories
// 	categoriesList, err := c.GetCategories("", "") // Call updated to match new signature of GetCategories
// 	if err != nil {
// 		return fmt.Errorf("failed to fetch all categories: %v", err)
// 	}

// 	// Search for the category with the given name
// 	for _, category := range categoriesList.Results {
// 		if category.Name == name {
// 			// Parse the ID from string to int
// 			id, err := strconv.Atoi(category.Id)
// 			if err != nil {
// 				return fmt.Errorf("failed to parse category ID: %v", err)
// 			}
// 			// Delete the category using its ID
// 			return c.DeleteCategoryByID(id)
// 		}
// 	}

// 	return fmt.Errorf("no category found with the name %s", name)
// }

// // DeleteMultipleCategoriesByID deletes multiple categories by their IDs
// func (c *Client) DeleteMultipleCategoriesByID(ids []string) error {
// 	endpoint := fmt.Sprintf("%s/delete-multiple", uriCategories)

// 	// Construct the request payload
// 	payload := struct {
// 		IDs []string `json:"ids"`
// 	}{
// 		IDs: ids,
// 	}

// 	resp, err := c.HTTP.DoRequest("POST", endpoint, payload, nil)
// 	if err != nil {
// 		return fmt.Errorf("failed to delete multiple categories: %v", err)
// 	}

// 	if resp != nil && resp.Body != nil {
// 		defer resp.Body.Close()
// 	}

// 	return nil
// }
