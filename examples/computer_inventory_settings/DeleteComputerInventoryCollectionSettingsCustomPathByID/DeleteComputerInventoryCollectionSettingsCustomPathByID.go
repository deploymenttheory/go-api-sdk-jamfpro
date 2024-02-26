package main

import (
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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// ID of the custom path to delete.
	// Replace "123" with the actual ID of the custom path you want to delete.
	id := "3"

	// Call the delete function.
	err = client.DeleteComputerInventoryCollectionSettingsCustomPathByID(id)
	if err != nil {
		log.Fatalf("Error deleting custom path with ID %s: %s", id, err)
	}

	// If no error, print success message.
	fmt.Printf("Custom path with ID %s deleted successfully.\n", id)
}
