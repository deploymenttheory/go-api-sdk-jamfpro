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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the ID of the macOS Configuration Profile you want to delete
	profileID := 78

	// Call the DeleteMacOSConfigurationProfileByID function
	err = client.DeleteMacOSConfigurationProfileByID(profileID)
	if err != nil {
		log.Fatalf("Failed to delete macOS Configuration Profile with ID %d: %v", profileID, err)
	}

	fmt.Printf("Profile with ID %d deleted successfully\n", profileID)
}
