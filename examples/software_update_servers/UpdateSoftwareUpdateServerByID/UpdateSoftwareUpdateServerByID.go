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

	// Construct the update data
	serverToUpdate := &jamfpro.ResourceSoftwareUpdateServer{
		ID:            1, // Set the ID of the server to update
		Name:          "Updated New York SUS",
		IPAddress:     "10.10.51.249",
		Port:          8088,
		SetSystemWide: true,
	}

	// Call UpdateSoftwareUpdateServerByID
	updatedServer, err := client.UpdateSoftwareUpdateServerByID(serverToUpdate.ID, serverToUpdate)
	if err != nil {
		log.Fatalf("Error updating software update server by ID: %v", err)
	}

	// Pretty print the details in XML
	updatedServerXML, err := xml.MarshalIndent(updatedServer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated server data: %v", err)
	}
	fmt.Println("Updated Software Update Server Details:\n", string(updatedServerXML))
}
