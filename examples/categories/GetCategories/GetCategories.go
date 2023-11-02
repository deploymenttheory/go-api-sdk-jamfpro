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

	// Define the query parameters
	page := 0
	pageSize := 100
	sort := "id:asc"
	filter := ""

	// Call the GetCategories function
	categories, err := client.GetCategories(page, pageSize, sort, filter)
	if err != nil {
		fmt.Printf("Error fetching categories: %v\n", err)
		return
	}

	// Print the fetched categories
	fmt.Println("Fetched Categories:")
	for _, category := range categories.Results {
		fmt.Printf("ID: %s, Name: %s, Priority: %d\n", category.Id, category.Name, category.Priority)
	}
}
