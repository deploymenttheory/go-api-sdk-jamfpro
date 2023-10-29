package main

import (
	"encoding/xml"
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

	// Define the ID of the Computer Prestage you want to fetch as a string
	prestageID := "2" // Replace with the actual prestage ID as a string

	// Call GetComputerPrestageByID function
	prestage, err := client.GetComputerPrestageByID(prestageID)
	if err != nil {
		log.Fatalf("Error fetching Computer Prestage by ID: %v", err)
	}

	// Pretty print the prestage in XML
	prestageXML, err := xml.MarshalIndent(prestage, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Prestage data: %v", err)
	}
	fmt.Println("Fetched Computer Prestage:\n", string(prestageXML))
}
