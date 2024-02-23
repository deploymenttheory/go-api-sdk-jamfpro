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
	// Define the computers for the static group
	computers := []jamfpro.ComputerGroupSubsetComputer{
		{
			ID:            2,
			Name:          "MacBook Pro",
			MacAddress:    "",
			AltMacAddress: "",
			SerialNumber:  "D2FHXH22QB",
		},
		{
			ID:            6,
			Name:          "MacBook Pro",
			MacAddress:    "",
			AltMacAddress: "",
			SerialNumber:  "LT6M4DTF88",
		},
	}

	// Create a new static computer group
	newStaticGroup := &jamfpro.ResourceComputerGroup{
		Name:      "SDK Static Group Test",
		IsSmart:   false,
		Site:      jamfpro.SharedResourceSite{ID: -1, Name: "None"},
		Computers: computers,
	}

	// Call CreateComputerGroup function
	createdGroup, err := client.CreateComputerGroup(newStaticGroup)
	if err != nil {
		log.Fatalf("Error creating Computer Group: %v", err)
	}

	// Pretty print the created group in XML
	groupXML, err := xml.MarshalIndent(createdGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Created Computer Group:\n", string(groupXML))
}
