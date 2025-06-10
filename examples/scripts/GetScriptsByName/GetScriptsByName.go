package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define a script name for testing
	scriptName := "test_script" // Replace this with an actual script name for testing

	// Call GetScriptsByName function
	script, err := client.GetScriptByName(scriptName)
	if err != nil {
		log.Fatalf("Error fetching script by Name: %v", err)
	}

	fmt.Println(script)

	// Pretty print the script details in XML
	scriptXML, err := json.MarshalIndent(script, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling script details data: %v", err)
	}
	fmt.Println("Fetched Script Details by Name:\n", string(scriptXML))
}
