package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
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

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// The ID of the static computer group you wish to update
	groupID := 49 // Replace with your actual group ID

	// Define the updated computers for the static group
	updatedComputers := []jamfpro.ComputerGroupComputerItem{
		{
			ID:            2,
			Name:          "MacBook Pro",
			MacAddress:    "",
			AltMacAddress: "",
			SerialNumber:  "D2FHXH22QB",
		},
		{
			ID:            6,
			Name:          "MacBook Pro",
			MacAddress:    "",
			AltMacAddress: "",
			SerialNumber:  "LT6M4DTF88",
		},
	}

	// Create the updated static computer group data
	updatedStaticGroup := &jamfpro.ResponseComputerGroup{
		Name:      "Updated Static Group Name",
		IsSmart:   false,
		Site:      jamfpro.ComputerGroupSite{ID: -1, Name: "None"},
		Computers: updatedComputers,
	}

	// Call UpdateComputerGroupByID function
	updatedGroup, err := client.UpdateComputerGroupByID(groupID, updatedStaticGroup)
	if err != nil {
		log.Fatalf("Error updating Computer Group by ID: %v", err)
	}

	// Pretty print the updated group in XML
	groupXML, err := xml.MarshalIndent(updatedGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Updated Computer Group:\n", string(groupXML))
}
