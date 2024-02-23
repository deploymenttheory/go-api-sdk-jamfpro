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
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

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
			LogLevel:            loadedConfig.ClientOptions.LogLevel,
			LogOutputFormat:     loadedConfig.ClientOptions.LogOutputFormat,
			LogConsoleSeparator: loadedConfig.ClientOptions.LogConsoleSeparator,
			HideSensitiveData:   loadedConfig.ClientOptions.HideSensitiveData,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Assemble the request body for creating an account group
	accountGroup := &jamfpro.ResourceAccountGroup{
		Name:         "jamf sdk test group",
		AccessLevel:  "Full Access", // Full Access / Site Access
		PrivilegeSet: "Custom",      // Administrator / Auditor / Enrollment Only / Custom
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
		Members: jamfpro.AccountGroupSubsetMembers{
			{User: jamfpro.MemberUser{ID: 12, Name: "Barry White"}},
			{User: jamfpro.MemberUser{ID: 2, Name: "dafydd.watkins"}},
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
