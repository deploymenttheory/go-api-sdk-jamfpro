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

	categoryName := "Updated Applications Category Name" // Replace with the actual category name you want to delete

	err = client.DeleteCategoryByName(categoryName)
	if err != nil {
		log.Fatalf("Error deleting category: %v", err)
	}

	fmt.Println("Category deleted successfully")
}
