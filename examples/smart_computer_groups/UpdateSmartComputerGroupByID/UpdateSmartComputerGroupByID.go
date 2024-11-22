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

	// Define group ID to update
	groupID := "1"
	siteID := "-1"

	// Create update data
	updateGroup := jamfpro.ResourceSmartComputerGroup{
		Name: "Updated Smart Group",
		Criteria: []jamfpro.SharedSubsetCriteriaJamfProAPI{
			{
				Name:         "Account",
				Priority:     0,
				AndOr:        "and",
				SearchType:   "is",
				Value:        "test",
				OpeningParen: jamfpro.FalsePtr(),
				ClosingParen: jamfpro.FalsePtr(),
			},
		},
		SiteId: &siteID,
	}

	// Call function
	updated, err := client.UpdateSmartComputerGroupByID(groupID, updateGroup)
	if err != nil {
		log.Fatalf("Error updating smart computer group: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(updated, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated group data: %v", err)
	}
	fmt.Println("Updated Smart Computer Group:\n", string(response))
}
