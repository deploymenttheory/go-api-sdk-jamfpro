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
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the name of the computer to update
	computerName := "Example-Computer-Name" // Replace with actual computer name

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

	// Call UpdateComputerByName function
	computer, err := client.UpdateComputerByName(computerName, updatedComputer)
	if err != nil {
		log.Fatalf("Error updating computer by name: %v", err)
	}

	// Pretty print the created department in JSON
	computerJSON, err := xml.MarshalIndent(computer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created computer data: %v", err)
	}
	fmt.Println("Created Computer:\n", string(computerJSON))
}
