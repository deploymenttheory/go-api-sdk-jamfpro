// classicapi_byoprofiles.go
// Jamf Pro Classic Api - Personal Device Profiles
// api reference: https://developer.jamf.com/jamf-pro/reference/byoprofiles
// Classic API requires the structs to support an XML data structure.

/*
Shared Resources in this Endpoint:
- SharedResourceSite
*/

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriBYOProfiles = "/JSSResource/byoprofiles"

// List

// ResponseBYOProfilesList represents the XML response for a list of BYO profiles.
type ResponseBYOProfilesList struct {
	Size        int                  `xml:"size"`
	BYOProfiles []BYOProfileListItem `xml:"byoprofile"`
}

type BYOProfileListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// Resource

// BYOProfile represents the details of a BYO profile.
type ResourceBYOProfile struct {
	General BYOProfileSubsetGeneral `xml:"general"`
}

// Subsets

type BYOProfileSubsetGeneral struct {
	ID          int                `xml:"id"`
	Name        string             `xml:"name"`
	Site        SharedResourceSite `xml:"site"`
	Enabled     bool               `xml:"enabled"`
	Description string             `xml:"description"`
}

// Responses

type ResponceBYOProfileCreatedAndUpdated struct {
	ID int `json:"id,omitempty" xml:"id,omitempty"`
}

// GetBYOProfiles gets a list of all BYO profiles.
func (c *Client) GetBYOProfiles() (*ResponseBYOProfilesList, error) {
	endpoint := uriBYOProfiles

	var byoProfiles ResponseBYOProfilesList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &byoProfiles)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all BYO Profiles: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &byoProfiles, nil
}

// GetBYOProfileByID retrieves a BYO profile by its ID.
func (c *Client) GetBYOProfileByID(id int) (*ResourceBYOProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriBYOProfiles, id)

	var profile ResourceBYOProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGetByID, "byo profile", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetBYOProfileByName retrieves a BYO profile by its name.
func (c *Client) GetBYOProfileByName(name string) (*ResourceBYOProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriBYOProfiles, name)

	var profile ResourceBYOProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch BYO Profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// CreateBYOProfile creates a new BYO profile.
func (c *Client) CreateBYOProfile(profile *ResourceBYOProfile) (*ResponceBYOProfileCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriBYOProfiles)

	requestBody := struct {
		XMLName xml.Name `xml:"byoprofile"`
		*ResourceBYOProfile
	}{
		ResourceBYOProfile: profile,
	}

	var createdProfile ResponceBYOProfileCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdProfile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "byo profile", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdProfile, nil
}

// UpdateBYOProfileByID updates an existing BYO profile by its ID.
func (c *Client) UpdateBYOProfileByID(id int, profile *ResourceBYOProfile) (*ResponceBYOProfileCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriBYOProfiles, id)

	requestBody := struct {
		XMLName xml.Name `xml:"byoprofile"`
		*ResourceBYOProfile
	}{
		ResourceBYOProfile: profile,
	}

	var updatedProfile ResponceBYOProfileCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByID, "byo profile", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

// UpdateBYOProfileByName updates a BYO profile by its name.
func (c *Client) UpdateBYOProfileByName(name string, profile *ResourceBYOProfile) (*ResponceBYOProfileCreatedAndUpdated, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriBYOProfiles, name)

	requestBody := struct {
		XMLName xml.Name `xml:"byoprofile"`
		*ResourceBYOProfile
	}{
		ResourceBYOProfile: profile,
	}

	var updatedProfile ResponceBYOProfileCreatedAndUpdated
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedUpdateByName, "byo profile", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

// DeleteBYOProfileByID deletes a BYO profile by its ID.
func (c *Client) DeleteBYOProfileByID(id string) error {
	endpoint := fmt.Sprintf("%s/id/%s", uriBYOProfiles, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByID, "byo profile", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}

// DeleteBYOProfileByName deletes a BYO profile by its name.
func (c *Client) DeleteBYOProfileByName(name string) error {
	endpoint := fmt.Sprintf("%s/name/%s", uriBYOProfiles, name)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf(errMsgFailedDeleteByName, "byo profile", name, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
