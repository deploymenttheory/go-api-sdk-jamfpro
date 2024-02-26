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

	profileID := 1 // Replace with actual ID

	err = client.DeleteMobileDeviceEnrollmentProfileByID(profileID)
	if err != nil {
		log.Fatalf("Error deleting profile by name: %v", err)
	}

	fmt.Println("Profile deleted successfully.")
}
