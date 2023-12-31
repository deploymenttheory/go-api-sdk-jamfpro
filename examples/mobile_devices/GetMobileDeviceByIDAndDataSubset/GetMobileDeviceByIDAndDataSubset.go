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
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelInfo // Adjust log level as needed

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

	// Example device ID and subset
	deviceID := 1       // Replace with an actual device ID
	subset := "General" // Replace with the desired subset

	// Get mobile device by ID and subset
	deviceSubset, err := client.GetMobileDeviceByIDAndDataSubset(deviceID, subset)
	if err != nil {
		log.Fatalf("Error fetching mobile device by ID and subset: %v", err)
	}

	// Pretty print the device subset data in XML
	deviceSubsetXML, err := xml.MarshalIndent(deviceSubset, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling device subset data: %v", err)
	}
	fmt.Println("Device Subset Data:\n", string(deviceSubsetXML))
}
