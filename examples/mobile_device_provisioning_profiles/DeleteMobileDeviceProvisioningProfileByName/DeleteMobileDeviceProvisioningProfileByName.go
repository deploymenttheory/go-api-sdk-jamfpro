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

	profileName := "ProfileName" // Replace with the actual name of the profile to delete

	err = client.DeleteMobileDeviceProvisioningProfileByName(profileName)
	if err != nil {
		log.Fatalf("Error deleting mobile device provisioning profile by name: %v", err)
	}

	fmt.Printf("Mobile Device Provisioning Profile with name '%s' deleted successfully\n", profileName)
}
