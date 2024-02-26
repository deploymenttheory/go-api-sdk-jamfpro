package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	scriptFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/scriptfile.sh" // Replace with your script file path

	// Read script contents from a file
	scriptFile, err := os.ReadFile(scriptFilePath)
	if err != nil {
		log.Fatalf("Error reading script file: %v", err)
	}
	scriptContents := string(scriptFile)

	// Define a sample script for testing
	sampleScript := &jamfpro.ResourceScript{
		Name:           "Sample Script",
		CategoryId:     "None",
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
		ScriptContents: scriptContents,
	}

	// Call CreateScriptByID function
	createdScript, err := client.CreateScript(sampleScript)
	if err != nil {
		log.Fatalf("Error creating script: %v", err)
	}

	// Pretty print the created script details in XML
	createdScriptXML, err := xml.MarshalIndent(createdScript, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created script data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(createdScriptXML))
}
