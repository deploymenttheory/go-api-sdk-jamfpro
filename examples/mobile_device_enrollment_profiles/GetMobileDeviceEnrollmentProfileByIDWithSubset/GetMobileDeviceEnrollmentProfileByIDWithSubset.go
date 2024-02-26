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

	profileID := 1      // Replace with the actual profile ID
	subset := "general" // Replace with the actual subset (e.g., "general")

	profile, err := client.GetMobileDeviceEnrollmentProfileByIDWithSubset(profileID, subset)
	if err != nil {
		fmt.Println("Error fetching profile by name and subset:", err)
		return
	}

	fmt.Printf("Profile: %+v\n", profile)
}
