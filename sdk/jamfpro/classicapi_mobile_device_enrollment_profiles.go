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

// ResponseMobileDeviceEnrollmentProfilesList represents the response for a list of mobile device enrollment profiles.
type ResponseMobileDeviceEnrollmentProfilesList struct {
	Size                          int                                 `xml:"size"`
	MobileDeviceEnrollmentProfile []MobileDeviceEnrollmentProfileItem `xml:"mobile_device_enrollment_profile"`
}

// MobileDeviceEnrollmentProfileItem represents a single mobile device enrollment profile item.
type MobileDeviceEnrollmentProfileItem struct {
	ID         int     `xml:"id"`
	Name       string  `xml:"name"`
	Invitation float64 `xml:"invitation"`
}

// ResponseMobileDeviceEnrollmentProfile represents the response structure for a mobile device enrollment profile.
type ResponseMobileDeviceEnrollmentProfile struct {
	General     MobileDeviceEnrollmentProfileGeneral      `xml:"general"`
	Location    MobileDeviceEnrollmentProfileLocation     `xml:"location,omitempty"`
	Purchasing  MobileDeviceEnrollmentProfilePurchasing   `xml:"purchasing,omitempty"`
	Attachments []MobileDeviceEnrollmentProfileAttachment `xml:"attachments,omitempty"`
}

// MobileDeviceEnrollmentProfileGeneral contains general information about the enrollment profile.
type MobileDeviceEnrollmentProfileGeneral struct {
	ID          int    `xml:"id"`
	Name        string `xml:"name"`
	Invitation  string `xml:"invitation,omitempty"`
	UDID        string `xml:"udid,omitempty"`
	Description string `xml:"description,omitempty"`
}

// MobileDeviceEnrollmentProfileLocation contains location information of the enrollment profile.
type MobileDeviceEnrollmentProfileLocation struct {
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

// MobileDeviceEnrollmentProfilePurchasing contains purchasing information of the enrollment profile.
type MobileDeviceEnrollmentProfilePurchasing struct {
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

// MobileDeviceEnrollmentProfileAttachment represents an attachment in the enrollment profile.
type MobileDeviceEnrollmentProfileAttachment struct {
	Attachment MobileDeviceEnrollmentProfileAttachmentItem `xml:"attachment"`
}

// MobileDeviceEnrollmentProfileAttachmentItem contains details of an attachment.
type MobileDeviceEnrollmentProfileAttachmentItem struct {
	ID       int    `xml:"id"`
	Filename string `xml:"filename"`
	URI      string `xml:"uri"`
}

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
func (c *Client) GetMobileDeviceEnrollmentProfileByID(id int) (*ResponseMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceEnrollmentProfiles, id)

	var profile ResponseMobileDeviceEnrollmentProfile
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
func (c *Client) GetMobileDeviceEnrollmentProfileByName(name string) (*ResponseMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceEnrollmentProfiles, name)

	var profile ResponseMobileDeviceEnrollmentProfile
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
func (c *Client) GetProfileByInvitation(invitation string) (*ResponseMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/invitation/%s", uriMobileDeviceEnrollmentProfiles, invitation)

	var profile ResponseMobileDeviceEnrollmentProfile
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
func (c *Client) GetMobileDeviceEnrollmentProfileByIDBySubset(id int, subset string) (*ResponseMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d/subset/%s", uriMobileDeviceEnrollmentProfiles, id, subset)

	var profile ResponseMobileDeviceEnrollmentProfile
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
func (c *Client) GetMobileDeviceEnrollmentProfileByNameBySubset(name string, subset string) (*ResponseMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s/subset/%s", uriMobileDeviceEnrollmentProfiles, name, subset)

	var profile ResponseMobileDeviceEnrollmentProfile
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
func (c *Client) CreateMobileDeviceEnrollmentProfile(profile *ResponseMobileDeviceEnrollmentProfile) (*ResponseMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriMobileDeviceEnrollmentProfiles)

	// Wrap the profile with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_enrollment_profile"`
		*ResponseMobileDeviceEnrollmentProfile
	}{
		ResponseMobileDeviceEnrollmentProfile: profile,
	}

	var responseProfile ResponseMobileDeviceEnrollmentProfile
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
func (c *Client) UpdateMobileDeviceEnrollmentProfileByID(id int, profile *ResponseMobileDeviceEnrollmentProfile) (*ResponseMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriMobileDeviceEnrollmentProfiles, id)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_enrollment_profile"`
		*ResponseMobileDeviceEnrollmentProfile
	}{
		ResponseMobileDeviceEnrollmentProfile: profile,
	}

	var responseProfile ResponseMobileDeviceEnrollmentProfile
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
func (c *Client) UpdateMobileDeviceEnrollmentProfileByName(name string, profile *ResponseMobileDeviceEnrollmentProfile) (*ResponseMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriMobileDeviceEnrollmentProfiles, name)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_enrollment_profile"`
		*ResponseMobileDeviceEnrollmentProfile
	}{
		ResponseMobileDeviceEnrollmentProfile: profile,
	}

	var responseProfile ResponseMobileDeviceEnrollmentProfile
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
func (c *Client) UpdateMobileDeviceEnrollmentProfileByInvitation(invitation string, profile *ResponseMobileDeviceEnrollmentProfile) (*ResponseMobileDeviceEnrollmentProfile, error) {
	endpoint := fmt.Sprintf("%s/invitation/%s", uriMobileDeviceEnrollmentProfiles, invitation)

	requestBody := struct {
		XMLName xml.Name `xml:"mobile_device_enrollment_profile"`
		*ResponseMobileDeviceEnrollmentProfile
	}{
		ResponseMobileDeviceEnrollmentProfile: profile,
	}

	var responseProfile ResponseMobileDeviceEnrollmentProfile
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
