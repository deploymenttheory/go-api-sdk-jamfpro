package jamfpro

import (
	"fmt"
)

const uriComputerInventoryCollectionSettings = "/api/v1/computer-inventory-collection-settings"

type ResponseComputerInventoryCollectionSettings struct {
	Preferences      ComputerInventoryCollectionPreferences `json:"computerInventoryCollectionPreferences"`
	ApplicationPaths []InventoryPath                        `json:"applicationPaths"`
	FontPaths        []InventoryPath                        `json:"fontPaths"`
	PluginPaths      []InventoryPath                        `json:"pluginPaths"`
}

type ResponseCustomPathCreation struct {
	Message string `json:"message"`
}

type ComputerInventoryCollectionPreferences struct {
	MonitorApplicationUsage                      bool `json:"monitorApplicationUsage"`
	IncludeFonts                                 bool `json:"includeFonts"`
	IncludePlugins                               bool `json:"includePlugins"`
	IncludePackages                              bool `json:"includePackages"`
	IncludeSoftwareUpdates                       bool `json:"includeSoftwareUpdates"`
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
}

type InventoryPath struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

type CustomPathCreation struct {
	Scope string `json:"scope"`
	Path  string `json:"path"`
}

func (c *Client) GetComputerInventoryCustomPathCollectionSettings() (*ResponseComputerInventoryCollectionSettings, error) {
	uri := uriComputerInventoryCollectionSettings

	var out ResponseComputerInventoryCollectionSettings
	err := c.DoRequest("GET", uri, nil, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to get computer inventory collection settings: %v", err)
	}

	return &out, nil
}

func (c *Client) UpdateComputerInventoryCustomPathCollectionSettings(settings *ResponseComputerInventoryCollectionSettings) (*ResponseComputerInventoryCollectionSettings, error) {
	uri := uriComputerInventoryCollectionSettings

	var out ResponseComputerInventoryCollectionSettings
	err := c.DoRequest("PATCH", uri, settings, nil, &out)
	if err != nil {
		return nil, fmt.Errorf("failed to update computer inventory collection settings: %v", err)
	}

	return &out, nil
}

func (c *Client) CreateComputerInventoryCustomPathCollection(scope, path string) (*ResponseCustomPathCreation, error) {
	uri := uriComputerInventoryCollectionSettings + "/custom-path"

	requestPayload := &CustomPathCreation{
		Scope: scope,
		Path:  path,
	}

	var response ResponseCustomPathCreation
	err := c.DoRequest("POST", uri, requestPayload, nil, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to create computer inventory collection custom path: %v", err)
	}

	return &response, nil
}

func (c *Client) DeleteComputerInventoryCollectionCustomPath(id string) error {
	uri := fmt.Sprintf("%s/custom-path/%s", uriComputerInventoryCollectionSettings, id)

	err := c.DoRequest("DELETE", uri, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to delete computer inventory collection custom path: %v", err)
	}

	return nil
}
