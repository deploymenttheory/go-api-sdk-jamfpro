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

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the ID for the file extension you want to fetch
	fileExtensionID := 142 // Replace with the desired ID

	// Call GetAllowedFileExtensionByID function
	allowedExtension, err := client.GetAllowedFileExtensionByID(fileExtensionID)
	if err != nil {
		log.Fatalf("Error fetching allowed file extension by ID: %v", err)
	}

	// Pretty print the fetched file extension in XML
	allowedExtensionXML, err := xml.MarshalIndent(allowedExtension, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling allowed file extension data: %v", err)
	}
	fmt.Println("Fetched Allowed File Extension by ID:\n", string(allowedExtensionXML))
}
