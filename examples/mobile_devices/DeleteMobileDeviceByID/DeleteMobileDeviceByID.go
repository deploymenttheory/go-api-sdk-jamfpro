package main

import (
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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Example device ID to delete
	deviceID := 1 // Replace with an actual device ID

	// Delete mobile device by ID
	err = client.DeleteMobileDeviceByID(deviceID)
	if err != nil {
		log.Fatalf("Error deleting mobile device by ID: %v", err)
	} else {
		log.Printf("Mobile device with ID %d has been successfully deleted.\n", deviceID)
	}
}
