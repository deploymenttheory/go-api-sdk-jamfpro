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

	// The name of the advanced mobile device search you want to retrieve
	searchName := "Advanced Search Name" // Replace with the actual name you want to retrieve

	// Call the GetAdvancedMobileDeviceSearchByName function
	search, err := client.GetAdvancedMobileDeviceSearchByName(searchName)
	if err != nil {
		log.Fatalf("Error fetching advanced mobile device search by name: %v", err)
	}

	// Convert the response into pretty XML for printing
	output, err := xml.MarshalIndent(search, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling search to XML: %v", err)
	}

	// Print the pretty XML
	fmt.Printf("Advanced Mobile Device Search (Name: %s):\n%s\n", searchName, string(output))
}
