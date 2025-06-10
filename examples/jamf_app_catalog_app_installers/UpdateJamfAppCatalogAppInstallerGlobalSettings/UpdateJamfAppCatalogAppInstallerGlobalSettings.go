package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the global settings configuration
	globalSettings := &jamfpro.ResponseJamfAppCatalogGlobalSettings{
		EndUserExperienceSettings: jamfpro.JamfAppCatalogDeploymentSubsetNotificationSettings{
			NotificationMessage:  "A new version is available for installation.",
			NotificationInterval: 24,
			DeadlineMessage:      "Please install within 24 hours.",
			Deadline:             1,
			QuitDelay:            3,
			CompleteMessage:      "Installation completed successfully.",
			Relaunch:             BoolPtr(true),
			Suppress:             BoolPtr(false),
		},
	}

	// Call UpdateJamfAppCatalogAppInstallerGlobalSettings function
	updatedSettings, err := client.UpdateJamfAppCatalogAppInstallerGlobalSettings(globalSettings)
	if err != nil {
		log.Fatalf("Error updating jamf app catalog global settings: %v", err)
	}

	// Pretty print the updated global settings in JSON
	updatedSettingsJSON, err := json.MarshalIndent(updatedSettings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated jamf app catalog global settings data: %v", err)
	}
	fmt.Println("Updated Jamf App Catalog Global Settings:\n", string(updatedSettingsJSON))
}

// BoolPtr returns a pointer to the bool value passed in.
func BoolPtr(b bool) *bool {
	return &b
}
