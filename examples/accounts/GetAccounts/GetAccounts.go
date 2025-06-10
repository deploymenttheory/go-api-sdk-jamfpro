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

	// Call GetAccounts function
	accountsList, err := client.GetAccounts()
	if err != nil {
		log.Fatalf("Error fetching accounts: %v", err)
	}

	// Pretty print the accounts details
	accountsXML, err := xml.MarshalIndent(accountsList, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling accounts data: %v", err)
	}
	fmt.Println("Fetched Accounts List:", string(accountsXML))
}
