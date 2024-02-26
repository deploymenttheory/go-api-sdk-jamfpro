package main

import (
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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	apiIntegrationID := 1 // Replace with the actual API Integration ID

	// Create an instance of the ApiIntegration with updated values
	integrationUpdate := &jamfpro.ResourceApiIntegration{
		AuthorizationScopes:        []string{"sdktest"}, // your api role names
		DisplayName:                "My API Integration - new name",
		Enabled:                    true,
		AccessTokenLifetimeSeconds: 1,
	}

	// Update the API Integration using its ID
	updatedIntegration, err := client.UpdateApiIntegrationByID(apiIntegrationID, integrationUpdate)
	if err != nil {
		fmt.Println("Error updating API Integration:", err)
		return
	}

	// Print the response
	fmt.Printf("Updated API Integration with ID: %d and Display Name: %s\n", apiIntegrationID, updatedIntegration.DisplayName)
}
