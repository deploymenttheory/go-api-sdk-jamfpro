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

	// Call GetApiIntegrations function
	integrations, err := client.GetApiIntegrations()
	if err != nil {
		log.Fatalf("Error fetching API Integrations: %v", err)
	}

	// Pretty print the integrations in JSON
	integrationsJSON, err := json.MarshalIndent(integrations, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling API Integrations data: %v", err)
	}
	fmt.Println("Fetched API Integrations:\n", string(integrationsJSON))
}
