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

	// The ID of the advanced mobile device search you want to retrieve
	searchID := "94" // Replace with the actual ID you want to retrieve

	// Call the GetAdvancedMobileDeviceSearchByID function
	search, err := client.GetAdvancedMobileDeviceSearchByID(searchID)
	if err != nil {
		log.Fatalf("Error fetching advanced mobile device search by ID: %v", err)
	}

	// Convert the response into pretty XML for printing
	output, err := xml.MarshalIndent(search, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling search to XML: %v", err)
	}

	// Print the pretty XML
	fmt.Printf("Advanced Mobile Device Search (ID: %s):\n%s\n", searchID, string(output))
}
