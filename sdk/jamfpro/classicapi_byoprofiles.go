// classicapi_byoprofiles.go
// Jamf Pro Classic Api - Personal Device Profiles
// api reference: https://developer.jamf.com/jamf-pro/reference/byoprofiles
// Classic API requires the structs to support an XML data structure.

package jamfpro

import (
	"encoding/xml"
	"fmt"
)

const uriBYOProfiles = "/JSSResource/byoprofiles"

// ResponseBYOProfilesList represents the XML response for a list of BYO profiles.
type ResponseBYOProfilesList struct {
	Size        int              `xml:"size"`
	BYOProfiles []BYOProfileItem `xml:"byoprofile"`
}

// BYOProfileItem represents a single BYO profile item in the list.
type BYOProfileItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

// BYOProfile represents the details of a BYO profile.
type ResponseBYOProfile struct {
	General BYOProfileGeneralInfo `xml:"general"`
}

// GeneralInfo represents the general section of a BYO profile.
type BYOProfileGeneralInfo struct {
	ID          int                `xml:"id"`
	Name        string             `xml:"name"`
	Site        BYOProfileSiteInfo `xml:"site"`
	Enabled     bool               `xml:"enabled"`
	Description string             `xml:"description"`
}

// SiteInfo represents the site information of a BYO profile.
type BYOProfileSiteInfo struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
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
func (c *Client) GetBYOProfileByID(id int) (*ResponseBYOProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriBYOProfiles, id)

	var profile ResponseBYOProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &profile)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch BYO Profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &profile, nil
}

// GetBYOProfileByName retrieves a BYO profile by its name.
func (c *Client) GetBYOProfileByName(name string) (*ResponseBYOProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriBYOProfiles, name)

	var profile ResponseBYOProfile
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
func (c *Client) CreateBYOProfile(profile *ResponseBYOProfile) (*ResponseBYOProfile, error) {
	endpoint := fmt.Sprintf("%s/id/0", uriBYOProfiles)

	// Wrap the profile request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"byoprofile"`
		*ResponseBYOProfile
	}{
		ResponseBYOProfile: profile,
	}

	var createdProfile ResponseBYOProfile
	resp, err := c.HTTP.DoRequest("POST", endpoint, &requestBody, &createdProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to create BYO Profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &createdProfile, nil
}

// UpdateBYOProfileByID updates an existing BYO profile by its ID.
func (c *Client) UpdateBYOProfileByID(id int, profile *ResponseBYOProfile) (*ResponseBYOProfile, error) {
	endpoint := fmt.Sprintf("%s/id/%d", uriBYOProfiles, id)

	// Wrap the profile request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"byoprofile"`
		*ResponseBYOProfile
	}{
		ResponseBYOProfile: profile,
	}

	var updatedProfile ResponseBYOProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update BYO Profile by ID: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

// UpdateBYOProfileByName updates a BYO profile by its name.
func (c *Client) UpdateBYOProfileByName(name string, profile *ResponseBYOProfile) (*ResponseBYOProfile, error) {
	endpoint := fmt.Sprintf("%s/name/%s", uriBYOProfiles, name)

	// Wrap the profile request with the desired XML name using an anonymous struct
	requestBody := struct {
		XMLName xml.Name `xml:"byoprofile"`
		*ResponseBYOProfile
	}{
		ResponseBYOProfile: profile,
	}

	var updatedProfile ResponseBYOProfile
	resp, err := c.HTTP.DoRequest("PUT", endpoint, &requestBody, &updatedProfile)
	if err != nil {
		return nil, fmt.Errorf("failed to update BYO Profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &updatedProfile, nil
}

// DeleteBYOProfileByID deletes a BYO profile by its ID.
func (c *Client) DeleteBYOProfileByID(id int) error {
	endpoint := fmt.Sprintf("%s/id/%d", uriBYOProfiles, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete BYO Profile by ID: %v", err)
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
		return fmt.Errorf("failed to delete BYO Profile by name: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return nil
}
