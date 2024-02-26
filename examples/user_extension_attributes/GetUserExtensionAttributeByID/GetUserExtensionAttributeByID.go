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

	// Example ID of the user extension attribute to retrieve
	attributeID := 1

	// Fetch user extension attribute by ID
	attribute, err := client.GetUserExtensionAttributeByID(attributeID)
	if err != nil {
		log.Fatalf("Error fetching user extension attribute by ID: %v", err)
	}

	// Pretty print the attribute details in XML
	attributeXML, err := xml.MarshalIndent(attribute, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling attribute data: %v", err)
	}
	fmt.Println("User Extension Attribute Details:\n", string(attributeXML))
}
