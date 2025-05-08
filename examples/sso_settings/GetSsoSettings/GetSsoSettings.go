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

	// Fetch SSO settings
	fmt.Println("Fetching SSO settings...")
	ssoSettings, err := client.GetSsoSettings()
	if err != nil {
		fmt.Printf("Error fetching SSO settings: %v\n", err)
		return
	}

	// Pretty print the JSON response for SSO settings
	jsonData, err := json.MarshalIndent(ssoSettings, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Printf("SSO Settings: %s\n", jsonData)
}
