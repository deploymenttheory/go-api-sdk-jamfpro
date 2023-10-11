// managedPreferenceProfiles.go
// Jamf Pro Classic Api
// Classic API requires the structs to support both XML and JSON.

package jamfpro

import (
	"fmt"
)

const uriManagedPreferenceProfiles = "/JSSResource/managedpreferenceprofiles"

type ResponseManagedPreferenceProfile struct {
	General GeneralInfo `json:"general" xml:"general"`
}

type GeneralInfo struct {
	ID      int    `json:"id" xml:"id"`
	Name    string `json:"name" xml:"name"`
	Enabled bool   `json:"enabled" xml:"enabled"`
	Plist   string `json:"plist" xml:"plist"`
}

type ManagedPreferenceProfileList struct {
	Size    int                                `json:"size" xml:"size"`
	Results []ResponseManagedPreferenceProfile `json:"managedpreferenceprofile" xml:"managedpreferenceprofile"`
}

// GetManagedPreferenceProfileByID retrieves the Managed Preference Profile by its ID
func (c *Client) GetManagedPreferenceProfileByID(id int) (*ResponseManagedPreferenceProfile, error) {
	url := fmt.Sprintf("%s/id/%d", uriManagedPreferenceProfiles, id)

	var profile ResponseManagedPreferenceProfile
	if err := c.DoRequest("GET", url, nil, nil, &profile); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &profile, nil
}

// GetManagedPreferenceProfileByName retrieves the Managed Preference Profile by its Name
func (c *Client) GetManagedPreferenceProfileByName(name string) (*ResponseManagedPreferenceProfile, error) {
	url := fmt.Sprintf("%s/name/%s", uriManagedPreferenceProfiles, name)

	var profile ResponseManagedPreferenceProfile
	if err := c.DoRequest("GET", url, nil, nil, &profile); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &profile, nil
}

// GetManagedPreferenceProfileByNameWithSubset retrieves the Managed Preference Profile by its Name with data subset
func (c *Client) GetManagedPreferenceProfileByNameWithSubset(name, subset string) (*ResponseManagedPreferenceProfile, error) {
	url := fmt.Sprintf("%s/name/%s/subset/%s", uriManagedPreferenceProfiles, name, subset)

	var profile ResponseManagedPreferenceProfile
	if err := c.DoRequest("GET", url, nil, nil, &profile); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &profile, nil
}

// GetManagedPreferenceProfileByIDWithSubset retrieves the Managed Preference Profile by its ID with data subset
func (c *Client) GetManagedPreferenceProfileByIDWithSubset(id int, subset string) (*ResponseManagedPreferenceProfile, error) {
	url := fmt.Sprintf("%s/id/%d/subset/%s", uriManagedPreferenceProfiles, id, subset)

	var profile ResponseManagedPreferenceProfile
	if err := c.DoRequest("GET", url, nil, nil, &profile); err != nil {
		return nil, fmt.Errorf("failed to execute request: %v", err)
	}

	return &profile, nil
}

// CreateManagedPreferenceProfile creates a new Managed Preference Profile
func (c *Client) CreateManagedPreferenceProfile(profile *ResponseManagedPreferenceProfile) (*ResponseManagedPreferenceProfile, error) {
	url := fmt.Sprintf("%s/id/0", uriManagedPreferenceProfiles)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"managed_preference_profile"`
		*ResponseManagedPreferenceProfile
	}{
		ResponseManagedPreferenceProfile: profile,
	}

	// Execute the request
	var responseProfile ResponseManagedPreferenceProfile
	if err := c.DoRequest("POST", url, reqBody, nil, &responseProfile); err != nil {
		return nil, fmt.Errorf("failed to create Managed Preference Profile: %v", err)
	}

	return &responseProfile, nil
}

// UpdateManagedPreferenceProfileById updates an existing Managed Preference Profile by its ID
func (c *Client) UpdateManagedPreferenceProfileById(id int, profile *ResponseManagedPreferenceProfile) (*ResponseManagedPreferenceProfile, error) {
	url := fmt.Sprintf("%s/id/%d", uriManagedPreferenceProfiles, id)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"managed_preference_profile"`
		*ResponseManagedPreferenceProfile
	}{
		ResponseManagedPreferenceProfile: profile,
	}

	// Execute the request
	var responseProfile ResponseManagedPreferenceProfile
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseProfile); err != nil {
		return nil, fmt.Errorf("failed to update Managed Preference Profile: %v", err)
	}

	return &responseProfile, nil
}

// UpdateManagedPreferenceProfileByName updates an existing Managed Preference Profile by its name
func (c *Client) UpdateManagedPreferenceProfileByName(name string, profile *ResponseManagedPreferenceProfile) (*ResponseManagedPreferenceProfile, error) {
	url := fmt.Sprintf("%s/name/%s", uriManagedPreferenceProfiles, name)

	// Construct a custom request body structure for proper XML serialization
	reqBody := &struct {
		XMLName struct{} `xml:"managed_preference_profile"`
		*ResponseManagedPreferenceProfile
	}{
		ResponseManagedPreferenceProfile: profile,
	}

	// Execute the request
	var responseProfile ResponseManagedPreferenceProfile
	if err := c.DoRequest("PUT", url, reqBody, nil, &responseProfile); err != nil {
		return nil, fmt.Errorf("failed to update Managed Preference Profile by name: %v", err)
	}

	return &responseProfile, nil
}

// DeleteManagedPreferenceProfileById deletes an existing Managed Preference Profile by its ID
func (c *Client) DeleteManagedPreferenceProfileById(id int) error {
	url := fmt.Sprintf("%s/id/%d", uriManagedPreferenceProfiles, id)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Managed Preference Profile: %v", err)
	}

	return nil
}

// DeleteManagedPreferenceProfileByName deletes an existing Managed Preference Profile by its name
func (c *Client) DeleteManagedPreferenceProfileByName(name string) error {
	url := fmt.Sprintf("%s/name/%s", uriManagedPreferenceProfiles, name)

	if err := c.DoRequest("DELETE", url, nil, nil, nil); err != nil {
		return fmt.Errorf("failed to delete Managed Preference Profile by name: %v", err)
	}

	return nil
}
