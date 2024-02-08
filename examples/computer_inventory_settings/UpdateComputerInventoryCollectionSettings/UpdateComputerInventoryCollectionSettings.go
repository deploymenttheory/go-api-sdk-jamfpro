package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the new settings
	newSettings := &jamfpro.ResourceComputerInventoryCollectionSettings{
		ComputerInventoryCollectionPreferences: struct {
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
		}{
			MonitorApplicationUsage:       false,
			IncludeFonts:                  false,
			IncludePlugins:                false,
			IncludePackages:               true,
			IncludeSoftwareUpdates:        false,
			IncludeSoftwareId:             true,
			IncludeAccounts:               true,
			CalculateSizes:                false,
			IncludeHiddenAccounts:         false,
			IncludePrinters:               true,
			IncludeServices:               true,
			CollectSyncedMobileDeviceInfo: false,
			UpdateLdapInfoOnComputerInventorySubmissions: false,
			MonitorBeacons:               false,
			AllowChangingUserAndLocation: true,
			UseUnixUserPaths:             true,
			CollectUnmanagedCertificates: true,
		},
		ApplicationPaths: []jamfpro.ComputerInventoryCollectionSettingsSubsetPathItem{
			{
				ID:   "1",
				Path: "/Example/Path/To/App/",
			},
		},
		FontPaths: []jamfpro.ComputerInventoryCollectionSettingsSubsetPathItem{
			{
				ID:   "2",
				Path: "/Example/Path/To/Font/",
			},
		},
		PluginPaths: []jamfpro.ComputerInventoryCollectionSettingsSubsetPathItem{
			{
				ID:   "3",
				Path: "/Example/Path/To/Plugin/",
			},
		},
	}

	// Update computer inventory collection settings
	updatedSettings, err := client.UpdateComputerInventoryCollectionSettings(newSettings)
	if err != nil {
		log.Fatalf("Error updating Computer Inventory Collection Settings: %s", err)
	}

	// Convert the updated settings to pretty-printed JSON
	updatedSettingsJSON, err := json.MarshalIndent(updatedSettings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling updated Computer Inventory Collection Settings to JSON: %s", err)
	}

	// Print the pretty-printed JSON of the updated settings
	fmt.Println("Updated Computer Inventory Collection Settings:")
	fmt.Println(string(updatedSettingsJSON))
}
