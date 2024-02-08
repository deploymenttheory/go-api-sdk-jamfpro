package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
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
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
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
