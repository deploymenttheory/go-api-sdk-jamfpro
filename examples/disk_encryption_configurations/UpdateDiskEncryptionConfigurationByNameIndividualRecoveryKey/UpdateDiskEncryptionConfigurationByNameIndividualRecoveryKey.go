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
	configToUpdate := &jamfpro.ResourceDiskEncryptionConfiguration{
		Name:                  "Corporate Encryption Name Updated",
		KeyType:               "Individual",
		FileVaultEnabledUsers: "Management Account",
	}

	updatedConfig, err := client.UpdateDiskEncryptionConfigurationByName("Corporate Encryption", configToUpdate)
	if err != nil {
		log.Fatalf("Error updating Disk Encryption Configuration by Name: %v", err)
	}

	configXML, err := xml.MarshalIndent(updatedConfig, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated configuration to XML: %v", err)
	}

	fmt.Printf("Updated Disk Encryption Configuration by Name:\n%s\n", configXML)
}
