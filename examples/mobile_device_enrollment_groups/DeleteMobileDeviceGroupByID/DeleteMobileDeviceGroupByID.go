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

	// Define the ID of the group you want to delete
	groupID := "1" // Replace with the actual group ID

	// Call the DeleteMobileDeviceGroupByID function
	err = client.DeleteMobileDeviceGroupByID(groupID)
	if err != nil {
		log.Fatalf("Error deleting mobile device group: %s\n", err)
	}

	log.Println("Mobile device group deleted successfully")
}
