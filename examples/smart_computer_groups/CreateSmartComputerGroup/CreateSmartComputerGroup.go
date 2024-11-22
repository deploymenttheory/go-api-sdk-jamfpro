package main

import (
	"encoding/json"
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

	// Create new smart computer group
	siteID := "-1"
	newGroup := jamfpro.ResourceSmartComputerGroup{
		Name: "Test Smart Group",
		Criteria: []jamfpro.SharedSubsetCriteriaJamfProAPI{
			{
				Name:         "Application Title",
				Priority:     0,
				AndOr:        "and",
				SearchType:   "is",
				Value:        "test",
				OpeningParen: false,
				ClosingParen: false,
			},
		},
		SiteId: &siteID,
	}

	// Call function
	created, err := client.CreateSmartComputerGroup(newGroup)
	if err != nil {
		log.Fatalf("Error creating smart computer group: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(created, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created group data: %v", err)
	}
	fmt.Println("Created Smart Computer Group:\n", string(response))
}
