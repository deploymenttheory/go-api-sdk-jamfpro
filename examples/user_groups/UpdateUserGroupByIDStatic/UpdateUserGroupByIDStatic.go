package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Example user group to be updated
	updatedUserGroup := &jamfpro.ResourceUserGroup{
		Name:             "Static Group",
		IsSmart:          false,
		IsNotifyOnChange: false,
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Users: []jamfpro.UserGroupSubsetUserItem{
			{
				ID:           1938,
				Username:     "Mercy",
				EmailAddress: "mercy@company.com",
			},
			{
				ID:           1939,
				Username:     "Aaron",
				EmailAddress: "aaron@company.com",
			},
		},
	}

	// Replace with the actual ID of the user group you want to update
	userGroupID := 1

	// Call UpdateUserGroupByID to update the user group
	updatedGroup, err := client.UpdateUserGroupByID(userGroupID, updatedUserGroup)
	if err != nil {
		fmt.Println("Error updating user group:", err)
		return
	}

	// Pretty print the created user group details in XML
	createdUserGroupXML, err := xml.MarshalIndent(updatedGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created user group data: %v", err)
	}
	fmt.Println("Created User Group Details:\n", string(createdUserGroupXML))
}
