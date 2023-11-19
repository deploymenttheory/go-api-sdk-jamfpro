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

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the name of the dock item to retrieve
	dockItemName := "Safari" // Replace with the actual name of the dock item

	// Call the GetDockItemsByName function
	dockItem, err := client.GetDockItemsByName(dockItemName)
	if err != nil {
		log.Fatalf("Error fetching dock item by name: %v", err)
	}

	// Output the details of the fetched dock item
	fmt.Printf("Fetched Dock Item Details:\n")
	fmt.Printf("ID: %d\n", dockItem.ID)
	fmt.Printf("Name: %s\n", dockItem.Name)
	fmt.Printf("Type: %s\n", dockItem.Type)
	fmt.Printf("Path: %s\n", dockItem.Path)
	fmt.Printf("Contents: %s\n", dockItem.Contents)
}
