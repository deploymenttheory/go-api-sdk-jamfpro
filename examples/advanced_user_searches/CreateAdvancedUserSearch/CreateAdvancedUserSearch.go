package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the advanced user search to create
	newAdvancedUserSearch := &jamfpro.ResourceAdvancedUserSearch{
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
