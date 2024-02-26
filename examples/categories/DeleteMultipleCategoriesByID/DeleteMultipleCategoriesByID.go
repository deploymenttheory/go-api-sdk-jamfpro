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

	// Define the category IDs to delete as strings
	categoryIDsToDelete := []string{"3", "4"} // Replace with the actual category IDs you want to delete

	// Call DeleteMultipleCategoriesByID function
	err = client.DeleteMultipleCategoriesByID(categoryIDsToDelete)
	if err != nil {
		log.Fatalf("Error deleting multiple categories: %v", err)
	}

	fmt.Println("Categories deleted successfully")
}
