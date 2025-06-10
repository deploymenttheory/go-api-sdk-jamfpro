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

	scriptFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/scriptfile.sh"

	scriptNameToUpdate := "Embedded Sample Script" // The name of the script to update.

	file, err := os.Open(scriptFilePath)
	if err != nil {
		log.Fatalf("Error opening script file: %v", err)
	}
	defer file.Close()

	scriptContents, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading script file: %v", err)
	}

	sampleScript := &jamfpro.ResourceScript{
		Name:           scriptNameToUpdate,
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
		ScriptContents: string(scriptContents),
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
