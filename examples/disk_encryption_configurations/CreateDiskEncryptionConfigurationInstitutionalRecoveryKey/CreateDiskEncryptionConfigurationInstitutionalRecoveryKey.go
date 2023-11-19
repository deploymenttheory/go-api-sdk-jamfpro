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

	institutionalConfig := &jamfpro.DiskEncryptionConfiguration{
		Name:                  "Corporate Encryption",
		KeyType:               "Institutional",
		FileVaultEnabledUsers: "Management Account",
		InstitutionalRecoveryKey: &jamfpro.DiskEncryptionConfigurationDataSubsetInstitutionalRecoveryKey{
			Key:             "OID.2.5.4.13=admins-MacBook-Pro.local, CN=FileVault Recovery Key",
			CertificateType: "PKCS12",
			Password:        "PKCS12Password",
			Data:            "Base64EncodedCertificateData",
		},
	}

	createdConfig, err := client.CreateDiskEncryptionConfiguration(institutionalConfig)
	if err != nil {
		log.Fatalf("Error creating Institutional Key Configuration: %v", err)
	}

	configXML, err := xml.MarshalIndent(createdConfig, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created configuration to XML: %v", err)
	}

	fmt.Printf("Created Institutional Disk Encryption Configuration:\n%s\n", configXML)
}
