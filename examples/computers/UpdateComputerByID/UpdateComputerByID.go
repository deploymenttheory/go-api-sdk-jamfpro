package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the computer ID to update
	computerID := 6 // Replace with actual computer ID

	// Define the computer configuration to be updated
	updatedComputer := jamfpro.ResponseComputer{
		// Populate with the updated fields
		General: jamfpro.ComputerSubsetGeneral{
			Name:         "Steve Job's iMac",
			SerialNumber: "XXXQ7KHTGXXX",                         // Must be Unique
			UDID:         "EBBFF74D-C6B7-5589-93A9-19E8BDXXXXXX", // Must be Unique
			RemoteManagement: jamfpro.ComputerSubsetGeneralRemoteManagement{
				Managed: true,
			},
			Site: jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		// ... other struct fields ...
	}

	// Call UpdateComputerByID function
	computer, err := client.UpdateComputerByID(computerID, updatedComputer)
	if err != nil {
		log.Fatalf("Error updating computer: %v", err)
	}

	// Pretty print the created department in JSON
	computerJSON, err := xml.MarshalIndent(computer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created computer data: %v", err)
	}
	fmt.Println("Created Computer:\n", string(computerJSON))
}
