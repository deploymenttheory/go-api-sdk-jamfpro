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

	// Define the computers for the static group
	computers := []jamfpro.ComputerGroupSubsetComputer{
		{
			ID: 21,
		},
		{
			ID: 16,
		},
	}

	// Create a new static computer group
	newStaticGroup := &jamfpro.ResourceComputerGroup{
		Name:    "jamfpro-go-sdk-test-static-group",
		IsSmart: false,
		Site: &jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},

		Computers: &computers,
	}

	// Call CreateComputerGroup function
	createdGroup, err := client.CreateComputerGroup(newStaticGroup)
	if err != nil {
		log.Fatalf("Error creating Computer Group: %v", err)
	}

	// Pretty print the created group in XML
	groupXML, err := xml.MarshalIndent(createdGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Created Computer Group:\n", string(groupXML))
}
