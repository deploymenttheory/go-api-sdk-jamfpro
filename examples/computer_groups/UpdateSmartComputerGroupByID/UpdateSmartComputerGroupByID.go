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

	// Define the computer group details for update
	groupUpdate := &jamfpro.ResourceComputerGroup{
		Name:    "Group Name",
		IsSmart: true,
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: jamfpro.SharedContainerCriteria{
			Criterion: []jamfpro.SharedSubsetCriteria{
				{
					Name:       "Operating System Version",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "macOS 14",
				},
			},
		},
		// Include other fields if necessary
	}

	// Call UpdateComputerGroupByID function
	updatedGroup, err := client.UpdateComputerGroupByID(123, groupUpdate) // Assuming 123 is the ID of the computer group you want to update
	if err != nil {
		log.Fatalf("Error updating Computer Group by ID: %v", err)
	}

	// Pretty print the updated computer group in XML
	updatedGroupXML, err := xml.MarshalIndent(updatedGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated Computer Group data: %v", err)
	}
	fmt.Println("Updated Computer Group:\n", string(updatedGroupXML))
}
