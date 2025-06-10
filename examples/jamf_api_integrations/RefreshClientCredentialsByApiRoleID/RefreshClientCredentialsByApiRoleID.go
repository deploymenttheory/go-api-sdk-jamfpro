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

	integrationID := "5" // Replace with the actual API Integration ID

	// Reset client credentials for the given API Integration ID
	response, err := client.RefreshClientCredentialsByApiRoleID(integrationID)
	if err != nil {
		fmt.Println("Error resetting client credentials:", err)
		return
	}

	privilegesJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling API role privileges data: %v", err)
	}
	fmt.Println("Fetched API Role Privileges:", string(privilegesJSON))
}
