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

	// Let's assume you want to get the disk encryption configuration with ID 1
	configID := 36

	configuration, err := client.GetDiskEncryptionConfigurationByID(configID)
	if err != nil {
		log.Fatalf("Error fetching disk encryption configuration by ID: %v", err)
	}

	// Print the configuration in a pretty XML format (assuming the response is XML)
	configXML, err := xml.MarshalIndent(configuration, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling configuration data: %v", err)
	}
	fmt.Printf("Fetched Disk Encryption Configuration by ID:\n%s\n", configXML)
}
