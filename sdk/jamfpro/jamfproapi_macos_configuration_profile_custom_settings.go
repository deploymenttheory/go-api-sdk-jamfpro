// jamfproapi_macos_configuration_profile_custom_settings.go
// Jamf Pro Api - Config profile Custom Settings Schema
// api reference: Undocumented
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriConfigProfiles = "/api/config-profiles/macos"

// ResponseCustomSettingsSchemaList represents the top-level response for custom settings schema list
type ResponseCustomSettingsSchemaList []ResourceCustomSettingsBucket

// ResourceCustomSettingsBucket represents a single bucket of custom settings schemas
type ResourceCustomSettingsBucket struct {
	BucketName            string                                  `json:"bucketName"`
	DisplayName           string                                  `json:"displayName"`
	CustomSettingsDomains map[string]ResourceCustomSettingsDomain `json:"customSettingsDomains"`
}

// ResourceCustomSettingsDomain represents a domain of custom settings
type ResourceCustomSettingsDomain struct {
	SettingsDomain string                           `json:"settingsDomain"`
	Versions       map[string]ResourceDomainVersion `json:"versions"`
}

// ResourceDomainVersion represents a specific version of a domain
type ResourceDomainVersion struct {
	Version  string   `json:"version"`
	Variants []string `json:"variants"`
}

//--------------

// ResourceConfigProfile represents a macOS configuration profile
type ResourceConfigProfile struct {
	PayloadUUID    string               `json:"payloadUUID"`
	PayloadContent []PayloadContentItem `json:"payloadContent"`
	Level          string               `json:"level,omitempty"`
}

// PayloadContentItem represents an item in the payload content
type PayloadContentItem struct {
	PayloadType         string          `json:"payloadType"`
	PayloadVersion      int             `json:"payloadVersion"`
	PayloadIdentifier   string          `json:"payloadIdentifier"`
	PayloadUUID         string          `json:"payloadUUID"`
	PayloadOrganization string          `json:"payloadOrganization,omitempty"`
	PreferenceDomain    string          `json:"preferenceDomain,omitempty"`
	Forced              *ForcedSettings `json:"forced,omitempty"`
	PayloadDisplayName  string          `json:"payloadDisplayName,omitempty"`
}

// ForcedSettings represents forced settings in a payload
type ForcedSettings struct {
	Plist         string `json:"plist,omitempty"`
	JsonSchema    string `json:"jsonSchema,omitempty"`
	SchemaSource  string `json:"schemaSource,omitempty"`
	SchemaDomain  string `json:"schemaDomain,omitempty"`
	SchemaVersion string `json:"schemaVersion,omitempty"`
	SchemaVariant string `json:"schemaVariant,omitempty"`
}

// Response

// ResponseConfigProfileCreate represents the response when creating a configuration profile
type ResponseConfigProfileCreate struct {
	UUID string `json:"uuid"`
}

// GetCustomSettingsSchemaList retrieves the list of custom settings schemas
func (c *Client) GetCustomSettingsSchemaList() (*ResponseCustomSettingsSchemaList, error) {
	endpoint := fmt.Sprintf("%s/custom-settings/v1/schema-list", uriConfigProfiles)
	var out ResponseCustomSettingsSchemaList
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to get custom settings schema list: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// GetConfigProfileByPayloadUUID retrieves a macOS configuration profile by payload UUID
func (c *Client) GetConfigProfileByPayloadUUID(id string) (*ResourceConfigProfile, error) {
	endpoint := fmt.Sprintf("%s/%s", uriConfigProfiles, id)
	var out ResourceConfigProfile
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to get macOS configuration profile with ID %s: %v", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}

// CreateConfigProfileWithCustomSettingsSchema creates a new macOS configuration profile
func (c *Client) CreateConfigProfileWithCustomSettingsSchema(profile *ResourceConfigProfile) (*ResponseConfigProfileCreate, error) {
	endpoint := uriConfigProfiles
	var out ResponseConfigProfileCreate
	resp, err := c.HTTP.DoRequest("POST", endpoint, profile, &out)

	if err != nil {
		return nil, fmt.Errorf("failed to create macOS configuration profile: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &out, nil
}
