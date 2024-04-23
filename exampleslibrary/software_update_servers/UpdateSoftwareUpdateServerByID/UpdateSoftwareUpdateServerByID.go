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

	// Construct the update data
	serverToUpdate := &jamfpro.ResourceSoftwareUpdateServer{
		ID:            1, // Set the ID of the server to update
		Name:          "Updated New York SUS",
		IPAddress:     "10.10.51.249",
		Port:          8088,
		SetSystemWide: true,
	}

	// Call UpdateSoftwareUpdateServerByID
	updatedServer, err := client.UpdateSoftwareUpdateServerByID(serverToUpdate.ID, serverToUpdate)
	if err != nil {
		log.Fatalf("Error updating software update server by ID: %v", err)
	}

	// Pretty print the details in XML
	updatedServerXML, err := xml.MarshalIndent(updatedServer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated server data: %v", err)
	}
	fmt.Println("Updated Software Update Server Details:\n", string(updatedServerXML))
}
