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

	// Example display name
	patchSoftwareTitleConfigurationName := "1Password 7" // Replace with an actual device name

	// Get patch software title configuration by name
	patchTitle, err := client.GetPatchSoftwareTitleConfigurationByName(patchSoftwareTitleConfigurationName)
	if err != nil {
		log.Fatalf("Error fetching patch software title configuration by name: %v", err)
	}

	// Pretty print the network segments in XML
	mobileDeviceXML, err := xml.MarshalIndent(patchTitle, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling network segments data: %v", err)
	}
	fmt.Println("Network Segments:\n", string(mobileDeviceXML))
}
