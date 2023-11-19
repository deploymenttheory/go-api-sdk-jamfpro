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

	// Define the advanced mobile device search request based on the provided XML
	search := jamfpro.ResponseAdvancedMobileDeviceSearches{
		Name:   "Advanced Search Name",
		ViewAs: "Standard Web Page",
		Criteria: []jamfpro.AdvancedMobileDeviceSearchesCriteria{
			{
				Criterion: jamfpro.Criterion{
					Name:       "Last Inventory Update",
					Priority:   0,
					AndOr:      "and",
					SearchType: "more than x days ago",
					Value:      7,
				},
			},
		},
		DisplayFields: []jamfpro.AdvancedMobileDeviceSearchesDisplayField{
			{
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

	// Create the new advanced mobile device search by ID
	createdSearch, err := client.CreateAdvancedMobileDeviceSearchByID(1, &search)
	if err != nil {
		log.Fatalf("Failed to create advanced mobile device search by ID: %v", err)
	}

	// Convert the created search into pretty XML for printing
	output, err := xml.MarshalIndent(createdSearch, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling created search to XML: %v", err)
	}

	// Print the pretty XML
	fmt.Printf("Created Advanced Mobile Device Search:\n%s\n", string(output))
}
