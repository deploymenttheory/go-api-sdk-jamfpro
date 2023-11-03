package main

import (
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	// Define the ID of the mobile device search you want to delete
	SearchIDToDelete = 123
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for Jamf Pro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new Jamf Pro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Use the client to delete an advanced mobile device search by ID
	// Replace 123 with the actual ID
	err = client.DeleteAdvancedMobileDeviceSearchByID(SearchIDToDelete)
	if err != nil {
		log.Fatalf("Error deleting advanced mobile device search by ID: %v", err)
	} else {
		log.Printf("Successfully deleted Advanced Mobile Device Search with ID: %d\n", SearchIDToDelete)
	}
}
