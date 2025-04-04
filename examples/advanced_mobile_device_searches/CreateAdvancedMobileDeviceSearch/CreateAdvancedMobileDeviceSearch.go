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

	// Create a new search
	siteID := "-1"
	newSearch := jamfpro.ResourceAdvancedMobileDeviceSearch{
		Name: "Test Search 2",
		Criteria: []jamfpro.SharedSubsetCriteriaJamfProAPI{
			{
				Name:         "Building",
				Priority:     0,
				AndOr:        "and",
				SearchType:   "is",
				Value:        "test",
				OpeningParen: jamfpro.TruePtr(),
				ClosingParen: jamfpro.FalsePtr(),
			},
			{
				Name:         "iTunes Store Account",
				Priority:     1,
				AndOr:        "and",
				SearchType:   "is",
				Value:        "test",
				OpeningParen: jamfpro.FalsePtr(),
				ClosingParen: jamfpro.TruePtr(),
			},
		},
		DisplayFields: []string{"Building", "iTunes Store Account"},
		SiteId:        &siteID,
	}

	// Call Function
	created, err := client.CreateAdvancedMobileDeviceSearch(newSearch)
	if err != nil {
		log.Fatalf("Error creating advanced mobile device search: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(created, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created search data: %v", err)
	}
	fmt.Println("Created Advanced Mobile Device Search:\n", string(response))
}
