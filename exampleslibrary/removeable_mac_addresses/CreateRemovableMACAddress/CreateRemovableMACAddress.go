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

	// Create a new Removable MAC Address
	newMACAddress := &jamfpro.ResourceRemovableMacAddress{
		Name: "E0:AC:CB:97:36:G4", // Replace with the actual MAC address name
		// ID: [set the ID if necessary]
	}

	createdMACAddress, err := client.CreateRemovableMACAddress(newMACAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Pretty print the created MAC address details in XML
	createdMACAddressXML, err := xml.MarshalIndent(createdMACAddress, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created MAC address data: %v", err)
	}
	fmt.Println("Created MAC Address Details:\n", string(createdMACAddressXML))
}
