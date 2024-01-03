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

	// Assemble the request body for updating an account
	accountDetail := &jamfpro.ResourceAccount{
		Name:                "John Smith Updated by resource ID",
		DirectoryUser:       false,
		FullName:            "John Smith Updated",
		Email:               "john.smith.updated@company.com",
		EmailAddress:        "john.smith.updated@company.com",
		Enabled:             "Enabled",
		ForcePasswordChange: true,
		AccessLevel:         "Full Access",
		PrivilegeSet:        "Custom", // Administrator / Auditor / Enrollment Only / Custom
		Password:            "sampleUpdated",
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Privileges: jamfpro.AccountSubsetPrivileges{
			JSSObjects:    []string{"Advanced Computer Searches"},
			JSSSettings:   []string{"updatedString"},
			JSSActions:    []string{"updatedString"},
			Recon:         []string{"updatedString"},
			CasperAdmin:   []string{"updatedString"},
			CasperRemote:  []string{"updatedString"},
			CasperImaging: []string{"updatedString"},
		},
	}

	// Let's assume we are updating an account with ID 1.
	accountID := 10

	// Call UpdateAccountByID function
	updatedAccount, err := client.UpdateAccountByID(accountID, accountDetail)

	if err != nil {
		log.Fatalf("Error updating account by ID: %v", err)
	}

	// Pretty print the updated account details
	accountXML, err := xml.MarshalIndent(updatedAccount, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling account data: %v", err)
	}
	fmt.Println("Updated Account Details:", string(accountXML))
}
