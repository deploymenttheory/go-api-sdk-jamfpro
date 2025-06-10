package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example user group to be created
	newUserGroup := &jamfpro.ResourceUserGroup{
		Name:             "Teachers",
		IsSmart:          true,
		IsNotifyOnChange: true,
		Site: &jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: []jamfpro.SharedSubsetCriteria{
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
	}

	// Call CreateUserGroup to create a new user group
	createdUserGroup, err := client.CreateUserGroup(newUserGroup)
	if err != nil {
		fmt.Println("Error creating user group:", err)
		return
	}

	// Pretty print the created user group details in XML
	createdUserGroupXML, err := xml.MarshalIndent(createdUserGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created user group data: %v", err)
	}
	fmt.Println("Created User Group Details:\n", string(createdUserGroupXML))
}
