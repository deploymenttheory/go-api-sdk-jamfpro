package main

import (
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

	// ID of the mobile extension attribute to delete
	id := 1 // Replace with the actual ID

	if err := client.DeleteMobileExtensionAttributeByID(id); err != nil {
		log.Fatalf("Error deleting mobile extension attribute by ID: %v", err)
	}

	log.Println("Mobile extension attribute deleted successfully.")
}
