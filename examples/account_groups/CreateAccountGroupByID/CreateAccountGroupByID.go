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

	// Assemble the request body for creating an account group
	accountGroup := &jamfpro.ResourceAccountGroup{
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
			{ID: 7, Name: "John  Doe"},
			{ID: 2, Name: "dafydd.watkins"},
		},
	}

	// Call CreateAccountGroupByID function
	createdAccountGroup, err := client.CreateAccountGroup(accountGroup)

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
