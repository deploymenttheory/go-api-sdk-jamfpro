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
	// Define the ID of the Jamf API Role you want to delete
	roleID := "13"

	// Call DeleteJamfApiRoleByID function
	err = client.DeleteJamfApiRoleByID(roleID)
	if err != nil {
		log.Fatalf("Error deleting Jamf API role by ID: %v", err)
	}

	log.Println("Successfully deleted the Jamf API role.")
}
