package main

import (
	"encoding/xml"
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

	// Define a script ID for testing
	scriptID := 187

	// Call GetScriptsByID function
	script, err := client.GetScriptsByID(scriptID)
	if err != nil {
		log.Fatalf("Error fetching script by ID: %v", err)
	}

	// Pretty print the script details in XML
	scriptXML, err := xml.MarshalIndent(script, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling script details data: %v", err)
	}
	fmt.Println("Fetched Script Details:\n", string(scriptXML))
}
