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

	// Define the name of the dock item to delete
	dockItemName := "Safari" // Replace with the actual name

	// Call the DeleteDockItemsByName function
	err = client.DeleteDockItemByName(dockItemName)
	if err != nil {
		log.Fatalf("Error deleting dock item by name: %v", err)
	}

	fmt.Printf("Successfully deleted dock item with name: %s\n", dockItemName)
}
