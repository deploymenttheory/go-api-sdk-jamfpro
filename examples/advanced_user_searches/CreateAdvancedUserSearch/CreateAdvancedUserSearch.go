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

	// Define the advanced user search to create
	newAdvancedUserSearch := &jamfpro.AdvancedUserSearch{
		Name: "Advanced User Search Name",
		Criteria: []jamfpro.AdvancedUserSearchCriteriaDetail{
			{
				Criterion: jamfpro.AdvancedUserSearchCriterionDetail{
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

	// Call CreateAdvancedUserSearch function
	createdSearch, err := client.CreateAdvancedUserSearch(newAdvancedUserSearch)
	if err != nil {
		log.Fatalf("Error creating advanced user search: %v", err)
	}

	// Pretty print the created advanced user search in XML
	createdSearchXML, err := xml.MarshalIndent(createdSearch, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created advanced user search data: %v", err)
	}
	fmt.Println("Created Advanced User Search:\n", string(createdSearchXML))
}
