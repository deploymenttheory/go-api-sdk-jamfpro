// jamfproapi_self_service_branding_macos.go
// Jamf Pro Api - Self Service Branding macOS
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-self-service-branding-macos
// Classic API requires the structs to support an JSON data structure.

package jamfpro

import "fmt"

const uriAPISelfServiceBrandingMacOS = "/api/v1/self-service/branding/macos"

// ResponseSelfServiceBranding is the structure that holds the list of self-service branding configurations for macOS.
type ResponseSelfServiceBranding struct {
	TotalCount int                         `json:"totalCount"`
	Results    []SelfServiceBrandingDetail `json:"results"`
}

// SelfServiceBrandingDetail represents the details of a self-service branding configuration.
type SelfServiceBrandingDetail struct {
	ID                    string `json:"id"`
	ApplicationName       string `json:"applicationName"`
	BrandingName          string `json:"brandingName"`
	BrandingNameSecondary string `json:"brandingNameSecondary"`
	IconId                int    `json:"iconId"`
	BrandingHeaderImageId int    `json:"brandingHeaderImageId"`
}

// GetSelfServiceBrandingMacOS retrieves the list of self-service branding configurations for macOS.
func (c *Client) GetSelfServiceBrandingMacOS() (*ResponseSelfServiceBranding, error) {
	var out ResponseSelfServiceBranding

	resp, err := c.HTTP.DoRequest("GET", uriAPISelfServiceBrandingMacOS, nil, &out)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf("Failed to fetch self-service branding for macOS", "Error", err)
		return nil, err
	}

	return &out, nil
}

// GetSelfServiceBrandingMacOSByID retrieves a specific self-service branding configuration for macOS by ID.
func (c *Client) GetSelfServiceBrandingMacOSByID(id string) (*SelfServiceBrandingDetail, error) {
	var out SelfServiceBrandingDetail
	// Construct the URL with the ID
	endpoint := fmt.Sprintf("%s/%s", uriAPISelfServiceBrandingMacOS, id)

	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		fmt.Printf("Failed to fetch self-service branding for macOS by ID", "Error", err)
		return nil, err
	}

	return &out, nil
}

// GetSelfServiceBrandingMacOSByNameByID retrieves a specific self-service branding configuration for macOS by its name.
func (c *Client) GetSelfServiceBrandingMacOSByNameByID(name string) (*SelfServiceBrandingDetail, error) {
	// First, get all branding configurations.
	response, err := c.GetSelfServiceBrandingMacOS()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch self-service branding for macOS: %v", err)
	}

	// Now, find the branding with the given name.
	for _, branding := range response.Results {
		if branding.BrandingName == name || branding.BrandingNameSecondary == name {
			return &branding, nil
		}
	}

	// If no branding is found with the given name, return an error.
	return nil, fmt.Errorf("no self-service branding found with the name %s", name)
}

// CreateSelfServiceBrandingMacOS creates a new self-service branding configuration for macOS.
func (c *Client) CreateSelfServiceBrandingMacOS(branding *SelfServiceBrandingDetail) (*SelfServiceBrandingDetail, error) {
	endpoint := uriAPISelfServiceBrandingMacOS

	var response SelfServiceBrandingDetail
	resp, err := c.HTTP.DoRequest("POST", endpoint, branding, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create self-service branding: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateSelfServiceBrandingMacOSByID updates an existing self-service branding configuration for macOS.
func (c *Client) UpdateSelfServiceBrandingMacOSByID(id string, branding *SelfServiceBrandingDetail) (*SelfServiceBrandingDetail, error) {
	endpoint := fmt.Sprintf("%s/%s", uriAPISelfServiceBrandingMacOS, id)

	var response SelfServiceBrandingDetail
	resp, err := c.HTTP.DoRequest("PUT", endpoint, branding, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to update self-service branding: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// UpdateSelfServiceBrandingMacOSByName updates a self-service branding configuration for macOS by name.
func (c *Client) UpdateSelfServiceBrandingMacOSByName(name string, newBranding *SelfServiceBrandingDetail) (*SelfServiceBrandingDetail, error) {
	// First, get all branding configurations.
	response, err := c.GetSelfServiceBrandingMacOS()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch self-service branding for macOS: %v", err)
	}

	// Now, find the branding with the given name.
	var existingBranding *SelfServiceBrandingDetail
	for _, branding := range response.Results {
		if branding.BrandingName == name || branding.BrandingNameSecondary == name {
			existingBranding = &branding
			break
		}
	}

	if existingBranding == nil {
		return nil, fmt.Errorf("no self-service branding found with the name %s", name)
	}

	// Call the update by ID function with the found ID and the new branding details
	return c.UpdateSelfServiceBrandingMacOSByID(existingBranding.ID, newBranding)
}

// DeleteSelfServiceBrandingMacOSByID deletes a self-service branding configuration for macOS by ID.
func (c *Client) DeleteSelfServiceBrandingMacOSByID(id string) error {
	endpoint := fmt.Sprintf("%s/%s", uriAPISelfServiceBrandingMacOS, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete self-service branding: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
