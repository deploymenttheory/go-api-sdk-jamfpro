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

	// Delete profile by ID
	if err := client.DeleteMobileDeviceConfigurationProfileByID("2"); err != nil {
		fmt.Println("Error deleting profile by ID:", err)
	} else {
		fmt.Println("Profile deleted successfully by ID")
	}

}
