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

	// Define the name and details of the dock item to update
	dockItemName := "Safari" // Replace with the actual name
	updatedDockItem := &jamfpro.ResponseDockItem{
		// Set the fields you want to update
		Name:     "Updated Safari",
		Type:     "App",
		Path:     "file://localhost/Applications/Safari.app/",
		Contents: "Updated Contents",
	}

	// Call the UpdateDockItemByName function
	result, err := client.UpdateDockItemsByName(dockItemName, updatedDockItem)
	if err != nil {
		log.Fatalf("Error updating dock item by Name: %v", err)
	}

	// Output the result
	fmt.Printf("Updated Dock Item by Name: %+v\n", result)
}
