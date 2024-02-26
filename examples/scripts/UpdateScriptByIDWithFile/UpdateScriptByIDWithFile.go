package main

import (
	"encoding/xml"
	"fmt"
	"io"
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

	scriptFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/scriptfile.sh" // Path to the script file

	// Load script contents from a file
	file, err := os.Open(scriptFilePath)
	if err != nil {
		log.Fatalf("Failed to open script file: %v", err)
	}
	defer file.Close()

	scriptContents, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read script file: %v", err)
	}

	// Define a sample script for testing
	updatedScript := &jamfpro.ResourceScript{
		ID:             "3", // Assuming ID 194 for this example
		Name:           "Updated Sample Script",
		CategoryId:     "None",
		Info:           "Updated Script information",
		Notes:          "Updated Sample Script",
		Priority:       "After",
		OSRequirements: "string",
		ScriptContents: string(scriptContents),
	}

	// Call UpdateScriptByID function
	resultScript, err := client.UpdateScriptByID(updatedScript.ID, updatedScript)
	if err != nil {
		log.Fatalf("Error updating script: %v", err)
	}

	// Pretty print the updated script details in XML
	resultScriptXML, err := xml.MarshalIndent(resultScript, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated script data: %v", err)
	}
	fmt.Println("Updated Script Details:\n", string(resultScriptXML))
}
