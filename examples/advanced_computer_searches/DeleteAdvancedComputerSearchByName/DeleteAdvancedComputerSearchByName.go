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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	advancedComputerSearchName := "Advanced Computer Search Name" // Replace with the actual advanced computer search name

	// Delete the advanced computer search by name
	err = client.DeleteAdvancedComputerSearchByName(advancedComputerSearchName)
	if err != nil {
		log.Fatalf("Error deleting advanced computer search by name: %v", err)
	}

	log.Println("Advanced computer search deleted successfully by name.")
}
