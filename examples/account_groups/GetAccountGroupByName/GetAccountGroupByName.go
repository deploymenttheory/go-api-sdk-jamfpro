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

	// Define the variable for the group name
	groupName := "Test Group" // Change this value as needed

	// Call GetAccountGroupByName function
	group, err := client.GetAccountGroupByName(groupName)
	if err != nil {
		log.Fatalf("Error fetching group by name: %v", err)
	}

	// Pretty print the group details
	accountsXML, err := xml.MarshalIndent(group, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling group data: %v", err)
	}
	fmt.Println("Fetched Group Details:", string(accountsXML))
}
