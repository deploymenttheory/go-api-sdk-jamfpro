package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	integrationID := 1

	// Call GetApiIntegrationByID function with an integer ID
	integration, err := client.GetApiIntegrationByID(integrationID)
	if err != nil {
		log.Fatalf("Error fetching Jamf API Integration by ID: %v", err)
	}

	// Pretty print the integration in JSON
	integrationJSON, err := json.MarshalIndent(integration, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf API Integration data: %v", err)
	}
	fmt.Println("Fetched Jamf API Integration:\n", string(integrationJSON))
}
