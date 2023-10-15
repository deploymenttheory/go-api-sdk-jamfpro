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

	// Assemble the request body for creating an account group
	accountGroup := &jamfpro.ResponseAccountGroup{
		Name:         "Test Group",
		AccessLevel:  "Full Access",
		PrivilegeSet: "Administrator",
		Site: jamfpro.AccountDataSubsetSite{
			ID:   -1,
			Name: "None",
		},
		Privileges: jamfpro.AccountDataSubsetPrivileges{
			JSSObjects:    []string{"string"},
			JSSSettings:   []string{"string"},
			JSSActions:    []string{"string"},
			Recon:         []string{"string"},
			CasperAdmin:   []string{"string"},
			CasperRemote:  []string{"string"},
			CasperImaging: []string{"string"},
		},
		Members: []jamfpro.AccountDataSubsetUser{
			{
				ID:   1,
				Name: "string",
			},
		},
	}

	// Call CreateAccountGroupByID function
	createdAccountGroup, err := client.CreateAccountGroupByID(accountGroup)

	if err != nil {
		log.Fatalf("Error creating account group by ID: %v", err)
	}

	// Pretty print the created account group details
	accountGroupXML, err := xml.MarshalIndent(createdAccountGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling account group data: %v", err)
	}
	fmt.Println("Created Account Group Details:", string(accountGroupXML))
}
