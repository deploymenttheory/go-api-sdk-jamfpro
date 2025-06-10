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

	err = client.DeleteUserExtensionAttributeByID("1") // Replace 1 with the actual ID
	if err != nil {
		log.Fatalf("Error deleting user extension attribute by ID: %v", err)
	}

	log.Println("User extension attribute deleted successfully by ID")
}
