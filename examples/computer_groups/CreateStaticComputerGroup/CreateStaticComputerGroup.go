package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the computers for the static group
	computers := []jamfpro.ComputerGroupComputerItem{
		{
			ID:            1,
			Name:          "Device Name",
			MacAddress:    "",
			AltMacAddress: "9A:00:01:B7:A7:90",
			SerialNumber:  "C02Q7KHTGFWF",
		},
		{
			ID:            3,
			Name:          "Alishia's MacBook Air",
			MacAddress:    "00:16:3E:45:A7:90",
			AltMacAddress: "",
			SerialNumber:  "C02Q7KHTGFWF",
		},
	}

	// Create a new static computer group
	newStaticGroup := &jamfpro.ComputerGroupRequest{
		Name:      "Static Group",
		IsSmart:   false,
		Site:      jamfpro.Site{ID: -1, Name: "None"},
		Computers: computers,
	}

	// Call CreateComputerGroup function
	createdGroup, err := client.CreateComputerGroup(newStaticGroup)
	if err != nil {
		log.Fatalf("Error creating Computer Group: %v", err)
	}

	// Pretty print the created group in XML
	groupXML, err := xml.MarshalIndent(createdGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Created Computer Group:\n", string(groupXML))
}
