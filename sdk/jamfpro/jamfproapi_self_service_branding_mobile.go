// jamfproapi_self_service_branding_ios.go
// Jamf Pro Api - Self Service Branding iOS
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-ios
// Jamf Pro Api requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
	"net/url"

	"github.com/mitchellh/mapstructure"
)

const uriSelfServiceBrandingIOS = "/api/v1/self-service/branding/ios"

// Response

// ResponseSelfServiceBrandingIOSList holds the list of self-service branding configurations for iOS.
type ResponseSelfServiceBrandingIOSList struct {
	TotalCount int                                    `json:"totalCount"`
	Results    []ResourceSelfServiceBrandingIOSDetail `json:"results"`
}

// Resource

// ResourceSelfServiceBrandingIOSDetail represents the details of a self-service branding configuration for iOS.
type ResourceSelfServiceBrandingIOSDetail struct {
	ID                        string `json:"id"`
	BrandingName              string `json:"brandingName"`
	IconId                    *int   `json:"iconId,omitempty"`
	HeaderBackgroundColorCode string `json:"headerBackgroundColorCode"`
	MenuIconColorCode         string `json:"menuIconColorCode"`
	BrandingNameColorCode     string `json:"brandingNameColorCode"`
	StatusBarTextColor        string `json:"statusBarTextColor"`
}

// CreateResponseSelfServiceBrandingIOS is the minimal response returned by the API when creating a branding
// configuration. It mirrors the lightweight server response { "id": "...", "href": "..." }.
type CreateResponseSelfServiceBrandingIOS struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// CRUD

// GetSelfServiceBrandingIOS retrieves the list of self-service branding configurations for iOS.
func (c *Client) GetSelfServiceBrandingIOS(params url.Values) (*ResponseSelfServiceBrandingIOSList, error) {
	resp, err := c.DoPaginatedGet(uriSelfServiceBrandingIOS, params)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "self service branding (ios)", err)
	}

	var out ResponseSelfServiceBrandingIOSList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceSelfServiceBrandingIOSDetail
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "self service branding (ios)", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetSelfServiceBrandingIOSByID retrieves a specific self-service branding configuration for iOS by ID.
func (c *Client) GetSelfServiceBrandingIOSByID(id string) (*ResourceSelfServiceBrandingIOSDetail, error) {
	var out ResourceSelfServiceBrandingIOSDetail
	endpoint := fmt.Sprintf("%s/%s", uriSelfServiceBrandingIOS, id)

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf(errMsgFailedGetByID, "self service branding (ios)", id, err)
		return nil, err
	}

	return &out, nil
}

// GetSelfServiceBrandingIOSByName retrieves a specific self-service branding configuration for iOS by its name.
func (c *Client) GetSelfServiceBrandingIOSByName(name string) (*ResourceSelfServiceBrandingIOSDetail, error) {
	all, err := c.GetSelfServiceBrandingIOS(nil)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "self service brandings (ios)", err)
	}

	for _, value := range all.Results {
		if value.BrandingName == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "self service branding (ios)", name, errMsgNoName)
}

// CreateSelfServiceBrandingIOS creates a new self-service branding configuration for iOS.
func (c *Client) CreateSelfServiceBrandingIOS(branding *ResourceSelfServiceBrandingIOSDetail) (*CreateResponseSelfServiceBrandingIOS, error) {
	endpoint := uriSelfServiceBrandingIOS

	var minimal CreateResponseSelfServiceBrandingIOS
	resp, err := c.HTTP.DoRequest("POST", endpoint, branding, &minimal)
	if err != nil {
		return nil, fmt.Errorf("failed to create self-service branding (ios): %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &minimal, nil
}

// UpdateSelfServiceBrandingIOSByID updates an existing self-service branding configuration for iOS.
func (c *Client) UpdateSelfServiceBrandingIOSByID(id string, brandingUpdate *ResourceSelfServiceBrandingIOSDetail) (*ResourceSelfServiceBrandingIOSDetail, error) {
	endpoint := fmt.Sprintf("%s/%s", uriSelfServiceBrandingIOS, id)

	var response ResourceSelfServiceBrandingIOSDetail
	resp, err := c.HTTP.DoRequest("PUT", endpoint, brandingUpdate, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update self-service branding (ios): %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateSelfServiceBrandingIOSByName updates a self-service branding configuration for iOS by name.
func (c *Client) UpdateSelfServiceBrandingIOSByName(name string, brandingUpdate *ResourceSelfServiceBrandingIOSDetail) (*ResourceSelfServiceBrandingIOSDetail, error) {
	target, err := c.GetSelfServiceBrandingIOSByName(name)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "self service branding (ios)", name, err)
	}

	targetID := target.ID
	resp, err := c.UpdateSelfServiceBrandingIOSByID(targetID, brandingUpdate)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "self service branding (ios)", name, err)
	}

	return resp, nil
}

// DeleteSelfServiceBrandingIOSByID deletes a self-service branding configuration for iOS by ID.
func (c *Client) DeleteSelfServiceBrandingIOSByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriSelfServiceBrandingIOS, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "self service branding (ios)", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteSelfServiceBrandingIOSByName deletes a self-service branding configuration for iOS by name.
func (c *Client) DeleteSelfServiceBrandingIOSByName(name string) error {
	target, err := c.GetSelfServiceBrandingIOSByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "self service branding (ios)", name, err)
	}

	targetID := target.ID
	err = c.DeleteSelfServiceBrandingIOSByID(targetID)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "self service branding (ios)", name, err)
	}

	return nil
}
