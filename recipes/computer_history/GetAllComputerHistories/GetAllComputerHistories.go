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

	// Call the GetComputers method
	computers, err := client.GetComputers()
	if err != nil {
		log.Fatalf("Error fetching computers: %v", err)
	}

	// Iterate through each computer to fetch its history
	for _, computer := range computers.Results {
		// Fetch computer history by ID
		computerHistory, err := client.GetComputerHistoryByComputerID(computer.ID)
		if err != nil {
			log.Printf("Error fetching computer history for ID %d: %v", computer.ID, err)
			continue
		}

		// Pretty print the response
		prettyXML, err := xml.MarshalIndent(computerHistory, "", "    ")
		if err != nil {
			log.Printf("Failed to generate pretty XML for computer ID %d: %v", computer.ID, err)
			continue
		}
		fmt.Printf("Computer ID: %d\n", computer.ID)
		fmt.Printf("%s\n", prettyXML)
	}
}
