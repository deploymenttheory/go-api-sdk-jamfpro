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

	// Configuration for the Jamf Pro API
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new Jamf Pro API client
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the category details for creation
	newCategory := &jamfpro.ResponseCategories{
		Name:     jamfpro.String("New Category Name"),
		Priority: jamfpro.Int(1),
	}

	// Create the category
	createdCategory, err := client.CreateCategory(newCategory)
	if err != nil {
		fmt.Println("Error creating category:", err)
		return
	}

	// Print the response
	fmt.Printf("Created Category with ID: %s and Name: %s\n", jamfpro.StringValue(createdCategory.Id), jamfpro.StringValue(createdCategory.Name))
}
