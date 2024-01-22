package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
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

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the updated advanced user search details
	updatedAdvancedUserSearch := &jamfpro.ResourceAdvancedUserSearch{
		Name: "Advanced User Search Name by jamf pro sdk",
		Criteria: jamfpro.SharedContainerCriteria{
			Criterion: []jamfpro.SharedSubsetCriteria{
				{
					Name:         "Email Address",
					Priority:     0,
					AndOr:        "and",
					SearchType:   "like",
					Value:        "company.com",
					OpeningParen: false,
					ClosingParen: false,
				},
			},
		},
		DisplayFields: []jamfpro.SharedAdvancedSearchContainerDisplayField{
			{
				DisplayField: []jamfpro.SharedAdvancedSearchSubsetDisplayField{
					{
						Name: "Computers",
					},
					{
						Name: "Content Name",
					},
					{
						Name: "Roster Course Source",
					},
					// Additional display fields can be added here
				},
			},
		},
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	searchName := "Advanced User Search Name by jamf pro sdk" // Replace with actual search name

	// Update by ID
	updatedByID, err := client.UpdateAdvancedUserSearchByName(searchName, updatedAdvancedUserSearch)
	if err != nil {
		log.Fatalf("Error updating advanced user search by ID: %v", err)
	}
	// Print updated search by ID
	updatedByIDXML, err := xml.MarshalIndent(updatedByID, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated advanced user search by ID: %v", err)
	}
	fmt.Println("Updated Advanced User Search by ID:\n", string(updatedByIDXML))
}
