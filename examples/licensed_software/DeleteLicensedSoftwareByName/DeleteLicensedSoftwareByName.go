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

	licensedSoftwareName := "Adobe Creative Suite"

	// Delete licensed software by name
	err = client.DeleteLicensedSoftwareByName(licensedSoftwareName)
	if err != nil {
		log.Fatalf("Error deleting licensed software by name: %v", err)
	}

	log.Println("Licensed software deleted successfully by name")
}
