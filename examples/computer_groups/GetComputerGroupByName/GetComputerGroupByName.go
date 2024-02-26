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
	// Call GetComputerGroupByName function
	groupName := "your-computer-group-name" // Placeholder name, replace with a valid name
	group, err := client.GetComputerGroupByName(groupName)
	if err != nil {
		log.Fatalf("Error fetching Computer Group by name: %v", err)
	}

	// Pretty print the group in XML
	groupXML, err := xml.MarshalIndent(group, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Fetched Computer Group by name:\n", string(groupXML))
}
