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

	// Replace with the correct name
	softwareUpdateServerByName, err := client.GetSoftwareUpdateServerByName("New York SUS") // Example name
	if err != nil {
		log.Fatalf("Error fetching software update server by name: %v", err)
	}

	// Pretty print the details in XML
	softwareUpdateServerByNameXML, err := xml.MarshalIndent(softwareUpdateServerByName, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)
	}
	fmt.Println("Software Update Server by Name Details:\n", string(softwareUpdateServerByNameXML))
}
