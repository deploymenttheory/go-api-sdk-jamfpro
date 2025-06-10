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

	// Example ID of the directory binding to fetch
	bindingID := "1"

	// Fetch directory binding by ID
	binding, err := client.GetDirectoryBindingByID(bindingID)
	if err != nil {
		fmt.Println("Error fetching directory binding:", err)
		return
	}

	// Pretty print the directory binding in xml
	bindingXML, err := xml.MarshalIndent(binding, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling binding data: %v", err)
	}
	fmt.Printf("Fetched Directory Binding:\n%s\n", string(bindingXML))
}
