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

	fmt.Println(response)

	// Print the response
	fmt.Printf("Created API Integration with ID: %d and Display Name: %s\n", response.ID, response.DisplayName)
}
