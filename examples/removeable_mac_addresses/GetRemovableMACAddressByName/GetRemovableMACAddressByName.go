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

	macAddressName := "E0:AC:CB:97:36:G4" // Replace with the actual name

	macAddress, err := client.GetRemovableMACAddressByName(macAddressName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	macAddressXML, err := xml.MarshalIndent(macAddress, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling MAC address data: %v", err)
	}
	fmt.Println("MAC Address Details:\n", string(macAddressXML))
}
