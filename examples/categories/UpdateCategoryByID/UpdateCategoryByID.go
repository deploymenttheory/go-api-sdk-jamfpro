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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Define the category ID you want to update and the updated category details
	categoryID := "1" // Replace with the actual category ID you want to update
	updatedCategory := &jamfpro.ResourceCategory{
		Name:     "Updated Category Name", // Replace with the updated name
		Priority: 10,                      // Replace with the updated priority
	}

	// Call UpdateCategoryByID function
	updatedCategoryResult, err := client.UpdateCategoryByID(categoryID, updatedCategory)
	if err != nil {
		log.Fatalf("Error updating category: %v", err)
	}

	// Pretty print the updated category in JSON
	categoryJSON, err := json.MarshalIndent(updatedCategoryResult, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated category data: %v", err)
	}
	fmt.Println("Updated Category:\n", string(categoryJSON))
}
