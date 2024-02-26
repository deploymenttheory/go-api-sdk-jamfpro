package main

import (
	"encoding/json"
	"fmt"
	"log"

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

	// Define the new category you want to create
	newCategory := &jamfpro.ResourceCategory{
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
