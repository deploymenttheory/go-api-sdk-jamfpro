package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
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
