package main

import (
	"encoding/json"
	"fmt"
	"log"

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

	// Define the new category you want to create
	newCategory := &jamfpro.ResponseCategories{
		Name:     "Applications",
		Priority: 9,
	}

	// Call CreateCategory function
	createdCategory, err := client.CreateCategory(newCategory)
	if err != nil {
		log.Fatalf("Error creating category: %v", err)
	}

	// Pretty print the created category in JSON
	categoryJSON, err := json.MarshalIndent(createdCategory, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling category data: %v", err)
	}
	fmt.Println("Created Category:\n", string(categoryJSON))
}