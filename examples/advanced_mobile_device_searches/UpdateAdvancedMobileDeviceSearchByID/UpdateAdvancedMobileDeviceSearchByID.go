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

	// ID of the search to update
	searchID := "94"
	siteID := "-1"

	// Create updated search
	updatedSearch := jamfpro.ResourceAdvancedMobileDeviceSearch{
		Name: "Updated Test Search",
		Criteria: []jamfpro.SharedSubsetCriteriaJamfProAPI{
			{
				Name:         "Identity",
				AndOr:        "and",
				SearchType:   "is",
				Value:        "test",
				OpeningParen: jamfpro.TruePtr(),
				ClosingParen: jamfpro.FalsePtr(),
			},
			{
				Name:         "Languages",
				Priority:     1,
				AndOr:        "and",
				SearchType:   "is",
				Value:        "test",
				OpeningParen: jamfpro.FalsePtr(),
				ClosingParen: jamfpro.TruePtr(),
			},
		},
		DisplayFields: []string{"App Name", "Device Name"},
		SiteId:        &siteID,
	}

	// Call Function
	updated, err := client.UpdateAdvancedMobileDeviceSearchByID(searchID, updatedSearch)
	if err != nil {
		log.Fatalf("Error updating advanced mobile device search: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(updated, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated search data: %v", err)
	}
	fmt.Println("Updated Advanced Mobile Device Search:\n", string(response))
}
