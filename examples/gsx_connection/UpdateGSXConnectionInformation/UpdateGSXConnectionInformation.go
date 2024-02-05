package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

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

	// Define new GSX Connection settings
	newGSXSettings := &jamfpro.ResourceGSXConnection{
		Enabled:       false,
		Username:      "", // Empty string to denote no username
		AccountNumber: 0,  // Zero to denote no account number
		URI:           "https://partner-connect.apple.com/gsx/api",
	}

	// Call the UpdateGSXConnectionInformation function
	err = client.UpdateGSXConnectionInformation(newGSXSettings)
	if err != nil {
		log.Fatalf("Error updating GSX Connection Information: %v", err)
	}

	fmt.Println("GSX Connection Information updated successfully.")
}
