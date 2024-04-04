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

	// Perform the GetAccounts operation 100 times
	for i := 0; i < 100; i++ {
		accountsList, err := client.GetAccounts()
		if err != nil {
			log.Printf("Error fetching accounts at iteration %d: %v", i, err)
			continue // Skip this iteration if there's an error
		}

		// Pretty print the accounts details
		accountsXML, err := xml.MarshalIndent(accountsList, "", "    ") // Indent with 4 spaces
		if err != nil {
			log.Printf("Error marshaling accounts data at iteration %d: %v", i, err)
			continue // Skip this iteration if there's an error
		}
		fmt.Printf("Fetched Accounts List at iteration %d:\n%s\n", i, string(accountsXML))
	}
}
