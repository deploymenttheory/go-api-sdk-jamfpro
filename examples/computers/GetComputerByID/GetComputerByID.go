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

	// Define the computer ID you want to retrieve
	computerID := 21 // Replace with the actual ID of the computer

	// Call GetComputerByID function
	computer, err := client.GetComputerByID(computerID)
	if err != nil {
		log.Fatalf("Error retrieving computer by ID: %v", err)
	}

	// Pretty print the computer in XML
	computerXML, err := xml.MarshalIndent(computer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling computer data: %v", err)
	}
	fmt.Println("Retrieved Computer:\n", string(computerXML))
}
