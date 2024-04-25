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

	// Call GetAccountPreferences function
	groups, err := client.GetAccountPreferences()
	if err != nil {
		log.Fatalf("Error fetching Account Preferences: %v", err)
	}

	// Pretty print the groups in JSON
	accountPreferencesJSON, err := json.MarshalIndent(groups, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Account Preferences data: %v", err)
	}
	fmt.Println("Fetched Account Preferences:\n", string(accountPreferencesJSON))
}
