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

	// Name of the mobile extension attribute to delete
	name := "New Attribute" // Replace with the actual name

	if err := client.DeleteMobileExtensionAttributeByName(name); err != nil {
		log.Fatalf("Error deleting mobile extension attribute by name: %v", err)
	}

	log.Println("Mobile extension attribute deleted successfully.")
}
