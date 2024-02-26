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

	profileID := 171

	// Call GetMacOSConfigurationProfileByID function
	profile, err := client.GetMacOSConfigurationProfileByID(profileID)
	if err != nil {
		log.Fatalf("Error fetching macOS Configuration Profile by ID: %v", err)
	}

	// Pretty print the profile in XML
	profileXML, err := xml.MarshalIndent(profile, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling macOS Configuration Profile data: %v", err)
	}
	fmt.Println("Fetched macOS Configuration Profile:\n", string(profileXML))
}
