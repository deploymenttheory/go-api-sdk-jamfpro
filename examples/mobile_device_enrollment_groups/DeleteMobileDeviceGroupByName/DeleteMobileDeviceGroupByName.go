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
	// Define the name of the group you want to delete
	groupName := "Sample Smart Group" // Replace with the actual group name

	// Call the DeleteMobileDeviceGroupByName function
	err = client.DeleteMobileDeviceGroupByName(groupName)
	if err != nil {
		log.Fatalf("Error deleting mobile device group: %s\n", err)
	}

	log.Println("Mobile device group deleted successfully")
}
