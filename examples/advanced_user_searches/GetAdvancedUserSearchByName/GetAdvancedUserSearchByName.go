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

	// Name of the advanced user search to retrieve
	advancedUserSearchName := "YourSearchName" // Replace with the actual name

	// Call GetAdvancedUserSearchByName function
	advancedUserSearch, err := client.GetAdvancedUserSearchByName(advancedUserSearchName)
	if err != nil {
		log.Fatalf("Error fetching advanced user search by name: %v", err)
	}

	// Pretty print the advanced user search in XML
	advancedUserSearchXML, err := xml.MarshalIndent(advancedUserSearch, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling advanced user search data: %v", err)
	}
	fmt.Println("Fetched Advanced User Search by Name:\n", string(advancedUserSearchXML))
}
