// get_device_communication_settings.go
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

	// Get device communication settings
	settings, err := client.GetDeviceCommunicationSettings()
	if err != nil {
		log.Fatalf("Error getting device communication settings: %v", err)
	}

	// Pretty print the device communication settings using JSON marshaling
	settingsJSON, err := json.MarshalIndent(settings, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling device communication settings data: %v", err)
	}
	fmt.Println("Device Communication Settings:", string(settingsJSON))
}
