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

	// The name of the static computer group you wish to update
	groupName := "Updated Static Group Name" // Replace with your actual group name

	// Define the updated computers for the static group
	updatedComputers := []jamfpro.ComputerGroupComputerItem{
		{
			ID:            2,
			Name:          "MacBook Pro",
			MacAddress:    "",
			AltMacAddress: "",
			SerialNumber:  "D2FCXH22FH",
		},
		// ... add more updated computers if needed
	}

	// Create the updated static computer group data
	updatedStaticGroup := &jamfpro.ComputerGroupRequest{
		Name:      "Static Group Name",
		IsSmart:   false,
		Site:      jamfpro.Site{ID: -1, Name: "None"},
		Computers: updatedComputers,
	}

	// Call UpdateComputerGroupByName function
	updatedGroup, err := client.UpdateComputerGroupByName(groupName, updatedStaticGroup)
	if err != nil {
		log.Fatalf("Error updating Computer Group by Name: %v", err)
	}

	// Pretty print the updated group in XML
	groupXML, err := xml.MarshalIndent(updatedGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Updated Computer Group:\n", string(groupXML))
}
