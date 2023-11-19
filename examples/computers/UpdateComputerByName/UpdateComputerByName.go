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

	// Define the name of the computer to update
	computerName := "Example-Computer-Name" // Replace with actual computer name

	// Define the computer configuration to be updated
	updatedComputer := jamfpro.ResponseComputer{
		// Populate with the updated fields
		General: jamfpro.General{
			Name:         "Updated Computer Name", // Updated name or other fields
			SerialNumber: "XXXQ7KHTGXXX",
			UDID:         "EBBFF74D-C6B7-5589-93A9-19E8BDXXXXXX",
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

	// Call UpdateComputerByName function
	response, err := client.UpdateComputerByName(computerName, updatedComputer)
	if err != nil {
		log.Fatalf("Error updating computer by name: %v", err)
	}

	// Output the result
	fmt.Printf("Updated Computer by Name: %+v\n", response)
}
