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

	SearchIDToDelete := 123

	// Use the client to delete an advanced mobile device search by ID
	// Replace 123 with the actual ID
	err = client.DeleteAdvancedMobileDeviceSearchByID(SearchIDToDelete)
	if err != nil {
		log.Fatalf("Error deleting advanced mobile device search by ID: %v", err)
	} else {
		log.Printf("Successfully deleted Advanced Mobile Device Search with ID: %d\n", SearchIDToDelete)
	}
}
