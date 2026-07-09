// jamfproapi_users.go
// Jamf Pro Api - Users
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-users
// Jamf Pro API requires the structs to support a JSON data structure.
//
// Note: This is the modern Jamf Pro API (/api/v1/users) users resource. The Classic API
// users resource (/JSSResource/users) lives in classicapi_users.go.

package jamfpro

import (
	"fmt"
	"net/url"
)

const uriUsersV1 = "/api/v1/users"

// List

// ResponseUsersListV1 represents a page of Jamf Pro API user results.
type ResponseUsersListV1 struct {
	Page        int              `json:"page"`
	PageSize    int              `json:"pageSize"`
	TotalCount  int              `json:"totalCount"`
	TotalPages  int              `json:"totalPages"`
	HasNext     bool             `json:"hasNext"`
	HasPrevious bool             `json:"hasPrevious"`
	Results     []ResourceUserV1 `json:"results"`
}

// ResponseUsersCreateV1 represents the response returned when creating a user.
type ResponseUsersCreateV1 struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ResponseUserRecalculationResultsV1 represents the results of a smart group recalculation.
type ResponseUserRecalculationResultsV1 struct {
	Count int `json:"count"`
}

// Resource

// ResourceUserV1 represents a Jamf Pro API user. When used as a request body for create/update
// the read-only ID field is omitted.
type ResourceUserV1 struct {
	ID                   string `json:"id,omitempty"`
	Username             string `json:"username,omitempty"`
	Realname             string `json:"realname,omitempty"`
	Email                string `json:"email,omitempty"`
	Phone                string `json:"phone,omitempty"`
	Position             string `json:"position,omitempty"`
	EnableCustomPhotoURL bool   `json:"enableCustomPhotoUrl,omitempty"`
	CustomPhotoURL       string `json:"customPhotoUrl,omitempty"`
	ManagedAppleID       string `json:"managedAppleId,omitempty"`
}

// GetUsersV1 retrieves a page of users from the Jamf Pro API.
func (c *Client) GetUsersV1(params url.Values) (*ResponseUsersListV1, error) {
	endpoint := uriUsersV1
	if params != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, params.Encode())
	}

	var out ResponseUsersListV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "users", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetUserByIDV1 retrieves a single user by ID from the Jamf Pro API.
func (c *Client) GetUserByIDV1(id string) (*ResourceUserV1, error) {
	endpoint := fmt.Sprintf("%s/%s", uriUsersV1, id)

	var out ResourceUserV1
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "user", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateUserV1 creates a new user via the Jamf Pro API.
func (c *Client) CreateUserV1(user *ResourceUserV1) (*ResponseUsersCreateV1, error) {
	endpoint := uriUsersV1

	var out ResponseUsersCreateV1
	resp, err := c.HTTP.DoRequest("POST", endpoint, user, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "user", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// UpdateUserByIDV1 updates an existing user by ID via the Jamf Pro API.
func (c *Client) UpdateUserByIDV1(id string, user *ResourceUserV1) error {
	endpoint := fmt.Sprintf("%s/%s", uriUsersV1, id)

	resp, err := c.HTTP.DoRequest("PUT", endpoint, user, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedUpdateByID, "user", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteUserByIDV1 deletes a user by ID via the Jamf Pro API.
func (c *Client) DeleteUserByIDV1(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriUsersV1, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil || resp.StatusCode != 204 {
		return fmt.Errorf(errMsgFailedDeleteByID, "user", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// RecalculateUserSmartGroupsByIDV1 triggers a smart group membership recalculation for a user.
func (c *Client) RecalculateUserSmartGroupsByIDV1(id string) (*ResponseUserRecalculationResultsV1, error) {
	endpoint := fmt.Sprintf("%s/%s/recalculate-smart-groups", uriUsersV1, id)

	var out ResponseUserRecalculationResultsV1
	resp, err := c.HTTP.DoRequest("POST", endpoint, nil, &out)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedActionByID, "user", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
