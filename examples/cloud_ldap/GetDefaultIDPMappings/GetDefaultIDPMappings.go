package main

import (
	"encoding/json"
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
	resp, err := client.GetDefaultCloudIdentityProviderDefaultMappings("")
	if err != nil {
		log.Fatalf("Error fetching default cloud identity provider mappings: %v", err)
	}

	// Pretty print the attributes in XML
	attributesXML, err := json.MarshalIndent(resp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling default cloud identity provider mappings data: %v", err)
	}
	fmt.Println("Fetched Computer default cloud identity provider mappings:\n", string(attributesXML))
}
