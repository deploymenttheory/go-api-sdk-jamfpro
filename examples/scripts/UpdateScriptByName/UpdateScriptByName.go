package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
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
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
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

	scriptNameToUpdate := "Script Name" // Replace with the actual script name
	embeddedScript := `#!/bin/bash echo hello world`

	sampleScript := &jamfpro.ResourceScript{
		Name:           scriptNameToUpdate,
		CategoryName:   "None",
		Info:           "Script information",
		Notes:          "Sample Script",
		Priority:       "Before",
		Parameter4:     "string",
		Parameter5:     "string",
		Parameter6:     "string",
		Parameter7:     "string",
		Parameter8:     "string",
		Parameter9:     "string",
		Parameter10:    "string",
		Parameter11:    "string",
		OSRequirements: "string",
		ScriptContents: embeddedScript,
	}

	updatedScript, err := client.UpdateScriptByName(scriptNameToUpdate, sampleScript)
	if err != nil {
		log.Fatalf("Error updating script by name: %v", err)
	}

	updatedScriptXML, err := xml.MarshalIndent(updatedScript, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated script data: %v", err)
	}
	fmt.Println("Updated Script Details:\n", string(updatedScriptXML))
}
