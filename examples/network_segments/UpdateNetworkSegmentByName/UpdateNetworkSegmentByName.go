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

	// Define the new network segment
	updatedSegment := &jamfpro.ResourceNetworkSegment{
		Name:                "NY Office",
		StartingAddress:     "10.1.1.1",
		EndingAddress:       "10.10.1.1",
		OverrideBuildings:   false,
		OverrideDepartments: false,
	}

	segmentName := "NY Office" // Replace with actual name

	updated, err := client.UpdateNetworkSegmentByName(segmentName, updatedSegment)
	if err != nil {
		log.Fatalf("Error updating network segment by name: %v", err)
	}

	// Print the updated network segment details
	segmentXML, _ := xml.MarshalIndent(updated, "", "    ")
	fmt.Println("Updated Network Segment:", string(segmentXML))
}
