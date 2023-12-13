// jamfproapi_computer_inventory_collection_settings.go
// Jamf Pro Api - Computer Inventory Collection Settings
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v1-computer-inventory-collection-settings
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"encoding/json"
	"fmt"
)

const uriComputerInventoryCollectionSettings = "/api/v1/computer-inventory-collection-settings"

type ResourceComputerInventoryCollectionSettings struct {
	ComputerInventoryCollectionPreferences ResourceDataInventoryCollectionPreference `json:"computerInventoryCollectionPreferences"`
	ApplicationPaths                       []ResourceDataPathItem                    `json:"applicationPaths"`
	FontPaths                              []ResourceDataPathItem                    `json:"fontPaths"`
	PluginPaths                            []ResourceDataPathItem                    `json:"pluginPaths"`
}
type ResourceDataInventoryCollectionPreference struct {
	MonitorApplicationUsage                      bool `json:"monitorApplicationUsage"`
	IncludeFonts                                 bool `json:"includeFonts"`
	IncludePlugins                               bool `json:"includePlugins"`
	IncludePackages                              bool `json:"includePackages"`
	IncludeSoftwareUpdates                       bool `json:"includeSoftwareUpdates"`
	IncludeSoftwareId                            bool `json:"includeSoftwareId"`
	IncludeAccounts                              bool `json:"includeAccounts"`
	CalculateSizes                               bool `json:"calculateSizes"`
	IncludeHiddenAccounts                        bool `json:"includeHiddenAccounts"`
	IncludePrinters                              bool `json:"includePrinters"`
	IncludeServices                              bool `json:"includeServices"`
	CollectSyncedMobileDeviceInfo                bool `json:"collectSyncedMobileDeviceInfo"`
	UpdateLdapInfoOnComputerInventorySubmissions bool `json:"updateLdapInfoOnComputerInventorySubmissions"`
	MonitorBeacons                               bool `json:"monitorBeacons"`
	AllowChangingUserAndLocation                 bool `json:"allowChangingUserAndLocation"`
	UseUnixUserPaths                             bool `json:"useUnixUserPaths"`
	CollectUnmanagedCertificates                 bool `json:"collectUnmanagedCertificates"`
}

type ResourceDataPathItem struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

// ComputerInventoryCollectionSettingsCustomPath defines the request body for creating a custom path.
type ResourceComputerInventoryCollectionSettingsCustomPath struct {
	Scope string `json:"scope"`
	Path  string `json:"path"`
}

func (c *Client) GetComputerInventoryCollectionSettings() (*ResourceComputerInventoryCollectionSettings, error) {
	endpoint := uriComputerInventoryCollectionSettings

	var settings ResourceComputerInventoryCollectionSettings
	resp, err := c.HTTP.DoRequest("GET", endpoint, nil, &settings)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedGet, "computer inventory collection settings", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &settings, nil
}

// UpdateComputerInventoryCollectionSettings updates the computer inventory collection settings.
func (c *Client) UpdateComputerInventoryCollectionSettings(settings *ResourceComputerInventoryCollectionSettings) (*ResourceComputerInventoryCollectionSettings, error) {
	endpoint := uriComputerInventoryCollectionSettings

	// Marshal the settings into JSON for the request body
	requestBody, err := json.Marshal(settings)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal settings: %v", err)
	}

	// Perform the PATCH request
	resp, err := c.HTTP.DoRequest("PATCH", endpoint, requestBody, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to update computer inventory collection settings: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Read the response body into the same settings struct
	if err := json.NewDecoder(resp.Body).Decode(&settings); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %v", err)
	}

	return settings, nil
}

// CreateComputerInventoryCollectionSettingsCustomPath creates a custom path for computer inventory collection settings.
func (c *Client) CreateComputerInventoryCollectionSettingsCustomPath(customPath *ResourceComputerInventoryCollectionSettingsCustomPath) (*ResourceComputerInventoryCollectionSettingsCustomPath, error) {
	endpoint := fmt.Sprintf("%s/custom-path", uriComputerInventoryCollectionSettings)

	var response ResourceComputerInventoryCollectionSettingsCustomPath
	resp, err := c.HTTP.DoRequest("POST", endpoint, customPath, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create computer inventory collection settings custom path: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	return &response, nil
}

// DeleteComputerInventoryCollectionSettingsCustomPathByID deletes a custom path by ID.
func (c *Client) DeleteComputerInventoryCollectionSettingsCustomPathByID(id string) error {
	endpoint := fmt.Sprintf("%s/custom-path/%s", uriComputerInventoryCollectionSettings, id)

	resp, err := c.HTTP.DoRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete computer inventory collection settings custom path: %v", err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}

	// Success, no error
	return nil
}
