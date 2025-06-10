package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
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
		category.Id, category.Name, category.Priority)
}
