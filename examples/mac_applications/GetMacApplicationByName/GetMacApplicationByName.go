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

	name := "TextWrangler.app" // Replace with your Mac application name
	macApp, err := client.GetMacApplicationByName(name)
	if err != nil {
		log.Fatalf("Error fetching Mac Application by Name: %v", err)
	}

	// Pretty print the profile in XML
	profileXML, err := xml.MarshalIndent(macApp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling VPP mac Application data: %v", err)
	}
	fmt.Println("Fetched VPP mac Application:\n", string(profileXML))
}
