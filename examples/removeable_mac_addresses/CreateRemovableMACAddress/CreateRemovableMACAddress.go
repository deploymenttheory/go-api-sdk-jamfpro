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

	// Create a new Removable MAC Address
	newMACAddress := &jamfpro.ResourceRemovableMacAddress{
		Name: "E0:AC:CB:97:36:G4", // Replace with the actual MAC address name
		// ID: [set the ID if necessary]
	}

	createdMACAddress, err := client.CreateRemovableMACAddress(newMACAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Pretty print the created MAC address details in XML
	createdMACAddressXML, err := xml.MarshalIndent(createdMACAddress, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created MAC address data: %v", err)
	}
	fmt.Println("Created MAC Address Details:\n", string(createdMACAddressXML))
}
