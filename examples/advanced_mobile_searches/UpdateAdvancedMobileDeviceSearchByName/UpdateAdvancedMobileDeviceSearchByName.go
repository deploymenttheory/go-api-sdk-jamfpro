package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for Jamf Pro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new Jamf Pro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Create a search struct with updated details
	searchToUpdate := jamfpro.ResponseAdvancedMobileDeviceSearches{
		Name:   "Advanced Search Name - Rename",
		ViewAs: "Standard Web Page",
		Criteria: []jamfpro.AdvancedMobileDeviceSearchesCriteria{
			{
				Size: 1,
				Criterion: jamfpro.Criterion{
					Name:         "Last Inventory Update",
					Priority:     0,
					AndOr:        "and",
					SearchType:   "more than x days ago",
					Value:        7,
					OpeningParen: false,
					ClosingParen: false,
				},
			},
		},
		DisplayFields: []jamfpro.AdvancedMobileDeviceSearchesDisplayField{
			{
				Size: 1,
				DisplayField: jamfpro.AdvancedMobileDeviceSearchesDisplayFieldItem{
					Name: "IP Address",
				},
			},
		},
		Site: jamfpro.AdvancedMobileDeviceSearchesSite{
			ID:   -1,
			Name: "None",
		},
	}

	// Use the struct we created above for the update
	updatedSearch, err := client.UpdateAdvancedMobileDeviceSearchByName("Advanced Search Name", &searchToUpdate) // Replace with the actual search name
	if err != nil {
		log.Fatalf("Error updating advanced mobile device search by name: %v", err)
	}

	// Output the updated search
	output, err := xml.MarshalIndent(updatedSearch, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling updated search to XML: %v", err)
	}
	fmt.Printf("Updated Advanced Mobile Device Search (Name: %s):\n%s\n", "Advanced Search Name", string(output))
}
