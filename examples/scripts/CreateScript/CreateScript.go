package main

import (
	"encoding/json"
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
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

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
	sampleScript := &jamfpro.ResourceScript{
		Name:           "Sample Script",
		CategoryName:   "None",
		Info:           "Script information",
		Notes:          "Sample Script",
		Priority:       "BEFORE",
		Parameter4:     "string",
		Parameter5:     "string",
		Parameter6:     "string",
		Parameter7:     "string",
		Parameter8:     "string",
		Parameter9:     "string",
		Parameter10:    "string",
		Parameter11:    "string",
		OSRequirements: "string",
		ScriptContents: "echo \"Sample script\"",
	}

	// Call CreateScriptByID function
	createdScript, err := client.CreateScript(sampleScript)
	if err != nil {
		log.Fatalf("Error creating script: %v", err)
	}

	// Pretty print the created script details in XML
	createdScriptXML, err := json.MarshalIndent(createdScript, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created script data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(createdScriptXML))
}
