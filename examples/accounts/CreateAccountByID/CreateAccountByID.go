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
	logLevel := http_client.LogLevelInfo // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

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

	// Assemble the request body for creating an account
	accountDetail := &jamfpro.ResourceAccount{
		Name:                "John James",
		DirectoryUser:       false,
		FullName:            "John James",
		Email:               "john.James@company.com",
		EmailAddress:        "john.James@company.com",
		Enabled:             "Enabled",
		ForcePasswordChange: true,
		AccessLevel:         "Full Access",
		PrivilegeSet:        "Administrator",
		Password:            "this is a really secure password 390423049823409894382092348092348",
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
	}

	// Call CreateAccountByID function
	createdAccount, err := client.CreateAccount(accountDetail)

	if err != nil {
		log.Fatal(err)
	}

	// Pretty print the created account details
	accountXML, err := xml.MarshalIndent(createdAccount, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling account data: %v", err)
	}
	fmt.Println("Created Account Details:", string(accountXML))
}
