package main

import (
	"fmt"
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

	// Define the ID of the VPP mac application you want to delete
	profileID := "1"

	// Call the DeleteMacOSConfigurationProfileByID function
	err = client.DeleteMacApplicationByID(profileID)
	if err != nil {
		log.Fatalf("Failed to delete VPP mac application with ID %d: %v", profileID, err)
	}

	fmt.Printf("VPP mac application with ID %d deleted successfully\n", profileID)
}
