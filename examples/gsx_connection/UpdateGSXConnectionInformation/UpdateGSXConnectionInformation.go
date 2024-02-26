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

	// Define new GSX Connection settings
	newGSXSettings := &jamfpro.ResourceGSXConnection{
		Enabled:       false,
		Username:      "", // Empty string to denote no username
		AccountNumber: 0,  // Zero to denote no account number
		URI:           "https://partner-connect.apple.com/gsx/api",
	}

	// Call the UpdateGSXConnectionInformation function
	err = client.UpdateGSXConnectionInformation(newGSXSettings)
	if err != nil {
		log.Fatalf("Error updating GSX Connection Information: %v", err)
	}

	fmt.Println("GSX Connection Information updated successfully.")
}
