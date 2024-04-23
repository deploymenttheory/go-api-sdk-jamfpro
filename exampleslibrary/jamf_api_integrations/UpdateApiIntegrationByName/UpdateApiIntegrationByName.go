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

	apiIntegrationName := "sdktest" // Replace with the actual API Integration display name

	// Create an instance of the ApiIntegration with updated values
	integrationUpdate := &jamfpro.ResourceApiIntegration{
		AuthorizationScopes:        []string{"sdktest2"},
		DisplayName:                "updated by sdk using name",
		Enabled:                    true,
		AccessTokenLifetimeSeconds: 300,
	}

	// Update the API Integration using its display name
	updatedIntegration, err := client.UpdateApiIntegrationByName(apiIntegrationName, integrationUpdate)
	if err != nil {
		fmt.Println("Error updating API Integration:", err)
		return
	}

	// Print the response
	fmt.Printf("Updated API Integration with Display Name: %s to New Display Name: %s\n", apiIntegrationName, updatedIntegration.DisplayName)
}
