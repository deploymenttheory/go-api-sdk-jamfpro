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

	// Replace with the actual beacon name
	iBeaconName := "Room 1 Beacon"

	iBeacon, err := client.GetIBeaconByName(iBeaconName)
	if err != nil {
		log.Fatalf("Error retrieving iBeacon by ID: %v", err)
	}

	// Pretty print the iBeacon details
	XML, err := xml.MarshalIndent(iBeacon, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling accounts data: %v", err)
	}
	fmt.Println("Fetched Accounts List:", string(XML))
}
