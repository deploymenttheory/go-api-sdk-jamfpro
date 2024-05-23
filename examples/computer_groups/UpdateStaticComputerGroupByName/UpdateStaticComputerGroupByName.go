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

	// The name of the static computer group you wish to update
	groupName := "Updated Static Group Name" // Replace with your actual group name

	// Define the updated computers for the static group
	updatedComputers := []jamfpro.ComputerGroupSubsetComputer{
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
	updatedStaticGroup := &jamfpro.ResourceComputerGroup{
		Name:      "Static Group Name",
		IsSmart:   false,
		Site:      jamfpro.SharedResourceSite{ID: -1, Name: "None"},
		Computers: &updatedComputers,
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
