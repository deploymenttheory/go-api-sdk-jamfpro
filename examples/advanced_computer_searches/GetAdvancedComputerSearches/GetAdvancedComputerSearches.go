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

	// Call GetAdvancedComputerSearches function
	advancedComputerSearches, err := client.GetAdvancedComputerSearches()
	if err != nil {
		log.Fatalf("Error fetching advanced computer searches: %v", err)
	}

	// Pretty print the advanced computer searches in XML
	advancedComputerSearchesXML, err := xml.MarshalIndent(advancedComputerSearches, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling advanced computer searches data: %v", err)
	}
	fmt.Println("Fetched Advanced Computer Searches:\n", string(advancedComputerSearchesXML))
}
