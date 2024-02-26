package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
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
