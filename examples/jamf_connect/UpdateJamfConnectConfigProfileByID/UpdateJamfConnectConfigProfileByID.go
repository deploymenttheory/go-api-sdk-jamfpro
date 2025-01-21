package main

import (
	"encoding/json"
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

	// Create update payload
	profileUpdate := &jamfpro.ResourceJamfConnectConfigProfileUpdate{
		JamfConnectVersion: "2.43.0",                  // 2.43.0 / 2.42.0 / 2.41.0 / 2.41.0 / 2.40.0 / 2.39.0 / 2.38.0 / 2.37.0 / 2.36.1 / 2.36.0
		AutoDeploymentType: "MINOR_AND_PATCH_UPDATES", // "NONE" / "PATCH_UPDATES" / "MINOR_AND_PATCH_UPDATES"
	}

	// Define the UUID of the profile to update
	profileID := 193

	updatedProfile, err := client.UpdateJamfConnectConfigProfileByID(profileID, profileUpdate)
	if err != nil {
		log.Fatalf("Error updating Jamf Connect config profile: %v", err)
	}

	response, err := json.MarshalIndent(updatedProfile, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated Jamf Connect config profile data: %v", err)
	}
	fmt.Println("Updated Jamf Connect config profile:\n", string(response))
}
