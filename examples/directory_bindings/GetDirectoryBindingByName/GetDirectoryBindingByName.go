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
	// Example name of the directory binding to fetch
	bindingName := "New Binding"

	// Fetch directory binding by Name
	binding, err := client.GetDirectoryBindingByName(bindingName)
	if err != nil {
		fmt.Println("Error fetching directory binding by name:", err)
		return
	}

	// Pretty print the directory binding in xml
	bindingXML, err := xml.MarshalIndent(binding, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling binding data: %v", err)
	}
	fmt.Printf("Fetched Directory Binding by Name:\n%s\n", string(bindingXML))
}
