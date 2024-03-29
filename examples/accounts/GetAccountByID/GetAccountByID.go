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

	// Define the variable for the ID number
	accountID := 55 // Change this value as needed

	// Call GetAccountByID function
	account, err := client.GetAccountByID(accountID)
	if err != nil {
		log.Fatalf("Error fetching account by ID: %v", err)
	}

	// Pretty print the account details
	accountsXML, err := xml.MarshalIndent(account, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling account data: %v", err)
	}
	fmt.Println("Fetched Account Details:", string(accountsXML))
}
