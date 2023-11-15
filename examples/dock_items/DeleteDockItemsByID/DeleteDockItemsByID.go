package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file for OAuth credentials
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro client
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     http_client.LogLevelDebug,
		Logger:       http_client.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the ID of the dock item to delete
	dockItemID := 123 // Replace with the actual ID

	// Call the DeleteDockItemsByID function
	err = client.DeleteDockItemsByID(dockItemID)
	if err != nil {
		log.Fatalf("Error deleting dock item by ID: %v", err)
	}

	fmt.Printf("Successfully deleted dock item with ID: %d\n", dockItemID)
}
