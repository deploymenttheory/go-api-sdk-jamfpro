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

	// Define the ID of the dock item to retrieve
	dockItemID := 1 // Replace with the actual ID of the dock item

	// Call the GetDockItemsByID function
	dockItem, err := client.GetDockItemsByID(dockItemID)
	if err != nil {
		log.Fatalf("Error fetching dock item by ID: %v", err)
	}

	// Output the details of the fetched dock item
	fmt.Printf("Fetched Dock Item Details:\n")
	fmt.Printf("ID: %d\n", dockItem.ID)
	fmt.Printf("Name: %s\n", dockItem.Name)
	fmt.Printf("Type: %s\n", dockItem.Type)
	fmt.Printf("Path: %s\n", dockItem.Path)
	fmt.Printf("Contents: %s\n", dockItem.Contents)
}
