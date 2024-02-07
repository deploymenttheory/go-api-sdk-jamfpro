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

	// Define the application name and the subset you want to retrieve
	appName := "TextWrangler.app" // Replace with the actual application name
	subset := "General"           // Subset values can be General, Scope, SelfService, VPPCodes and VPP.

	// Call GetMacApplicationByNameAndDataSubset
	macApp, err := client.GetMacApplicationByNameAndDataSubset(appName, subset)
	if err != nil {
		log.Fatalf("Error fetching Mac Application by Name and Subset: %v", err)
	}

	// Pretty print the response in XML
	macAppXML, err := xml.MarshalIndent(macApp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Mac Application data: %v", err)
	}
	fmt.Println("Fetched Mac Application Data:\n", string(macAppXML))
}
