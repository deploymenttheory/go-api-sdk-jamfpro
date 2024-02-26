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
	// Fetching mobile device groups
	groups, err := client.GetMobileDeviceGroups()
	if err != nil {
		fmt.Printf("Error fetching mobile device groups: %s\n", err)
		return
	}

	// Pretty print the mobile device groups in XML
	groupsXML, err := xml.MarshalIndent(groups, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling mobile device groups data: %v", err)
	}
	fmt.Println("Mobile Device Groups:\n", string(groupsXML))
}
