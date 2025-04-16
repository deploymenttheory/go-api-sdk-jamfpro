// jamfproapi_categories.go
// Jamf Pro Api - Categories
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-categories
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriCategories = "/api/v1/categories"

// List

type ResponseCategoriesList struct {
	TotalCount int                `json:"totalCount"`
	Results    []ResourceCategory `json:"results"`
}

// Resource

type ResourceCategory struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

// ResponseCategoryCreateAndUpdate represents the response structure for creating a category.
type ResponseCategoryCreateAndUpdate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD

// GetCategories retrieves all categories from the Jamf Pro API, handling pagination automatically.
// This function makes multiple API calls to fetch each page of category data and aggregates the results.
// It uses a loop to iterate through all available pages of categories.
// The default response contains information for 100 resources, this function is set to the maximum number of 2000.
// Parameters:
// - sort: A string specifying the sorting order of the returned categories.
// - filter: A string to filter the categories based on certain criteria.
func (c *Client) GetCategories(params url.Values) (*ResponseCategoriesList, error) {
	resp, err := c.DoPaginatedGet(uriCategories, standardPageSize, startingPageNumber, params)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "categories", err)
	}

	var out ResponseCategoriesList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		fmt.Println("LOOP")
		fmt.Println(value)
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
		return nil, fmt.Errorf(errMsgFailedGetByID, "categories", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &category, nil
}

// GetCategoryNameByID retrieves a category by its name and then retrieves its details using its ID
func (c *Client) GetCategoryByName(name string) (*ResourceCategory, error) {
	categories, err := c.GetCategories(url.Values{})
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "categories", err)
	}

	for _, value := range categories.Results {
		if value.Name == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "category", name, errMsgNoName)
}

// CreateCategory creates a new category
func (c *Client) CreateCategory(category *ResourceCategory) (*ResponseCategoryCreateAndUpdate, error) {
	endpoint := uriCategories

	var response ResponseCategoryCreateAndUpdate
	resp, err := c.HTTP.DoRequest("POST", endpoint, category, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "category", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateCategoryByID updates an existing category by its ID
func (c *Client) UpdateCategoryByID(id string, categoryUpdate *ResourceCategory) (*ResponseCategoryCreateAndUpdate, error) {
	endpoint := fmt.Sprintf("%s/%s", uriCategories, id)

	var response ResponseCategoryCreateAndUpdate
	resp, err := c.HTTP.DoRequest("PUT", endpoint, categoryUpdate, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "category", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateCategoryByNameByID updates a category by its name and then updates its details using its ID.
func (c *Client) UpdateCategoryByName(name string, categoryUpdate *ResourceCategory) (*ResponseCategoryCreateAndUpdate, error) {
	category, err := c.GetCategoryByName(name)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "category", name, err)
	}

	target_id := category.Id
	resp, err := c.UpdateCategoryByID(target_id, categoryUpdate)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "category", name, err)
	}

	return resp, nil
}

// DeleteCategoryByID deletes a category by its ID
func (c *Client) DeleteCategoryByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriCategories, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "category", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteCategoryByNameByID deletes a category by its name after inferring its ID.
func (c *Client) DeleteCategoryByName(name string) error {
	category, err := c.GetCategoryByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "category", name, err)
	}

	target_id := category.Id
	err = c.DeleteCategoryByID(target_id)

	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "category", name, err)
	}

	return nil
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
		return fmt.Errorf(errMsgFailedDeleteMultiple, "categories", ids, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
