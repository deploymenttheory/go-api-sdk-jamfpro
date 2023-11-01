package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define a category ID for testing
	categoryID := "1" // Replace with an actual category ID

	// Call GetCategoryByID function
	category, err := client.GetCategoryByID(categoryID)
	if err != nil {
		log.Fatalf("Error fetching category by ID: %v", err)
	}

	// Pretty print the category details
	fmt.Printf("Fetched Category Details:\nID: %s\nName: %s\nPriority: %d\n",
		*category.Id, *category.Name, *category.Priority)
}
