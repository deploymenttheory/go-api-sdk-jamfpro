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

	advancedComputerSearchID := 1 // Replace 1 with the actual advanced computer search ID

	// Call GetAdvancedComputerSearchByID function using the constant ID
	advancedComputerSearch, err := client.GetAdvancedComputerSearchByID(advancedComputerSearchID)
	if err != nil {
		log.Fatalf("Error fetching advanced computer search by ID: %v", err)
	}

	// Pretty print the advanced computer search in XML
	advancedComputerSearchXML, err := xml.MarshalIndent(advancedComputerSearch, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling advanced computer search data: %v", err)
	}
	fmt.Println("Fetched Advanced Computer Search by ID:\n", string(advancedComputerSearchXML))
}
