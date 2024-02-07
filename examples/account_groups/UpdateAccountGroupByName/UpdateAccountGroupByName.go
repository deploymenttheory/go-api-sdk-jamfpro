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

	// Assemble the request body for creating an account group
	updatedAccountGroup := &jamfpro.ResourceAccountGroup{
		Name:         "Test Group",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Site: jamfpro.SharedResourceSite{
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
		Members: jamfpro.AccountGroupSubsetMembers{
			{User: jamfpro.MemberUser{ID: 12, Name: "Barry White"}},
			{User: jamfpro.MemberUser{ID: 2, Name: "dafydd.watkins"}},
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
