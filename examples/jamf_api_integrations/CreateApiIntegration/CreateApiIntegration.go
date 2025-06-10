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

	// Sample request for creating a new API Integration
	integration := &jamfpro.ResourceApiIntegration{
		AuthorizationScopes:        []string{"sdktest"}, // insert api roles here
		DisplayName:                "My API Integration",
		Enabled:                    true,
		AccessTokenLifetimeSeconds: 300,
	}

	// Create the API Integration
	response, err := client.CreateApiIntegration(integration)
	if err != nil {
		fmt.Println("Error creating API Integration:", err)
		return
	}

	privilegesJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling API role privileges data: %v", err)
	}
	fmt.Println("Fetched API Role Privileges:", string(privilegesJSON))
}
