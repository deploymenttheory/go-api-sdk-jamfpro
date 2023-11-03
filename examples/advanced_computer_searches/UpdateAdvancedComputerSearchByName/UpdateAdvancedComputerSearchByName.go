package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// Define the name of the advanced computer search
const advancedComputerSearchName = "Advanced Search Name" // Replace with the actual name

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

	updatedSearch, err := client.UpdateAdvancedComputerSearchByName(advancedComputerSearchName, &jamfpro.ResponseAdvancedComputerSearch{
		Name:   "Advanced Search Name Updated",
		ViewAs: "Standard Web Page",
		Criteria: []jamfpro.AdvancedComputerSearchesCriteria{
			{
				Size: 1,
				Criterion: jamfpro.CriterionDetail{
					Name:         "Last Inventory Update",
					Priority:     0,
					AndOr:        "and",
					SearchType:   "more than x days ago",
					Value:        "7",
					OpeningParen: false,
					ClosingParen: false,
				},
			},
		},
		DisplayFields: []jamfpro.AdvancedComputerSearchesDisplayField{
			{
				Size: 1,
				DisplayField: jamfpro.DisplayFieldDetail{
					Name: "IP Address",
				},
			},
		},
		Site: jamfpro.AdvancedComputerSearchesSiteDetail{
			ID:   -1,
			Name: "None",
		},
	})
	if err != nil {
		fmt.Println("Error updating advanced computer search by name:", err)
		return
	}

	fmt.Println("Updated advanced computer Search object:", updatedSearch)
}
