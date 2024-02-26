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
	// Call GetSoftwareUpdateServers
	softwareUpdateServer, err := client.GetSoftwareUpdateServers()
	if err != nil {
		log.Fatalf("Error fetching software update servers: %v", err)
	}

	// Pretty print the details in XML
	softwareUpdateServerXML, err := xml.MarshalIndent(softwareUpdateServer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling server data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(softwareUpdateServerXML))
}
