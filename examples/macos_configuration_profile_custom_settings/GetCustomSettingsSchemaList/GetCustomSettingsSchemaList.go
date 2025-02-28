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

	macApps, err := client.GetCustomSettingsSchemaList()
	if err != nil {
		log.Fatalf("Error fetching macos configuration profile settings schema list: %v", err)
	}

	// Pretty print the account details
	byoprofileXML, err := json.MarshalIndent(macApps, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling account data: %v", err)
	}
	fmt.Println("Fetched macos configuration profile settings schema list Details:", string(byoprofileXML))
}
