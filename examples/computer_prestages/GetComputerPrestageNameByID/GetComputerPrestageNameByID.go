package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	prestageName = "pse-ade_lbgstaging_Jamf_Connect_New_Config-1.1-0000" // Replace with the actual prestage name you want to fetch
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

	// Create a new jamfpro client instance,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Call GetComputerPrestageByNameByID function
	prestage, err := client.GetComputerPrestageByNameByID(prestageName)
	if err != nil {
		log.Fatalf("Error fetching Jamf computer prestage by name: %v", err)
	}

	// Pretty print the prestage in JSON
	prestageJSON, err := json.MarshalIndent(prestage, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf computer prestage data: %v", err)
	}
	fmt.Println("Fetched Jamf computer prestage:\n", string(prestageJSON))
}
