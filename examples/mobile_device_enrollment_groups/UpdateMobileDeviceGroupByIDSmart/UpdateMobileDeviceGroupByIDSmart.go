package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
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
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
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

	// Define the ID of the group you want to update
	groupID := 123 // Replace with the actual group ID

	// Define the updated group data
	updatedSmartGroup := &jamfpro.ResourceMobileDeviceGroup{
		Name:    "Sample Smart Group",
		IsSmart: true,
		Criteria: jamfpro.SharedContainerCriteria{
			Size: 3, // The number of criteria
			Criterion: []jamfpro.SharedSubsetCriteria{
				{
					Name:         "Last Inventory Update",
					Priority:     0,
					AndOr:        "AND",
					SearchType:   "more than x days ago",
					Value:        "7",
					OpeningParen: true,
				},
				{
					Name:         "Department",
					Priority:     1,
					AndOr:        "and",
					SearchType:   "is",
					Value:        "marketing",
					ClosingParen: true,
				},
				{
					Name:         "Building",
					Priority:     2,
					AndOr:        "or",
					SearchType:   "is",
					Value:        "london wall",
					OpeningParen: true,
					ClosingParen: true,
				},
			},
		},
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		// other fields if necessary
	}

	// Call the UpdateMobileDeviceGroupByID function
	updatedGroup, err := client.UpdateMobileDeviceGroupByID(groupID, updatedSmartGroup)
	if err != nil {
		log.Fatalf("Error updating mobile device group: %s\n", err)
	}

	// Marshal and print the updated group
	updatedGroupXML, err := xml.MarshalIndent(updatedGroup, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated group data: %v", err)
	}
	fmt.Println("Updated Mobile Device Group:\n", string(updatedGroupXML))
}
