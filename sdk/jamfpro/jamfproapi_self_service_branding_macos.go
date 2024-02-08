// jamfproapi_self_service_branding_macos.go
// Jamf Pro Api - Self Service Branding macOS
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-macos
// Classic API requires the structs to support an JSON data structure.

package jamfpro

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

const uriSelfServiceBrandingMacOS = "/api/v1/self-service/branding/macos"

// Response

// ResponseSelfServiceBranding is the structure that holds the list of self-service branding configurations for macOS.
type ResponseSelfServiceBrandingList struct {
	TotalCount int                                 `json:"totalCount"`
	Results    []ResourceSelfServiceBrandingDetail `json:"results"`
}

// Resource

// SelfServiceBrandingDetail represents the details of a self-service branding configuration.
type ResourceSelfServiceBrandingDetail struct {
	ID                    string `json:"id"`
	ApplicationName       string `json:"applicationName"`
	BrandingName          string `json:"brandingName"`
	BrandingNameSecondary string `json:"brandingNameSecondary"`
	IconId                int    `json:"iconId"`
	BrandingHeaderImageId int    `json:"brandingHeaderImageId"`
}

// CRUD

// GetSelfServiceBrandingMacOS retrieves the list of self-service branding configurations for macOS.
func (c *Client) GetSelfServiceBrandingMacOS(sort_filter string) (*ResponseSelfServiceBrandingList, error) {
	resp, err := c.DoPaginatedGet(
		uriSelfServiceBrandingMacOS,
		standardPageSize,
		startingPageNumber,
		sort_filter,
	)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "self service branding", err)
	}

	var out ResponseSelfServiceBrandingList
	out.TotalCount = resp.Size

	for _, value := range resp.Results {
		var newObj ResourceSelfServiceBrandingDetail
		err := mapstructure.Decode(value, &newObj)
		if err != nil {
			return nil, fmt.Errorf(errMsgFailedMapstruct, "self service branding", err)
		}
		out.Results = append(out.Results, newObj)
	}

	return &out, nil
}

// GetSelfServiceBrandingMacOSByID retrieves a specific self-service branding configuration for macOS by ID.
func (c *Client) GetSelfServiceBrandingMacOSByID(id string) (*ResourceSelfServiceBrandingDetail, error) {
	var out ResourceSelfServiceBrandingDetail
	endpoint := fmt.Sprintf("%s/%s", uriSelfServiceBrandingMacOS, id)

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out, c.HTTP.Logger)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf(errMsgFailedGetByID, "self service branding", id, err)
		return nil, err
	}

	return &out, nil
}

// GetSelfServiceBrandingMacOSByNameByID retrieves a specific self-service branding configuration for macOS by its name.
func (c *Client) GetSelfServiceBrandingMacOSByName(name string) (*ResourceSelfServiceBrandingDetail, error) {
	all_ssbrandings, err := c.GetSelfServiceBrandingMacOS("")
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedPaginatedGet, "self service brandings", err)
	}

	for _, value := range all_ssbrandings.Results {
		if value.BrandingName == name {
			return &value, nil
		}
	}

	return nil, fmt.Errorf(errMsgFailedGetByName, "self service branding", name, errMsgNoName)
}

// CreateSelfServiceBrandingMacOS creates a new self-service branding configuration for macOS.
func (c *Client) CreateSelfServiceBrandingMacOS(branding *ResourceSelfServiceBrandingDetail) (*ResourceSelfServiceBrandingDetail, error) {
	endpoint := uriSelfServiceBrandingMacOS

	var response ResourceSelfServiceBrandingDetail
	resp, err := c.HTTP.DoRequest("POST", endpoint, branding, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to create self-service branding: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateSelfServiceBrandingMacOSByID updates an existing self-service branding configuration for macOS.
func (c *Client) UpdateSelfServiceBrandingMacOSByID(id string, brandingUpdate *ResourceSelfServiceBrandingDetail) (*ResourceSelfServiceBrandingDetail, error) {
	endpoint := fmt.Sprintf("%s/%s", uriSelfServiceBrandingMacOS, id)

	var response ResourceSelfServiceBrandingDetail
	resp, err := c.HTTP.DoRequest("PUT", endpoint, brandingUpdate, &response, c.HTTP.Logger)
	if err != nil {
		return nil, fmt.Errorf("failed to update self-service branding: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateSelfServiceBrandingMacOSByName updates a self-service branding configuration for macOS by name.
func (c *Client) UpdateSelfServiceBrandingMacOSByName(name string, brandingUpdate *ResourceSelfServiceBrandingDetail) (*ResourceSelfServiceBrandingDetail, error) {
	target, err := c.GetSelfServiceBrandingMacOSByName(name)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByName, "self service branding", name, err)
	}

	target_id := target.ID
	resp, err := c.UpdateSelfServiceBrandingMacOSByID(target_id, brandingUpdate)

	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "self service branding", name, err)
	}

	return resp, nil
}

// DeleteSelfServiceBrandingMacOSByID deletes a self-service branding configuration for macOS by ID.
func (c *Client) DeleteSelfServiceBrandingMacOSByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriSelfServiceBrandingMacOS, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil, c.HTTP.Logger)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "self service branding", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteSelfServiceBrandingMacOSByName deletes a self-service branding configuration for macOS by name.
func (c *Client) DeleteSelfServiceBrandingMacOSByName(name string) error {
	target, err := c.GetSelfServiceBrandingMacOSByName(name)
	if err != nil {
		return fmt.Errorf(errMsgFailedGetByName, "self service branding", name, err)
	}

	target_id := target.ID
	err = c.DeleteSelfServiceBrandingMacOSByID(target_id)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "self service branding", name, err)
	}

	return nil
}
