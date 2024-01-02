// Refactor Complete

// classicapi_mobile_device_enrollment_profiles.go
// Jamf Pro Classic Api - Mobile Device Enrollment Profiles
// API reference: https://developer.jamf.com/jamf-pro/reference/mobiledeviceenrollmentprofiles
// Jamf Pro Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriMobileDeviceEnrollmentProfiles = "/JSSResource/mobiledeviceenrollmentprofiles"

// List

// ResponseMobileDeviceEnrollmentProfilesList represents the response for a list of mobile device enrollment profiles.
type ResponseMobileDeviceEnrollmentProfilesList struct {
	Size                          int                                      `xml:"size"`
	MobileDeviceEnrollmentProfile []MobileDeviceEnrollmentProfilesListItem `xml:"mobile_device_enrollment_profile"`
}

type MobileDeviceEnrollmentProfilesListItem struct {
	ID         int     `xml:"id"`
	Name       string  `xml:"name"`
	Invitation float64 `xml:"invitation"`
}

// Resource

// ResourceMobileDeviceEnrollmentProfile represents the response structure for a mobile device enrollment profile.
type ResourceMobileDeviceEnrollmentProfile struct {
	General     MobileDeviceEnrollmentProfileSubsetGeneral          `xml:"general"`
	Location    MobileDeviceEnrollmentProfileSubsetLocation         `xml:"location,omitempty"`
	Purchasing  MobileDeviceEnrollmentProfileSubsetPurchasing       `xml:"purchasing,omitempty"`
	Attachments []MobileDeviceEnrollmentProfileContainerAttachments `xml:"attachments,omitempty"`
}

// Subsets & Containers

type MobileDeviceEnrollmentProfileSubsetGeneral struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	Invitation  string `xml:"invitation,omitempty"`
	UDID        string `xml:"udid,omitempty"`
	Description string `xml:"description,omitempty"`
}

type MobileDeviceEnrollmentProfileSubsetLocation struct {
	Username     string `xml:"username,omitempty"`
	Realname     string `xml:"realname,omitempty"`
	RealName     string `xml:"real_name,omitempty"`
	EmailAddress string `xml:"email_address,omitempty"`
	Position     string `xml:"position,omitempty"`
	Phone        string `xml:"phone,omitempty"`
	PhoneNumber  string `xml:"phone_number,omitempty"`
	Department   string `xml:"department,omitempty"`
	Building     string `xml:"building,omitempty"`
	Room         int    `xml:"room,omitempty"`
}

type MobileDeviceEnrollmentProfileSubsetPurchasing struct {
	IsPurchased          bool   `xml:"is_purchased"`
	IsLeased             bool   `xml:"is_leased"`
	PONumber             string `xml:"po_number,omitempty"`
	Vendor               string `xml:"vendor,omitempty"`
	ApplecareID          string `xml:"applecare_id,omitempty"`
	PurchasePrice        string `xml:"purchase_price,omitempty"`
	PurchasingAccount    string `xml:"purchasing_account,omitempty"`
	PODate               string `xml:"po_date,omitempty"`
	PODateEpoch          int64  `xml:"po_date_epoch,omitempty"`
	PODateUTC            string `xml:"po_date_utc,omitempty"`
	WarrantyExpires      string `xml:"warranty_expires,omitempty"`
	WarrantyExpiresEpoch int64  `xml:"warranty_expires_epoch,omitempty"`
	WarrantyExpiresUTC   string `xml:"warranty_expires_utc,omitempty"`
	LeaseExpires         string `xml:"lease_expires,omitempty"`
	LeaseExpiresEpoch    int64  `xml:"lease_expires_epoch,omitempty"`
	LeaseExpiresUTC      string `xml:"lease_expires_utc,omitempty"`
	LifeExpectancy       int    `xml:"life_expectancy,omitempty"`
	PurchasingContact    string `xml:"purchasing_contact,omitempty"`
}

type MobileDeviceEnrollmentProfileContainerAttachments struct {
	Attachment MobileDeviceEnrollmentProfileSubsetAttachments `xml:"attachment"`
}

type MobileDeviceEnrollmentProfileSubsetAttachments struct {
	ID       int    `xml:"id"`
	Filename string `xml:"filename"`
	URI      string `xml:"uri"`
}

// CRUD

// GetMobileDeviceEnrollmentProfiles retrieves a serialized list of mobile device enrollment profiles.
func (c *Client) GetMobileDeviceEnrollmentProfiles() (*ResponseMobileDeviceEnrollmentProfilesList, error) {
	endpoint := uriMobileDeviceEnrollmentProfiles

	var enrollmentProfiles ResponseMobileDeviceEnrollmentProfilesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &enrollmentProfiles)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device enrollment profiles: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &enrollmentProfiles, nil
}

// GetMobileDeviceEnrollmentProfileByID fetches a specific mobile device enrollment profile by its ID.
func (c *Client) GetMobileDeviceEnrollmentProfileByID(id int) (*ResourceMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceEnrollmentProfiles, id)

	var profile ResourceMobileDeviceEnrollmentProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device enrollment profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceEnrollmentProfileByName fetches a specific mobile device enrollment profile by its name.
func (c *Client) GetMobileDeviceEnrollmentProfileByName(name string) (*ResourceMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceEnrollmentProfiles, name)

	var profile ResourceMobileDeviceEnrollmentProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device enrollment profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetProfileByInvitation fetches a specific mobile device enrollment profile by its invitation.
func (c *Client) GetProfileByInvitation(invitation string) (*ResourceMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/invitation/%s", uriMobileDeviceEnrollmentProfiles, invitation)

	var profile ResourceMobileDeviceEnrollmentProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device enrollment profile by invitation: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceEnrollmentProfileByIDBySubset fetches a specific mobile device configuration profile by its ID and a specified subset.
func (c *Client) GetMobileDeviceEnrollmentProfileByIDBySubset(id int, subset string) (*ResourceMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriMobileDeviceEnrollmentProfiles, id, subset)

	var profile ResourceMobileDeviceEnrollmentProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device configuration profile by ID and subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetMobileDeviceEnrollmentProfileByNameBySubset fetches a specific mobile device configuration profile by its name and a specified subset.
func (c *Client) GetMobileDeviceEnrollmentProfileByNameBySubset(name string, subset string) (*ResourceMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriMobileDeviceEnrollmentProfiles, name, subset)

	var profile ResourceMobileDeviceEnrollmentProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch mobile device configuration profile by name and subset: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// CreateMobileDeviceEnrollmentProfile creates a new mobile device enrollment profile on the Jamf Pro server.
func (c *Client) CreateMobileDeviceEnrollmentProfile(profile *ResourceMobileDeviceEnrollmentProfile) (*ResourceMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceEnrollmentProfiles)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_enrollment_profile"`
		*ResourceMobileDeviceEnrollmentProfile
	}{
		ResourceMobileDeviceEnrollmentProfile: profile,
	}

	var responseProfile ResourceMobileDeviceEnrollmentProfile
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to create mobile device enrollment profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// UpdateMobileDeviceEnrollmentProfileByID updates a mobile device enrollment profile by its ID.
func (c *Client) UpdateMobileDeviceEnrollmentProfileByID(id int, profile *ResourceMobileDeviceEnrollmentProfile) (*ResourceMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceEnrollmentProfiles, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_enrollment_profile"`
		*ResourceMobileDeviceEnrollmentProfile
	}{
		ResourceMobileDeviceEnrollmentProfile: profile,
	}

	var responseProfile ResourceMobileDeviceEnrollmentProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device enrollment profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// UpdateMobileDeviceEnrollmentProfileByName updates a mobile device enrollment profile by its name.
func (c *Client) UpdateMobileDeviceEnrollmentProfileByName(name string, profile *ResourceMobileDeviceEnrollmentProfile) (*ResourceMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceEnrollmentProfiles, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_enrollment_profile"`
		*ResourceMobileDeviceEnrollmentProfile
	}{
		ResourceMobileDeviceEnrollmentProfile: profile,
	}

	var responseProfile ResourceMobileDeviceEnrollmentProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device enrollment profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// UpdateMobileDeviceEnrollmentProfileByInvitation updates a mobile device enrollment profile by its invitation.
func (c *Client) UpdateMobileDeviceEnrollmentProfileByInvitation(invitation string, profile *ResourceMobileDeviceEnrollmentProfile) (*ResourceMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/invitation/%s", uriMobileDeviceEnrollmentProfiles, invitation)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_enrollment_profile"`
		*ResourceMobileDeviceEnrollmentProfile
	}{
		ResourceMobileDeviceEnrollmentProfile: profile,
	}

	var responseProfile ResourceMobileDeviceEnrollmentProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &responseProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update mobile device enrollment profile by invitation: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &responseProfile, nil
}

// DeleteMobileDeviceEnrollmentProfileByID deletes a mobile device enrollment profile by its ID.
func (c *Client) DeleteMobileDeviceEnrollmentProfileByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceEnrollmentProfiles, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device enrollment profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceEnrollmentProfileByName deletes a mobile device enrollment profile by its name.
func (c *Client) DeleteMobileDeviceEnrollmentProfileByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceEnrollmentProfiles, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device enrollment profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteMobileDeviceEnrollmentProfileByInvitation deletes a mobile device enrollment profile by its invitation.
func (c *Client) DeleteMobileDeviceEnrollmentProfileByInvitation(invitation string) error {
	endpoint := fmt.Sprintf("%s/invitation/%s", uriMobileDeviceEnrollmentProfiles, invitation)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete mobile device enrollment profile by invitation: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
