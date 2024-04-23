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

	// Fetch directory bindings
	bindings, err := client.GetDirectoryBindings()
	if err != nil {
		fmt.Println("Error fetching directory bindings:", err)
		return
	}

	// Pretty print the directory bindings in xml
	configurationsXML, err := xml.MarshalIndent(bindings, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling configurations data: %v", err)
	}
	fmt.Println("Fetched Disk Encryption Configurations:\n", string(configurationsXML))
}
