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

	// Create new settings object
	newSettings := jamfpro.ResourceAccessManagementSettings{
		AutomatedDeviceEnrollmentServerUuid: "12345678-1234-1234-1234-123456789012",
	}

	// Create the settings
	createdSettings, err := client.CreateAccessManagementSettings(newSettings)
	if err != nil {
		log.Fatalf("Failed to create access management settings: %v", err)
	}

	// Pretty print the results
	settingsJSON, err := json.MarshalIndent(createdSettings, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}

	fmt.Println(string(settingsJSON))
}
