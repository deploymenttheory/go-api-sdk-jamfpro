package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
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

	// The name of the advanced mobile device search you want to retrieve
	searchName := "Advanced Search Name" // Replace with the actual name you want to retrieve

	// Call the GetAdvancedMobileDeviceSearchByName function
	search, err := client.GetAdvancedMobileDeviceSearchByName(searchName)
	if err != nil {
		log.Fatalf("Error fetching advanced mobile device search by name: %v", err)
	}

	// Convert the response into pretty XML for printing
	output, err := xml.MarshalIndent(search, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling search to XML: %v", err)
	}

	// Print the pretty XML
	fmt.Printf("Advanced Mobile Device Search (Name: %s):\n%s\n", searchName, string(output))
}
