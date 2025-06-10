package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the ID of the dock item to delete
	dockItemID := "1" // Replace with the actual ID

	// Call the DeleteDockItemsByID function
	err = client.DeleteDockItemByID(dockItemID)
	if err != nil {
		log.Fatalf("Error deleting dock item by ID: %v", err)
	}

	fmt.Printf("Successfully deleted dock item with ID: %d\n", dockItemID)
}
