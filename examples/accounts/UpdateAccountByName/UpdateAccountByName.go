package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
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
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
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

	// Assemble the request body for updating an account
	accountToUpdate := &jamfpro.ResourceAccount{
		Name:                "John Smith updated by resource name",
		DirectoryUser:       false,
		FullName:            "John Smith Updated",
		Email:               "john.smith.updated@company.com",
		EmailAddress:        "john.smith.updated@company.com",
		Enabled:             "Enabled",
		ForcePasswordChange: true,
		AccessLevel:         "Full Access",
		PrivilegeSet:        "Administrator",
		Password:            "sampleUpdated",
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Privileges: jamfpro.AccountSubsetPrivileges{
			JSSObjects:    []string{"updatedString"},
			JSSSettings:   []string{"updatedString"},
			JSSActions:    []string{"updatedString"},
			Recon:         []string{"updatedString"},
			CasperAdmin:   []string{"updatedString"},
			CasperRemote:  []string{"updatedString"},
			CasperImaging: []string{"updatedString"},
		},
	}

	// Let's assume we are updating an account with the name "Bobby".
	accountName := "John Smith Updated2"

	// Call UpdateAccountByName function
	updatedAccount, err := client.UpdateAccountByName(accountName, accountToUpdate)

	if err != nil {
		log.Fatalf("Error updating account by name: %v", err)
	}

	// Pretty print the updated account details
	accountXML, err := xml.MarshalIndent(updatedAccount, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling account data: %v", err)
	}
	fmt.Println("Updated Account Details:", string(accountXML))
}
