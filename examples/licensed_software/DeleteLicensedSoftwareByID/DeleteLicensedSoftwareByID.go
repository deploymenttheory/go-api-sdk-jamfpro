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

	licensedSoftwareID := "1" // Set your ID here

	// Delete licensed software by ID
	err = client.DeleteLicensedSoftwareByID(licensedSoftwareID)
	if err != nil {
		log.Fatalf("Error deleting licensed software by ID: %v", err)
	}

	log.Println("Licensed software deleted successfully by ID")
}
