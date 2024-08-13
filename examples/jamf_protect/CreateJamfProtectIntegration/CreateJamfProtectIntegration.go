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

	// Set the autoInstall parameter
	autoInstall := true

	// Call CreateJamfProtectIntegration function
	integrationSettings, err := client.CreateJamfProtectIntegration(autoInstall)
	if err != nil {
		log.Fatalf("Error creating Jamf Protect integration: %v", err)
	}

	// Optionally, you can also print the entire response as JSON
	settingsJSON, err := json.MarshalIndent(integrationSettings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling Jamf Protect integration settings to JSON: %v", err)
	}
	fmt.Println("\nJamf Protect Integration Settings (JSON):")
	fmt.Println(string(settingsJSON))
}
