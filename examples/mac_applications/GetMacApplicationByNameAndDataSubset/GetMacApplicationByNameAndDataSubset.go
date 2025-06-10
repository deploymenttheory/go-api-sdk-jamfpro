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

	// Define the application name and the subset you want to retrieve
	appName := "TextWrangler.app" // Replace with the actual application name
	subset := "General"           // Subset values can be General, Scope, SelfService, VPPCodes and VPP.

	// Call GetMacApplicationByNameAndDataSubset
	macApp, err := client.GetMacApplicationByNameAndDataSubset(appName, subset)
	if err != nil {
		log.Fatalf("Error fetching Mac Application by Name and Subset: %v", err)
	}

	// Pretty print the response in XML
	macAppXML, err := xml.MarshalIndent(macApp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Mac Application data: %v", err)
	}
	fmt.Println("Fetched Mac Application Data:\n", string(macAppXML))
}
