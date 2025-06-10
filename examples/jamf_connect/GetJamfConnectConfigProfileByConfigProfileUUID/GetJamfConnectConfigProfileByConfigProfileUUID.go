package main

import (
	"encoding/json"
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

	// Example UUID - replace with an actual UUID from your environment
	uuid := "your-profile-uuid-here"

	// Call GetJamfConnectConfigProfileByID function
	profile, err := client.GetJamfConnectConfigProfileByConfigProfileUUID(uuid)
	if err != nil {
		log.Fatalf("Error fetching Jamf Connect config profile by ID: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(profile, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf Connect config profile data: %v", err)
	}
	fmt.Println("Fetched Jamf Connect config profile by ID:\n", string(response))
}
