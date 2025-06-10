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

	apiIntegrationID := "1" // Replace with the actual API Integration ID

	// Create an instance of the ApiIntegration with updated values
	integrationUpdate := &jamfpro.ResourceApiIntegration{
		AuthorizationScopes:        []string{"sdktest"}, // your api role names
		DisplayName:                "My API Integration - new name",
		Enabled:                    true,
		AccessTokenLifetimeSeconds: 1,
	}

	// Update the API Integration using its ID
	response, err := client.UpdateApiIntegrationByID(apiIntegrationID, integrationUpdate)
	if err != nil {
		fmt.Println("Error updating API Integration:", err)
		return
	}

	responseJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling API role privileges data: %v", err)
	}
	fmt.Println("Fetched API Role Privileges:", string(responseJSON))
}
