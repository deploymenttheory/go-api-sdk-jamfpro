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

	// Create a new Removable MAC Address
	newMACAddress := &jamfpro.ResourceRemovableMacAddress{
		Name: "E0:AC:CB:97:36:G4", // Replace with the actual MAC address name
		// ID: [set the ID if necessary]
	}

	createdMACAddress, err := client.CreateRemovableMACAddress(newMACAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Pretty print the created MAC address details in XML
	createdMACAddressXML, err := xml.MarshalIndent(createdMACAddress, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created MAC address data: %v", err)
	}
	fmt.Println("Created MAC Address Details:\n", string(createdMACAddressXML))
}
