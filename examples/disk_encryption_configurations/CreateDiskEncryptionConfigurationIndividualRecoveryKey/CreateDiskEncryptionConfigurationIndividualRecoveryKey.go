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
	individualConfig := &jamfpro.ResourceDiskEncryptionConfiguration{
		Name:                  "jamfpro-sdk-example-IndividualRecoveryKey-config",
		KeyType:               "Individual",
		FileVaultEnabledUsers: "Management Account",
	}

	createdConfig, err := client.CreateDiskEncryptionConfiguration(individualConfig)
	if err != nil {
		log.Fatalf("Error creating Individual Key Configuration: %v", err)
	}

	configXML, err := xml.MarshalIndent(createdConfig, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created configuration to XML: %v", err)
	}

	fmt.Printf("Created Individual Disk Encryption Configuration:\n%s\n", configXML)
}
