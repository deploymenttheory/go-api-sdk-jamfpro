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

	// Call GetComputerGroupByID function
	groupID := 195 // Placeholder ID, replace with a valid ID
	group, err := client.GetComputerGroupByID(groupID)
	if err != nil {
		log.Fatalf("Error fetching Computer Group by ID: %v", err)
	}

	// Pretty print the group in XML
	groupXML, err := xml.MarshalIndent(group, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Fetched Computer Group by ID:\n", string(groupXML))
}
