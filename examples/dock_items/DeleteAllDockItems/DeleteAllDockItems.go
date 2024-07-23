package main

import (
	"fmt"
	"log"
	"strconv"

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

	// Fetch all dock items
	dockItems, err := client.GetDockItems()
	if err != nil {
		log.Fatalf("Error fetching dock items: %v", err)
	}

	fmt.Println("Dock items fetched. Starting deletion process:")

	// Iterate over each dock item and delete
	for _, dockItem := range dockItems.DockItems {
		fmt.Printf("Deleting dock item ID: %d, Name: %s\n", dockItem.ID, dockItem.Name)

		err = client.DeleteDockItemByID(strconv.Itoa(dockItem.ID))
		if err != nil {
			log.Printf("Error deleting dock item ID %d: %v\n", dockItem.ID, err)
			continue // Move to the next dock item if there's an error
		}

		fmt.Printf("Dock item ID %d deleted successfully.\n", dockItem.ID)
	}

	fmt.Println("Dock item deletion process completed.")

}
