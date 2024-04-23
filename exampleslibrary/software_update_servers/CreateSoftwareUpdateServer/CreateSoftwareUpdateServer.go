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

	// Example payload for creating a new software update server
	newServer := &jamfpro.ResourceSoftwareUpdateServer{
		Name:          "New York SUS",
		IPAddress:     "10.10.51.248",
		Port:          8088,
		SetSystemWide: true,
	}

	// Call CreateSoftwareUpdateServer
	createdServer, err := client.CreateSoftwareUpdateServer(newServer)
	if err != nil {
		log.Fatalf("Error creating software update server: %v", err)
	}

	// Pretty print the created server details in XML
	createdServerXML, err := xml.MarshalIndent(createdServer, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created server data: %v", err)
	}
	fmt.Println("Created Software Update Server Details:\n", string(createdServerXML))
}
