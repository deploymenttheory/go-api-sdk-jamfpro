package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	configToUpdate := &jamfpro.DiskEncryptionConfiguration{
		Name:                  "Corporate Encryption",
		KeyType:               "Individual",
		FileVaultEnabledUsers: "Management Account",
	}

	updatedConfig, err := client.UpdateDiskEncryptionConfigurationByID(1, configToUpdate)
	if err != nil {
		log.Fatalf("Error updating Disk Encryption Configuration by ID: %v", err)
	}

	configXML, err := xml.MarshalIndent(updatedConfig, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated configuration to XML: %v", err)
	}

	fmt.Printf("Updated Disk Encryption Configuration by ID:\n%s\n", configXML)
}
