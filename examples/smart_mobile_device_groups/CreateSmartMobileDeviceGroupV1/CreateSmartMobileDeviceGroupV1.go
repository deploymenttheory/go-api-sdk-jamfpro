package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create new smart mobile device group
	siteID := "-1"
	newGroup := jamfpro.ResourceSmartMobileDeviceGroupV1{
		GroupName:        "Operating System Version like 18",
		GroupDescription: "This is a description",
		Criteria: []jamfpro.SharedSubsetCriteriaJamfProAPI{
			{
				Name:         "OS Version",
				Priority:     0,
				AndOr:        "and",
				SearchType:   "like",
				Value:        "18",
				OpeningParen: jamfpro.TruePtr(),
				ClosingParen: jamfpro.FalsePtr(),
			},
			{
				Name:         "Model",
				Priority:     1,
				AndOr:        "and",
				SearchType:   "like",
				Value:        "iPad",
				OpeningParen: jamfpro.FalsePtr(),
				ClosingParen: jamfpro.TruePtr(),
			},
		},
		SiteId: &siteID,
	}

	// Call function
	created, err := client.CreateSmartMobileDeviceGroupV1(newGroup)
	if err != nil {
		log.Fatalf("Error creating smart mobile device group: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(created, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created group data: %v", err)
	}
	fmt.Println("Created Smart Mobile Device Group:\n", string(response))
}
