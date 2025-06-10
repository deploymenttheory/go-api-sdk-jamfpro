package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
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
