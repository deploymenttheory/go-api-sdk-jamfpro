package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	//logLevel := logger.LogLevelInfo // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Assemble the request body for creating an account
	accountDetail := &jamfpro.ResourceAccount{
		Name:                "Barry White",
		DirectoryUser:       false,
		FullName:            "Barry White",
		Email:               "Barry.White@company.com",
		EmailAddress:        "Barry.White@company.com",
		Enabled:             "Enabled",
		ForcePasswordChange: true,
		AccessLevel:         "Full Access", // Full Access / Site Access
		PrivilegeSet:        "Custom",      // Administrator / Auditor / Enrollment Only / Custom
		Password:            "this is a really secure password 390423049823409894382092348092348",
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Privileges: jamfpro.AccountSubsetPrivileges{
			JSSObjects:    []string{"Update Webhooks", "Delete Webhooks"},
			JSSSettings:   []string{"Read SSO Settings", "Update User-Initiated Enrollment"},
			JSSActions:    []string{"Send Computer Bluetooth Command", "Computer Delete User Account Command"},
			Recon:         []string{"string"},
			CasperAdmin:   []string{"Use Casper Admin", "Save With Casper Admin"},
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
