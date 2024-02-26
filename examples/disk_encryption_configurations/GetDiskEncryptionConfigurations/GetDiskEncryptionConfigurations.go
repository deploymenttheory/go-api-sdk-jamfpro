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

	// Call GetDiskEncryptionConfigurations function
	configurations, err := client.GetDiskEncryptionConfigurations()
	if err != nil {
		log.Fatalf("Error fetching disk encryption configurations: %v", err)
	}

	// Pretty print the configurations in XML
	configurationsXML, err := xml.MarshalIndent(configurations, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling configurations data: %v", err)
	}
	fmt.Println("Fetched Disk Encryption Configurations:\n", string(configurationsXML))
}
