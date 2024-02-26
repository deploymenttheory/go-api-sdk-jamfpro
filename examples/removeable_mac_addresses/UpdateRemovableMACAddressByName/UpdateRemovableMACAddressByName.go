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

	macAddressName := "Existing_Name" // Replace with the current name
	updatedMACAddress := &jamfpro.ResourceRemovableMacAddress{
		Name: "New_Name", // Replace with the new name
	}

	macAddress, err := client.UpdateRemovableMACAddressByName(macAddressName, updatedMACAddress)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	updatedMACAddressXML, err := xml.MarshalIndent(macAddress, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated MAC address data: %v", err)
	}
	fmt.Println("Updated MAC Address Details:\n", string(updatedMACAddressXML))
}
