package main

import (
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

	// Configuration for the jamfpro client
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     http_client.LogLevelDebug,
		Logger:       http_client.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
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
		General: jamfpro.General{
			Name:         "Steve Job's iMac",
			SerialNumber: "XXXQ7KHTGXXX",                         // Must be Unique
			UDID:         "EBBFF74D-C6B7-5589-93A9-19E8BDXXXXXX", // Must be Unique
			RemoteManagement: jamfpro.RemoteManagement{
				Managed: true,
			},
			Site: jamfpro.Site{
				ID:   -1,
				Name: "None",
			},
		},
		// ... other struct fields ...
	}

	// Call UpdateComputerByID function
	response, err := client.UpdateComputerByID(computerID, updatedComputer)
	if err != nil {
		log.Fatalf("Error updating computer: %v", err)
	}

	// Output the result
	fmt.Printf("Updated Computer: %+v\n", response)
}
