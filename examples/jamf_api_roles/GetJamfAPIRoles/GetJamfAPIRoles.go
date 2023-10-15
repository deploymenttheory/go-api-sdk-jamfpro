package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
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

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Fetch API roles
	apiRoles, err := client.GetJamfAPIRoles()
	if err != nil {
		log.Fatalf("Error fetching API roles: %v", err)
	}

	// Pretty print the fetched API roles using JSON marshaling
	rolesJSON, err := json.MarshalIndent(apiRoles, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling API roles data: %v", err)
	}
	fmt.Println("Fetched API Roles:", string(rolesJSON))
}
