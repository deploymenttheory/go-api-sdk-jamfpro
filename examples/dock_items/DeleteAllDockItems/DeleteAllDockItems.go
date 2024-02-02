package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
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

		err = client.DeleteDockItemByID(dockItem.ID)
		if err != nil {
			log.Printf("Error deleting dock item ID %d: %v\n", dockItem.ID, err)
			continue // Move to the next dock item if there's an error
		}

		fmt.Printf("Dock item ID %d deleted successfully.\n", dockItem.ID)
	}

	fmt.Println("Dock item deletion process completed.")

}
