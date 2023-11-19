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
	updatedAdvancedUserSearch := &jamfpro.AdvancedUserSearch{
		Name: "Updated Advanced User Search Name",
		Criteria: []jamfpro.AdvancedUserSearchCriteriaDetail{
			{
				Criterion: jamfpro.AdvancedUserSearchCriterionDetail{
					Name:         "Email Address",
					Priority:     0,
					AndOr:        "and",
					SearchType:   "like",
					Value:        "updatedcompany.com",
					OpeningParen: false,
					ClosingParen: false,
				},
			},
		},
		DisplayFields: []jamfpro.AdvancedUserSearchSiteDisplayFieldDetail{
			{
				DisplayField: struct {
					Name string `xml:"name"`
				}{Name: "Email Address"},
			},
		},
		Site: jamfpro.AdvancedUserSearchSiteDetail{
			ID:   -1,
			Name: "None",
		},
	}

	// Update by Name
	updatedByName, err := client.UpdateAdvancedUserSearchByName("Original Search Name", updatedAdvancedUserSearch) // Replace "Original Search Name" with the actual name
	if err != nil {
		log.Fatalf("Error updating advanced user search by name: %v", err)
	}
	// Print updated search by Name
	updatedByNameXML, err := xml.MarshalIndent(updatedByName, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated advanced user search by name: %v", err)
	}
	fmt.Println("Updated Advanced User Search by Name:\n", string(updatedByNameXML))
}
