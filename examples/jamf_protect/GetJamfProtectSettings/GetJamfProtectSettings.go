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

	// Get Jamf Protect settings
	settings, err := client.GetJamfProtectSettings()
	if err != nil {
		log.Fatalf("Error fetching Jamf Protect settings: %v", err)
	}

	// Print formatted settings details
	fmt.Println("Jamf Protect Settings:")
	fmt.Printf("ID: %s\n", settings.ID)
	fmt.Printf("Protect URL: %s\n", settings.ProtectURL)
	fmt.Printf("Sync Status: %s\n", settings.SyncStatus)
	fmt.Printf("API Client ID: %s\n", settings.APIClientID)
	fmt.Printf("Auto Install: %t\n", settings.AutoInstall)
	fmt.Printf("Last Sync Time: %s\n", settings.LastSyncTime)
	fmt.Printf("API Client Name: %s\n", settings.APIClientName)
	fmt.Printf("Registration ID: %s\n", settings.RegistrationID)
	fmt.Println("--------------------")

	// Print the entire response as JSON
	settingsJSON, err := json.MarshalIndent(settings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling Jamf Protect settings to JSON: %v", err)
	}
	fmt.Println("\nJamf Protect Settings (JSON):")
	fmt.Println(string(settingsJSON))
}
