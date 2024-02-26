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

	// Define the name of the Jamf API Role you want to delete
	roleName := "One Role to Rule them all"

	// Call DeleteJamfApiRoleByName function
	err = client.DeleteJamfApiRoleByName(roleName)
	if err != nil {
		log.Fatalf("Error deleting Jamf API role by name: %v", err)
	}

	log.Println("Successfully deleted Jamf API Role with name:", roleName)
}
