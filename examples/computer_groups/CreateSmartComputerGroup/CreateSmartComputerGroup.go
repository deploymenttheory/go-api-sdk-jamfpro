package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Sample data for creating a new computer group (replace with actual data as needed)
	newSmartGroup := &jamfpro.ResponseComputerGroup{
		Name:    "Operating System Version like 15",
		IsSmart: true,
		Site: jamfpro.ComputerGroupSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: []jamfpro.ComputerGroupCriterion{
			{
				Name:        "Operating System Version",
				Priority:    0,
				AndOr:       "and",
				SearchType:  "like",
				SearchValue: "macOS 14",
				//OpeningParen: false,
				//ClosingParen: false,
			},
		},
	}

	// Call CreateComputerGroup function
	createdGroup, err := client.CreateComputerGroup(newSmartGroup)
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
