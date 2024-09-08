package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig_oauth2_tfproduction.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Fetch t&c status
	status, err := client.GetJamfAppCatalogAppInstallerTermsAndConditionsStatus()
	if err != nil {
		log.Fatalf("Error fetching Terms And Conditions Status: %v", err)
	}

	// Pretty print the fetched Terms And Conditions Status using JSON marshaling
	statusJSON, err := json.MarshalIndent(status, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Terms And Conditions Status data: %v", err)
	}
	fmt.Println("Fetched Terms And Conditions Status:", string(statusJSON))
}
