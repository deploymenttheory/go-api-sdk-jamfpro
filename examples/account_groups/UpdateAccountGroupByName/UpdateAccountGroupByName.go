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

	// Assemble the request body for creating an account group
	updatedAccountGroup := &jamfpro.ResourceAccountGroup{
		Name:         "Test Group",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Site: &jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Privileges: jamfpro.AccountSubsetPrivileges{
			JSSObjects:    []string{"string"},
			JSSSettings:   []string{"string"},
			JSSActions:    []string{"string"},
			Recon:         []string{"string"},
			CasperAdmin:   []string{"string"},
			CasperRemote:  []string{"string"},
			CasperImaging: []string{"string"},
		},
		Members: []jamfpro.MemberUser{
			{ID: 12, Name: "Barry White"},
			{ID: 2, Name: "dafydd.watkins"},
		},
	}

	// Let's assume we are updating a group with the name "Test Group".
	groupName := "Test Group"

	// Call UpdateAccountGroupByName function
	updatedGroup, err := client.UpdateAccountGroupByName(groupName, updatedAccountGroup)

	if err != nil {
		log.Fatalf("Error updating account group by name: %v", err)
	}

	// Pretty print the updated group details
	groupXML, err := xml.MarshalIndent(updatedGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling group data: %v", err)
	}
	fmt.Println("Updated Group Details:", string(groupXML))
}
