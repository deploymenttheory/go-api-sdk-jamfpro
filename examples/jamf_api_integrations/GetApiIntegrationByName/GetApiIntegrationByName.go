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

	integrationName := "API Integration Name" // Replace with the actual API Integration name

	response, err := client.GetApiIntegrationByName(integrationName)
	if err != nil {
		log.Fatalf("Error fetching Jamf API Integration by name: %v", err)
	}

	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling API role privileges data: %v", err)
	}
	fmt.Println("Fetched API Role Privileges:", string(responseJSON))
}
