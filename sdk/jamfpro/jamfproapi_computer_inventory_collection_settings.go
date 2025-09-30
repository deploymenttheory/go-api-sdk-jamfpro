// jamfproapi_computer_inventory_collection_settings.go
// Jamf Pro Api - Computer Inventory Collection Settings
// api reference: https://developer.jamf.com/jamf-pro/reference/get_v2-computer-inventory-collection-settings
// Jamf Pro API requires the structs to support a JSON data structure.

package jamfpro

import (
	"fmt"
)

const uriComputerInventoryCollectionSettings = "/api/v2/computer-inventory-collection-settings"

// Resource
type ResourceComputerInventoryCollectionSettings struct {
	ComputerInventoryCollectionPreferences ComputerInventoryCollectionSettingsSubsetPreferences    `json:"computerInventoryCollectionPreferences"`
	ApplicationPaths                       []ComputerInventoryCollectionSettingsSubsetPathResponse `json:"applicationPaths"`
}

// Preferences
type ComputerInventoryCollectionSettingsSubsetPreferences struct {
	MonitorApplicationUsage                      bool `json:"monitorApplicationUsage"`
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

// ComputerInventoryCollectionSettingsSubsetPathItem for applicationPaths
type ComputerInventoryCollectionSettingsSubsetPathResponse struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

type ComputerInventoryCollectionSettingsSubsetPathItem struct {
	ID   string `json:"id"`
	Href string `json:"href"`
}

// ComputerInventoryCollectionSettingsCustomPath defines the request body for creating a custom path.
type ResourceComputerInventoryCollectionSettingsCustomPath struct {
	Scope string `json:"scope"`
	Path  string `json:"path"`
}

// CreateComputerInventoryCollectionSettingsCustomPath creates a custom path for computer inventory collection settings.
func (c *Client) CreateComputerInventoryCollectionSettingsCustomPath(customPath *ResourceComputerInventoryCollectionSettingsCustomPath) (*ComputerInventoryCollectionSettingsSubsetPathItem, error) {
	endpoint := fmt.Sprintf("%s/custom-path", uriComputerInventoryCollectionSettings)

	var response ComputerInventoryCollectionSettingsSubsetPathItem
	resp, err := c.HTTP.DoRequest("POST", endpoint, customPath, &response)
	if err != nil {
		return nil, fmt.Errorf(errMsgFailedCreate, "computer inventory collection settings custom path", err)
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
		return fmt.Errorf(errMsgFailedDeleteByID, "computer inventory collection settings custom path", id, err)
	}

	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	return nil
}

// GetComputerInventoryCollectionSettings retrieves a computer inventory collection settings.
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
func (c *Client) UpdateComputerInventoryCollectionSettings(settingsUpdate *ResourceComputerInventoryCollectionSettings) (*ResourceComputerInventoryCollectionSettings, error) {
	endpoint := uriComputerInventoryCollectionSettings

	resp, _ := c.HTTP.DoRequest("PATCH", endpoint, settingsUpdate, nil)

	if resp == nil {
		return nil, fmt.Errorf("failed to update computer inventory collection settings: received nil response")
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 204 {
		return nil, fmt.Errorf("failed to update computer inventory collection settings: unexpected status code %d", resp.StatusCode)
	}

	return nil, nil
}
