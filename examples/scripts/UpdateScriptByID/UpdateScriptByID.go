package main

import (
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
