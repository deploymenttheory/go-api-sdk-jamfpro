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

	integrationID := 1

	// Call GetApiIntegrationByID function with an integer ID
	integration, err := client.GetApiIntegrationByID(integrationID)
	if err != nil {
		log.Fatalf("Error fetching Jamf API Integration by ID: %v", err)
	}

	// Pretty print the integration in JSON
	integrationJSON, err := json.MarshalIndent(integration, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf API Integration data: %v", err)
	}
	fmt.Println("Fetched Jamf API Integration:\n", string(integrationJSON))
}
