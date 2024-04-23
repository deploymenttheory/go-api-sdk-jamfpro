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

	// Retrieve a network segment by name
	name := "NY Office" // Replace with actual name
	segment, err := client.GetNetworkSegmentByName(name)
	if err != nil {
		log.Fatalf("Error fetching network segment by name: %v", err)
	}

	// Pretty print the network segments in XML
	segmentsXML, err := xml.MarshalIndent(segment, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling network segments data: %v", err)
	}
	fmt.Println("Network Segments:\n", string(segmentsXML))
}
