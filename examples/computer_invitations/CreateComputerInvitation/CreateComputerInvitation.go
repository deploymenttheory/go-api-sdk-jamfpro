package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

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

	// Construct the invitation object based on the example XML provided
	invitation := &jamfpro.ResourceComputerInvitation{
		InvitationType:              "DEFAULT",
		ExpirationDate:              "2024-12-07 11:13:35",
		SSHUsername:                 "jamfadmin",
		SSHPassword:                 "accountpassword",
		MultipleUsersAllowed:        false,
		CreateAccountIfDoesNotExist: false,
		HideAccount:                 false,
		LockDownSSH:                 false,
		EnrollIntoSite: jamfpro.ComputerInvitationSubsetEnrollIntoState{
			ID:   -1,
			Name: "None",
		},
		KeepExistingSiteMembership: false,
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	// Call the CreateComputerInvitation function with the constructed invitation
	createdInvitation, err := client.CreateComputerInvitation(invitation)
	if err != nil {
		fmt.Printf("Error creating computer invitation: %s\n", err)
	} else {
		fmt.Printf("Created computer invitation: %+v\n", createdInvitation)
	}
}
