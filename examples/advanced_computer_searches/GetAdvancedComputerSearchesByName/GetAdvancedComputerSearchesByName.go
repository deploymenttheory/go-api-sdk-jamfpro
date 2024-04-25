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

	advancedComputerSearchName := "Advanced Computer Search Name" // Replace with the actual advanced computer search name

	// Call GetAdvancedComputerSearchesByName function using the constant name
	advancedComputerSearchByName, err := client.GetAdvancedComputerSearchByName(advancedComputerSearchName)
	if err != nil {
		log.Fatalf("Error fetching advanced computer search by name: %v", err)
	}

	// Pretty print the advanced computer search by name in XML
	advancedComputerSearchByNameXML, err := xml.MarshalIndent(advancedComputerSearchByName, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling advanced computer search by name data: %v", err)
	}
	fmt.Println("Fetched Advanced Computer Search by Name:\n", string(advancedComputerSearchByNameXML))
}
