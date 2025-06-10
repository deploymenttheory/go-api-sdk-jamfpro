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

	// Call GetAdvancedUserSearches function
	advancedUserSearches, err := client.GetAdvancedUserSearches()
	if err != nil {
		log.Fatalf("Error fetching advanced user searches: %v", err)
	}

	// Pretty print the advanced user searches in XML
	advancedUserSearchesXML, err := xml.MarshalIndent(advancedUserSearches, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling advanced user searches data: %v", err)
	}
	fmt.Println("Fetched Advanced User Searches:\n", string(advancedUserSearchesXML))
}
