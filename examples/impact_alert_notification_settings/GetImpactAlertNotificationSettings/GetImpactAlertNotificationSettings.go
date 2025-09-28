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

	// Initialize the Jamf Pro client
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Get Impact Alert Notification settings
	settings, err := client.GetImpactAlertNotificationSettings()
	if err != nil {
		log.Fatalf("Failed to get Impact Alert Notification settings: %v", err)
	}

	// Pretty print the results
	settingsJSON, err := json.MarshalIndent(settings, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal settings to JSON: %v", err)
	}

	fmt.Println("Impact Alert Notification Settings:")
	fmt.Println(string(settingsJSON))
}
