package main

import (
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

	// Delete advanced user search by ID
	err = client.DeleteAdvancedUserSearchByID("123") // Replace 123 with the actual ID
	if err != nil {
		log.Fatalf("Error deleting advanced user search by ID: %v", err)
	}
	log.Println("Advanced user search deleted successfully by ID.")
}
