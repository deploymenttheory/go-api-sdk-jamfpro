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

	// Call GetComputerInventoryCollectionInformation
	checkinSettings, err := client.GetComputerInventoryCollectionInformation()
	if err != nil {
		fmt.Printf("Error fetching computer check-in settings: %s\n", err)
		return
	}

	// Pretty print the created attribute in XML
	inventorySettingsXML, err := xml.MarshalIndent(checkinSettings, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling checkin settings data: %v", err)
	}
	fmt.Println("computer check-in settings:\n", string(inventorySettingsXML))
}
