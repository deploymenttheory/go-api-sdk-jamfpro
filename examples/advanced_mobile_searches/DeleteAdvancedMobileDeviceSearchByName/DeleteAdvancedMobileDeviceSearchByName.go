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

	SearchNameToDelete := "Search Name" // Replace with the actual search name

	// Use the client to delete an advanced mobile device search by name
	err = client.DeleteAdvancedMobileDeviceSearchByName(SearchNameToDelete)
	if err != nil {
		log.Fatalf("Error deleting advanced mobile device search by name: %v", err)
	} else {
		log.Printf("Successfully deleted Advanced Mobile Device Search with name: %s\n", SearchNameToDelete)
	}
}
