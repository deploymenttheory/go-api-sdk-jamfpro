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

	// Call GetComputerExtensionAttributes function
	attributes, err := client.GetComputerExtensionAttributes()
	if err != nil {
		log.Fatalf("Error fetching Computer Extension Attributes: %v", err)
	}

	// Pretty print the attributes in XML
	attributesXML, err := xml.MarshalIndent(attributes, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Extension Attributes data: %v", err)
	}
	fmt.Println("Fetched Computer Extension Attributes:\n", string(attributesXML))
}
