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
	// Example payload for creating a new software update server
	newServer := &jamfpro.ResourceSoftwareUpdateServer{
		Name:          "New York SUS",
		IPAddress:     "10.10.51.248",
		Port:          8088,
		SetSystemWide: true,
	}

	// Call CreateSoftwareUpdateServer
	createdServer, err := client.CreateSoftwareUpdateServer(newServer)
	if err != nil {
		log.Fatalf("Error creating software update server: %v", err)
	}

	// Pretty print the created server details in XML
	createdServerXML, err := xml.MarshalIndent(createdServer, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created server data: %v", err)
	}
	fmt.Println("Created Software Update Server Details:\n", string(createdServerXML))
}
