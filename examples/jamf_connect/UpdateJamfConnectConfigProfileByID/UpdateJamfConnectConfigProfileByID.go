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
		Version:            "2.3.0",
		AutoDeploymentType: "PATCH_UPDATES", // Can be PATCH_UPDATES, MINOR_AND_PATCH_UPDATES, INITIAL_INSTALLATION_ONLY, or NONE
	}

	// Define the UUID of the profile to update
	profileUUID := "d265dfd3-8fde-4bf2-aa56-b167c8b68069"

	updatedProfile, err := client.UpdateJamfConnectConfigProfileByID(profileUUID, profileUpdate)
	if err != nil {
		log.Fatalf("Error updating Jamf Connect config profile: %v", err)
	}

	response, err := json.MarshalIndent(updatedProfile, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated Jamf Connect config profile data: %v", err)
	}
	fmt.Println("Updated Jamf Connect config profile:\n", string(response))
}
