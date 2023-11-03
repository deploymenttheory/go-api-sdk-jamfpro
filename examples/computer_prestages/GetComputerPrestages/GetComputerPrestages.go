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

	// Define pagination and sorting parameters
	page := 0            // starting page
	pageSize := 10       // number of items per page
	sort := []string{""} // sort by name, for example

	// Fetch computer prestages using the V3 API
	prestages, err := client.GetComputerPrestagesV3(page, pageSize, sort)
	if err != nil {
		log.Fatalf("Error fetching computer prestages: %v", err)
	}

	// Print out the fetched computer prestages
	for _, prestage := range prestages.Results {
		fmt.Printf("Prestage Name: %s\n", prestage.DisplayName)
		// Add more details to print as needed
	}
}
