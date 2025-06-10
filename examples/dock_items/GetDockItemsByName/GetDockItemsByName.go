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

	// Define the name of the dock item to retrieve
	dockItemName := "Safari" // Replace with the actual name of the dock item

	// Call the GetDockItemsByName function
	dockItem, err := client.GetDockItemByName(dockItemName)
	if err != nil {
		log.Fatalf("Error fetching dock item by name: %v", err)
	}

	// Pretty print the updated group in XML
	groupXML, err := xml.MarshalIndent(dockItem, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Updated Computer Group:\n", string(groupXML))
}
