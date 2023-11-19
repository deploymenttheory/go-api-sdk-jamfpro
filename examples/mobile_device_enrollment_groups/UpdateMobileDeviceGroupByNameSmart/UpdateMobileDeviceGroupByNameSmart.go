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

	// Define the Name of the group you want to update
	groupName := "Smart-Group-Name" // Replace with the actual group Name

	// Define the updated group data
	updatedSmartGroup := &jamfpro.ResponseMobileDeviceGroup{
		Name:    "Sample Smart Group",
		IsSmart: true,
		Criteria: []jamfpro.MobileDeviceGroupCriteriaItem{
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
		Site: jamfpro.MobileDeviceGroupSite{
			ID:   -1,
			Name: "None",
		},
	}

	// Call the UpdateMobileDeviceGroupByName function
	updatedGroup, err := client.UpdateMobileDeviceGroupByName(groupName, updatedSmartGroup)
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
