package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
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

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Sample request for creating a new API Integration
	integration := &jamfpro.ApiIntegration{
		AuthorizationScopes:        []string{"api-role-1", "api-role-2"}, // insert api roles here
		DisplayName:                "My API Integration",
		Enabled:                    true,
		AccessTokenLifetimeSeconds: 300,
	}

	// Create the API Integration
	response, err := client.CreateApiIntegration(integration)
	if err != nil {
		fmt.Println("Error creating API Integration:", err)
		return
	}

	// Print the response
	fmt.Printf("Created API Integration with ID: %d and Display Name: %s\n", response.ID, response.DisplayName)
}
