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
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Create a search struct with updated details
	searchToUpdate := jamfpro.ResponseAdvancedMobileDeviceSearches{
		Name:   "Advanced Search Name",
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

	// Perform the update
	updatedSearch, err := client.UpdateAdvancedMobileDeviceSearchByID(123, &searchToUpdate) // Replace 123 with the actual ID
	if err != nil {
		log.Fatalf("Error updating advanced mobile device search by ID: %v", err)
	}

	// Output the updated search
	output, err := xml.MarshalIndent(updatedSearch, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling updated search to XML: %v", err)
	}
	fmt.Printf("Updated Advanced Mobile Device Search (ID: %d):\n%s\n", 123, string(output))
}
