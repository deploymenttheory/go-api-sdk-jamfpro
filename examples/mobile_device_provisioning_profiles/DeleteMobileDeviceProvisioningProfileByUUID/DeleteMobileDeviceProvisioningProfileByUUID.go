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

	profileUUID := "116AF1E6-7EB5-4335-B598-276CDE5E015B" // Replace with the actual UUID of the profile to delete

	err = client.DeleteMobileDeviceProvisioningProfileByUUID(profileUUID)
	if err != nil {
		log.Fatalf("Error deleting mobile device provisioning profile by UUID: %v", err)
	}

	fmt.Printf("Mobile Device Provisioning Profile with UUID '%s' deleted successfully\n", profileUUID)
}
