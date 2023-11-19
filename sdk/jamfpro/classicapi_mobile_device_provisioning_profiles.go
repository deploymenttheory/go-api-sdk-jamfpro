// classicapi_mobile_device_provisioning_profiles.go
// Jamf Pro Classic Api - Mobile Device Provisioning Profiles
// API reference: https://developer.jamf.com/jamf-pro/reference/mobiledeviceprovisioningprofiles
// Jamf Pro Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMobileDeviceProvisioningProfiles = "/JSSResource/mobiledeviceprovisioningprofiles"

// ResponseMobileDeviceProvisioningProfilesList represents the response for a list of mobile device provisioning profiles.
type ResponseMobileDeviceProvisioningProfilesList struct {
	Size                            int                                   `xml:"size"`
	MobileDeviceProvisioningProfile []MobileDeviceProvisioningProfileItem `xml:"mobile_device_provisioning_profile"`
}

// MobileDeviceProvisioningProfileItem represents a single mobile device provisioning profile item.
type MobileDeviceProvisioningProfileItem struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	DisplayName string `xml:"display_name"`
	UUID        string `xml:"uuid"`
}

// ResponseMobileDeviceProvisioningProfile represents the detailed structure for a mobile device provisioning profile.
type ResponseMobileDeviceProvisioningProfile struct {
	General MobileDeviceProvisioningProfileGeneral `xml:"general"`
}

// MobileDeviceProvisioningProfileGeneral contains general information about the provisioning profile.
type MobileDeviceProvisioningProfileGeneral struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	DisplayName string `xml:"display_name"`
	UUID        string `xml:"uuid"`
}

// GetMobileDeviceProvisioningProfiles retrieves a serialized list of mobile device provisioning profiles.
func (c *Client) GetMobileDeviceProvisioningProfiles() (*ResponseMobileDeviceProvisioningProfilesList, error) {
	endpoint := uriMobileDeviceProvisioningProfiles

	var profiles ResponseMobileDeviceProvisioningProfilesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profiles)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device provisioning profiles: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profiles, nil
}

// GetMobileDeviceProvisioningProfileByID fetches a specific mobile device provisioning profile by its ID.
func (c *Client) GetMobileDeviceProvisioningProfileByID(id int) (*ResponseMobileDeviceProvisioningProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceProvisioningProfiles, id)

	var profile ResponseMobileDeviceProvisioningProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device provisioning profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceProvisioningProfileByName fetches a specific mobile device provisioning profile by its name.
func (c *Client) GetMobileDeviceProvisioningProfileByName(name string) (*ResponseMobileDeviceProvisioningProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceProvisioningProfiles, name)

	var profile ResponseMobileDeviceProvisioningProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device provisioning profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceProvisioningProfileByUUID fetches a specific mobile device provisioning profile by its UUID.
func (c *Client) GetMobileDeviceProvisioningProfileByUUID(uuid string) (*ResponseMobileDeviceProvisioningProfile, error) {
	endpoint := fmt.Sprintf("%s/uuid/%s", uriMobileDeviceProvisioningProfiles, uuid)

	var profile ResponseMobileDeviceProvisioningProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device provisioning profile by UUID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// CreateMobileDeviceProvisioningProfileByID creates a new mobile device provisioning profile by its ID.
func (c *Client) CreateMobileDeviceProvisioningProfileByID(id int, profile *ResponseMobileDeviceProvisioningProfile) (*ResponseMobileDeviceProvisioningProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceProvisioningProfiles, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_provisioning_profile"`
		*ResponseMobileDeviceProvisioningProfile
	}{
		ResponseMobileDeviceProvisioningProfile: profile,
	}

	var responseProfile ResponseMobileDeviceProvisioningProfile
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to create mobile device provisioning profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// CreateMobileDeviceProvisioningProfileByName creates a new mobile device provisioning profile by its name.
func (c *Client) CreateMobileDeviceProvisioningProfileByName(name string, profile *ResponseMobileDeviceProvisioningProfile) (*ResponseMobileDeviceProvisioningProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceProvisioningProfiles, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_provisioning_profile"`
		*ResponseMobileDeviceProvisioningProfile
	}{
		ResponseMobileDeviceProvisioningProfile: profile,
	}

	var responseProfile ResponseMobileDeviceProvisioningProfile
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to create mobile device provisioning profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// CreateMobileDeviceProvisioningProfileByUUID creates a new mobile device provisioning profile by its UUID.
func (c *Client) CreateMobileDeviceProvisioningProfileByUUID(uuid string, profile *ResponseMobileDeviceProvisioningProfile) (*ResponseMobileDeviceProvisioningProfile, error) {
	endpoint := fmt.Sprintf("%s/uuid/%s", uriMobileDeviceProvisioningProfiles, uuid)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_provisioning_profile"`
		*ResponseMobileDeviceProvisioningProfile
	}{
		ResponseMobileDeviceProvisioningProfile: profile,
	}

	var responseProfile ResponseMobileDeviceProvisioningProfile
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to create mobile device provisioning profile by UUID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// UpdateMobileDeviceProvisioningProfileByID updates a mobile device provisioning profile by its ID.
func (c *Client) UpdateMobileDeviceProvisioningProfileByID(id int, profile *ResponseMobileDeviceProvisioningProfile) (*ResponseMobileDeviceProvisioningProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceProvisioningProfiles, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_provisioning_profile"`
		*ResponseMobileDeviceProvisioningProfile
	}{
		ResponseMobileDeviceProvisioningProfile: profile,
	}

	var updatedProfile ResponseMobileDeviceProvisioningProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device provisioning profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

// UpdateMobileDeviceProvisioningProfileByName updates a mobile device provisioning profile by its name.
func (c *Client) UpdateMobileDeviceProvisioningProfileByName(name string, profile *ResponseMobileDeviceProvisioningProfile) (*ResponseMobileDeviceProvisioningProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceProvisioningProfiles, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_provisioning_profile"`
		*ResponseMobileDeviceProvisioningProfile
	}{
		ResponseMobileDeviceProvisioningProfile: profile,
	}

	var updatedProfile ResponseMobileDeviceProvisioningProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device provisioning profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

// UpdateMobileDeviceProvisioningProfileByUUID updates a mobile device provisioning profile by its UUID.
func (c *Client) UpdateMobileDeviceProvisioningProfileByUUID(uuid string, profile *ResponseMobileDeviceProvisioningProfile) (*ResponseMobileDeviceProvisioningProfile, error) {
	endpoint := fmt.Sprintf("%s/uuid/%s", uriMobileDeviceProvisioningProfiles, uuid)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_provisioning_profile"`
		*ResponseMobileDeviceProvisioningProfile
	}{
		ResponseMobileDeviceProvisioningProfile: profile,
	}

	var updatedProfile ResponseMobileDeviceProvisioningProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device provisioning profile by UUID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

// DeleteMobileDeviceProvisioningProfileByID deletes a mobile device provisioning profile by ID
func (c *Client) DeleteMobileDeviceProvisioningProfileByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceProvisioningProfiles, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device provisioning profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceProvisioningProfileByName deletes a mobile device provisioning profile by Name
func (c *Client) DeleteMobileDeviceProvisioningProfileByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceProvisioningProfiles, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device provisioning profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceProvisioningProfileByUUID deletes a mobile device provisioning profile by UUID
func (c *Client) DeleteMobileDeviceProvisioningProfileByUUID(uuid string) error {
	endpoint := fmt.Sprintf("%s/uuid/%s", uriMobileDeviceProvisioningProfiles, uuid)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device provisioning profile by UUID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
