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

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Assemble the request body for creating an account
	accountDetail := &jamfpro.ResponseAccount{
		Name:                "John Smith",
		DirectoryUser:       false,
		FullName:            "John Smith",
		Email:               "john.smith@company.com",
		EmailAddress:        "john.smith@company.com",
		Enabled:             "Enabled",
		ForcePasswordChange: true,
		AccessLevel:         "Full Access",
		PrivilegeSet:        "Administrator",
		Password:            "sample",
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
	}

	// Call CreateAccountByID function
	createdAccount, err := client.CreateAccountByID(accountDetail)

	if err != nil {
		log.Fatalf("Error creating account by ID: %v", err)
	}

	// Pretty print the created account details
	accountXML, err := xml.MarshalIndent(createdAccount, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling account data: %v", err)
	}
	fmt.Println("Created Account Details:", string(accountXML))
}
