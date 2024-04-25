package main

import (
	"fmt"
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

	// Define the ID of the mobile device application you want to delete
	appID := 3 // Replace with the actual ID of the application you want to delete

	// Perform the deletion
	err = client.DeleteMobileDeviceApplicationpByID(appID)
	if err != nil {
		fmt.Println("Error deleting mobile device application:", err)
	} else {
		fmt.Println("Mobile Device Application deleted successfully")
	}
}
