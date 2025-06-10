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

	profileID := "1"           // Replace with the actual ID
	subset := "desired_subset" // Replace with the desired subset
	profile, err := client.GetMobileDeviceConfigurationProfileByIDWithSubset(profileID, subset)
	if err != nil {
		log.Fatalf("Error fetching mobile device configuration profile by ID and subset: %v", err)
	}

	fmt.Printf("Profile: %+v\n", profile)
}
