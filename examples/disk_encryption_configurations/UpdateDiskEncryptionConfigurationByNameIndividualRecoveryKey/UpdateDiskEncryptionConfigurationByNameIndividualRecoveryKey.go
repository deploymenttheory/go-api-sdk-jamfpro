package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
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
