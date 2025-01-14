// jamfapi_bookmarks.go
// Jamf Pro Api - Bookmarks
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-bookmarks

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriBookmarks = "/api/v1/bookmarks"

// List

// ResponseBookmarksList represents the structure of the response for the bookmarks list
type ResponseBookmarksList struct {
	TotalCount int                `json:"totalCount"`
	Results    []ResourceBookmark `json:"results"`
}

// Responses

// ResponseBookmarkCreate represents the response structure for creating a bookmark
type ResponseBookmarkCreate struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// Resource

// ResourceBookmark represents the structure of each bookmark item in the response
type ResourceBookmark struct {
	ID               string `json:"id"`
	SiteID           string `json:"siteId"`
	Priority         int    `json:"priority"`
	DisplayInBrowser *bool  `json:"displayInBrowser"`
	Name             string `json:"name"`
	Description      string `json:"description,omitempty"`
	ScopeDescription string `json:"scopeDescription,omitempty"`
	URL              string `json:"url"`
	IconID           string `json:"iconId"`
}

// CRUD

// GetBookmarks retrieves all bookmark information with optional sorting
func (c *Client) GetBookmarks(sort_filter string) (*ResponseBookmarksList, error) {
	resp, err := c.DoPaginatedGet(
		uriBookmarks,
		standardPageSize,
		startingPageNumber,
		sort_filter,
	)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "bookmarks", err)
	}

	var out ResponseBookmarksList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceBookmark
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "bookmark", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetBookmarkByID retrieves a single bookmark information by its ID
func (c *Client) GetBookmarkByID(id string) (*ResourceBookmark, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBookmarks, id)

	var bookmark ResourceBookmark
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &bookmark)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "bookmark", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &bookmark, nil
}

// GetBookmarkByName retrieves a single bookmark information by its name
func (c *Client) GetBookmarkByName(name string) (*ResourceBookmark, error) {
	bookmarks, err := c.GetBookmarks("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "bookmark", err)
	}

	for _, value := range bookmarks.Results {
		if value.Name == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "bookmark", name, errMsgNoName)
}

// CreateBookmark creates a new bookmark in Jamf Pro
func (c *Client) CreateBookmark(bookmark *ResourceBookmark) (*ResponseBookmarkCreate, error) {
	endpoint := uriBookmarks

	var responseBookmarkCreate ResponseBookmarkCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, bookmark, &responseBookmarkCreate)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "bookmark", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseBookmarkCreate, nil
}

// UpdateBookmarkByID updates a bookmark's information in Jamf Pro by its ID
func (c *Client) UpdateBookmarkByID(id string, bookmarkUpdate *ResourceBookmark) (*ResourceBookmark, error) {
	endpoint := fmt.Sprintf("%s/%s", uriBookmarks, id)

	var updatedBookmark ResourceBookmark
	resp, err := c.HTTP.DoRequest("PUT", endpoint, bookmarkUpdate, &updatedBookmark)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "bookmark", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedBookmark, nil
}

// UpdateBookmarkByName updates a bookmark's information in Jamf Pro by its name
func (c *Client) UpdateBookmarkByName(name string, bookmarkUpdate *ResourceBookmark) (*ResourceBookmark, error) {
	target, err := c.GetBookmarkByName(name)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "bookmark", name, err)
	}

	target_id := target.ID
	resp, err := c.UpdateBookmarkByID(target_id, bookmarkUpdate)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "bookmark", name, err)
	}

	return resp, nil
}

// DeleteBookmarkByID deletes a bookmark in Jamf Pro by its ID
func (c *Client) DeleteBookmarkByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriBookmarks, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "bookmark", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteBookmarkByName deletes a bookmark in Jamf Pro by its name
func (c *Client) DeleteBookmarkByName(name string) error {
	target, err := c.GetBookmarkByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "bookmark", name, err)
	}

	target_id := target.ID
	err = c.DeleteBookmarkByID(target_id)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "bookmark", name, err)
	}

	return nil
}
