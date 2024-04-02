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

	categoryID := "119" // Replace with the actual category ID you want to delete

	err = client.DeleteCategoryByID(categoryID)
	if err != nil {
		log.Fatalf("Error deleting category: %v", err)
	}

	fmt.Println("Category deleted successfully")
}
