package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the computer group name and details for update
	groupName := "NewGroupNameBySDKWithnoSiteset"
	groupUpdate := &jamfpro.ResponseComputerGroup{
		Name:    "UpdatedGroupNameBySDKWithnoSiteset",
		IsSmart: true,
		Site:    jamfpro.ComputerGroupSite{ID: -1, Name: "None"},
		Criteria: []jamfpro.CriterionContainer{
			{
				Size: 1,
				Criterion: jamfpro.ComputerGroupCriterion{
					Name:        "Last Inventory Update",
					Priority:    0,
					AndOr:       jamfpro.And,
					SearchType:  "more than x days ago",
					SearchValue: "7",
				},
			},
		},
	}

	// Call UpdateComputerGroupByName function
	updatedGroup, err := client.UpdateComputerGroupByName(groupName, groupUpdate)
	if err != nil {
		log.Fatalf("Error updating Computer Group by Name: %v", err)
	}

	// Pretty print the updated computer group in XML
	updatedGroupXML, err := xml.MarshalIndent(updatedGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated Computer Group data: %v", err)
	}
	fmt.Println("Updated Computer Group:\n", string(updatedGroupXML))
}
