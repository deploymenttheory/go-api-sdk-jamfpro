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

	name := "ProfileName" // Replace with the actual profile name
	subset := "general"   // Replace with the actual subset (e.g., "general")

	profile, err := client.GetMobileDeviceEnrollmentProfileByNameWithSubset(name, subset)
	if err != nil {
		fmt.Println("Error fetching profile by name and subset:", err)
		return
	}

	fmt.Printf("Profile: %+v\n", profile)
}
