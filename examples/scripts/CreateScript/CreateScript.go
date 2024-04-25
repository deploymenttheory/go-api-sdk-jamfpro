package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
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

	// Pretty print the created script details in JSON
	createdScriptJSON, err := json.MarshalIndent(createdScript, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created script data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(createdScriptJSON))
}
