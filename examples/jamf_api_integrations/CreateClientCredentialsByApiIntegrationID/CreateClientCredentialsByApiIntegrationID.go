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

	integrationID := "1" // Replace with the actual API Integration ID

	// Reset client credentials for the given API Integration ID
	response, err := client.RefreshClientCredentialsByApiRoleID(integrationID)
	if err != nil {
		fmt.Println("Error resetting client credentials:", err)
		return
	}

	// Print the updated credentials
	fmt.Printf("Refreshed client credentials - Client ID: %s, Client Secret: %s\n", response.ClientID, response.ClientSecret)
}
