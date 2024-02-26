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

	// Replace with the correct ID
	softwareUpdateServerByID, err := client.GetSoftwareUpdateServerByID(1) // Example ID
	if err != nil {
		log.Fatalf("Error fetching software update server by ID: %v", err)
	}

	// Pretty print the details in XML
	softwareUpdateServerByIDXML, err := xml.MarshalIndent(softwareUpdateServerByID, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling data: %v", err)
	}
	fmt.Println("Software Update Server by ID Details:\n", string(softwareUpdateServerByIDXML))
}
