package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// The ID of the computer prestage to retrieve the device scope for
	prestageID := "123" // Replace with the actual ID

	// Fetch the device scope for the specified computer prestage
	deviceScope, err := client.GetDeviceScopeForComputerPrestage(prestageID)
	if err != nil {
		log.Fatalf("Error fetching device scope for computer prestage: %v", err)
	}

	// Print out the fetched device scope
	fmt.Printf("Device Scope for Prestage ID: %s\n", prestageID)
	for _, assignment := range deviceScope.Assignments {
		fmt.Printf("Serial Number: %s, Assigned Date: %s, User Assigned: %s\n",
			assignment.SerialNumber,
			assignment.AssignmentDate,
			assignment.UserAssigned,
		)
		// Add more details to print as needed
	}
}
