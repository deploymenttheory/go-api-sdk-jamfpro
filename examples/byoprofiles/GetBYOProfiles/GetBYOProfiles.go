package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Call GetBYOProfiles function
	profiles, err := client.GetBYOProfiles()
	if err != nil {
		log.Fatalf("Failed to get BYO Profiles: %v", err)
	}

	// Print out the profiles
	fmt.Printf("Total BYO Profiles: %d\n", profiles.Size)
	for _, profile := range profiles.BYOProfiles {
		fmt.Printf("ID: %d, Name: %s\n", profile.ID, profile.Name)
	}
}
