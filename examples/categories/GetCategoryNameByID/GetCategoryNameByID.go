package main

import (
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

	// Define the category name to search for
	categoryName := "test" // Replace with the desired category name

	// Call GetCategoryNameByID function
	category, err := client.GetCategoryByName(categoryName)
	if err != nil {
		log.Fatalf("Error fetching category by name: %v", err)
	}

	// Print the category details
	fmt.Printf("Fetched Category Details:\nID: %s\nName: %s\nPriority: %d\n",
		category.Id, category.Name, category.Priority)
}
