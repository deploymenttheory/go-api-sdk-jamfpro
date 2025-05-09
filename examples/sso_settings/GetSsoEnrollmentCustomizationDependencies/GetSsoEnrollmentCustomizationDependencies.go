package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Fetch SSO dependencies
	fmt.Println("Fetching SSO dependencies...")
	ssoDependencies, err := client.GetSsoEnrollmentCustomizationDependencies()
	if err != nil {
		fmt.Printf("Error fetching SSO dependencies: %v\n", err)
		return
	}

	// Pretty print the JSON response for SSO dependencies
	dependenciesJSON, err := json.MarshalIndent(ssoDependencies, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal SSO dependencies JSON: %v", err)
	}
	fmt.Printf("SSO Dependencies:\n%s\n", dependenciesJSON)
}
