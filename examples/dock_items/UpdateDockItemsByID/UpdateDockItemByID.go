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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Define the ID and details of the dock item to update
	dockItemID := 1 // Replace with the actual ID
	updatedDockItem := &jamfpro.ResourceDockItem{
		// Set the fields you want to update
		Name:     "Updated Safari",
		Type:     "App",
		Path:     "file://localhost/Applications/Safari.app/",
		Contents: "Updated Contents",
	}

	// Call the UpdateDockItemByID function
	dockItem, err := client.UpdateDockItemByID(dockItemID, updatedDockItem)
	if err != nil {
		log.Fatalf("Error updating dock item by ID: %v", err)
	}

	// Pretty print the updated group in XML
	groupXML, err := xml.MarshalIndent(dockItem, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Updated Computer Group:\n", string(groupXML))
}
