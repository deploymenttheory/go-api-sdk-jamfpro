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
	id := 1 // Replace with your Mac application ID
	macApp, err := client.GetMacApplicationByID(id)
	if err != nil {
		log.Fatalf("Error fetching Mac Application by ID: %v", err)
	}

	// Pretty print the profile in XML
	profileXML, err := xml.MarshalIndent(macApp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling VPP mac Application data: %v", err)
	}
	fmt.Println("Fetched VPP mac Application:\n", string(profileXML))
}
