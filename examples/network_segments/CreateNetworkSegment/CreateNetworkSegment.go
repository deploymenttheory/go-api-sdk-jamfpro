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

	// Define the new network segment
	newSegment := &jamfpro.ResourceNetworkSegment{
		Name:                "NY Office",
		StartingAddress:     "10.1.1.1",
		EndingAddress:       "10.10.1.1",
		OverrideBuildings:   false,
		OverrideDepartments: false,
	}

	// Create the network segment
	createdSegment, err := client.CreateNetworkSegment(newSegment)
	if err != nil {
		log.Fatalf("Error creating network segment: %v", err)
	}

	// Pretty print the created network segment details in XML
	segmentXML, err := xml.MarshalIndent(createdSegment, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling network segment details: %v", err)
	}
	fmt.Println("Created Network Segment Details:\n", string(segmentXML))
}
