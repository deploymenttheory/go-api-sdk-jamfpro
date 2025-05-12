package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create settings update request
	updateSettings := jamfpro.ResourceJamfProtectSettings{
		AutoInstall: true, // Set this to false if you want to disable auto-install
	}

	// Update Jamf Protect settings
	updatedSettings, err := client.UpdateJamfProtectSettings(updateSettings)
	if err != nil {
		log.Fatalf("Error updating Jamf Protect settings: %v", err)
	}

	// Print success message and updated settings details
	fmt.Println("Successfully updated Jamf Protect settings")
	fmt.Println("\nUpdated Settings:")
	fmt.Printf("ID: %s\n", updatedSettings.ID)
	fmt.Printf("Protect URL: %s\n", updatedSettings.ProtectURL)
	fmt.Printf("Sync Status: %s\n", updatedSettings.SyncStatus)
	fmt.Printf("API Client ID: %s\n", updatedSettings.APIClientID)
	fmt.Printf("Auto Install: %t\n", updatedSettings.AutoInstall)
	fmt.Printf("Last Sync Time: %s\n", updatedSettings.LastSyncTime)
	fmt.Printf("API Client Name: %s\n", updatedSettings.APIClientName)
	fmt.Printf("Registration ID: %s\n", updatedSettings.RegistrationID)
	fmt.Println("--------------------")

	// Print the entire response as JSON
	settingsJSON, err := json.MarshalIndent(updatedSettings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling updated Jamf Protect settings to JSON: %v", err)
	}
	fmt.Println("\nUpdated Jamf Protect Settings (JSON):")
	fmt.Println(string(settingsJSON))
}
