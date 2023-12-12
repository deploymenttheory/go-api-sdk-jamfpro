package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/joseph/github/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelInfo // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define a sample script for testing
	updatedScript := &jamfpro.ResourceScript{
		Name: "Another new name",
	}

	// Call UpdateScriptByID function
	resultScript, err := client.UpdateScriptByID("2", updatedScript)
	if err != nil {
		log.Fatalf("Error updating script: %v", err)
	}

	fmt.Println(resultScript)

	// Pretty print the updated script details in XML
	// resultScriptXML, err := xml.MarshalIndent(resultScript, "", "    ") // Indent with 4 spaces
	// if err != nil {
	// 	log.Fatalf("Error marshaling updated script data: %v", err)
	// }
	// fmt.Println("Updated Script Details with Embedded Script:\n", string(resultScriptXML))
}
