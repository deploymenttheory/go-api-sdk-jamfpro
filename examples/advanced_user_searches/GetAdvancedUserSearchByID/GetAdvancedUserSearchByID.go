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

	// ID of the advanced user search to retrieve
	advancedUserSearchID := 29 // Replace with the actual ID

	// Call GetAdvancedUserSearchByID function
	advancedUserSearch, err := client.GetAdvancedUserSearchByID(advancedUserSearchID)
	if err != nil {
		log.Fatalf("Error fetching advanced user search by ID: %v", err)
	}

	// Pretty print the advanced user search in XML
	advancedUserSearchXML, err := xml.MarshalIndent(advancedUserSearch, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling advanced user search data: %v", err)
	}
	fmt.Println("Fetched Advanced User Search:\n", string(advancedUserSearchXML))
}
