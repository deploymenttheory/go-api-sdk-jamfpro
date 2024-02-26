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

	// Call the GetAdvancedMobileDeviceSearches function
	searches, err := client.GetAdvancedMobileDeviceSearches()
	if err != nil {
		fmt.Println("Error fetching advanced mobile device searches:", err)
		return
	}

	// Pretty print the results as XML
	searchesXML, err := xml.MarshalIndent(searches, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling searches to XML: %v", err)
	}
	fmt.Println("Advanced Mobile Device Searches XML:")
	fmt.Println(string(searchesXML))

}
