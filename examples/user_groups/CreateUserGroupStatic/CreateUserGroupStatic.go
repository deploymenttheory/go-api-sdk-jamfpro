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
	// Example user group to be created
	newUserGroup := &jamfpro.ResourceUserGroup{
		Name:             "Static Group",
		IsSmart:          false,
		IsNotifyOnChange: false,
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Users: []jamfpro.UserGroupSubsetUserItem{
			{
				ID:           1938,
				Username:     "Mercy",
				EmailAddress: "mercy@company.com",
			},
			{
				ID:           1939,
				Username:     "Aaron",
				EmailAddress: "aaron@company.com",
			},
		},
	}

	// Call CreateUserGroup to create a new user group
	createdUserGroup, err := client.CreateUserGroup(newUserGroup)
	if err != nil {
		fmt.Println("Error creating user group:", err)
		return
	}

	// Pretty print the created user group details in XML
	createdUserGroupXML, err := xml.MarshalIndent(createdUserGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created user group data: %v", err)
	}
	fmt.Println("Created User Group Details:\n", string(createdUserGroupXML))
}
