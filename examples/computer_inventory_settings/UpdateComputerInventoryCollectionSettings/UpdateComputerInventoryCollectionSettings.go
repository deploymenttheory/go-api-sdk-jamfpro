package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the new settings
	newSettings := &jamfpro.ResponseComputerInventoryCollectionSettings{
		ComputerInventoryCollectionPreferences: jamfpro.InventoryCollectionPreference{
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
		ApplicationPaths: []jamfpro.PathItem{
			{
				ID:   "1",
				Path: "/Example/Path/To/App/",
			},
		},
		FontPaths: []jamfpro.PathItem{
			{
				ID:   "2",
				Path: "/Example/Path/To/Font/",
			},
		},
		PluginPaths: []jamfpro.PathItem{
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
