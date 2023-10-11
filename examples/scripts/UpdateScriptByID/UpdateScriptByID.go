package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	concurrentRequests           = 10 // Number of simultaneous requests.
	maxConcurrentRequestsAllowed = 5  // Maximum allowed concurrent requests.
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
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
		InstanceName:          authConfig.InstanceName,
		DebugMode:             true,
		Logger:                jamfpro.NewDefaultLogger(),
		MaxConcurrentRequests: maxConcurrentRequestsAllowed,
		TokenLifespan:         defaultTokenLifespan,
		BufferPeriod:          defaultBufferPeriod,
		ClientID:              authConfig.ClientID,
		ClientSecret:          authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client := jamfpro.NewClient(config)

	// Use the embedded script
	embeddedScriptContents := `echo "Sample script"`

	// Define a sample script for testing
	updatedScript := &jamfpro.ResponseScript{
		ID:             195, // Assuming ID 1 for this example
		Name:           "Embedded Sample Script",
		Category:       "None",
		Filename:       "string",
		Info:           "Script information",
		Notes:          "Sample Script",
		Priority:       "Before",
		OSRequirements: "string",
		ScriptContents: embeddedScriptContents,
	}

	// Call UpdateScriptByID function
	resultScript, err := client.UpdateScriptByID(updatedScript)
	if err != nil {
		log.Fatalf("Error updating script: %v", err)
	}

	// Pretty print the updated script details in XML
	resultScriptXML, err := xml.MarshalIndent(resultScript, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated script data: %v", err)
	}
	fmt.Println("Updated Script Details with Embedded Script:\n", string(resultScriptXML))
}
