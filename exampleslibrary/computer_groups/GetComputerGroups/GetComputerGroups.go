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

	// Call GetComputerGroups function
	groups, err := client.GetComputerGroups()
	if err != nil {
		log.Fatalf("Error fetching Computer Groups: %v", err)
	}

	// Pretty print the groups in XML
	groupsXML, err := xml.MarshalIndent(groups, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Groups data: %v", err)
	}
	fmt.Println("Fetched Computer Groups:\n", string(groupsXML))
}
