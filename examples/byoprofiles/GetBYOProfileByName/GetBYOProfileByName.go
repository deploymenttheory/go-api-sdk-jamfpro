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

	// Replace with the actual name you want to retrieve
	profileName := "Personal Device Profile"

	profile, err := client.GetBYOProfileByName(profileName)
	if err != nil {
		log.Fatalf("Error fetching BYO Profile by name: %v", err)
	}

	// Pretty print the account details
	byoprofileXML, err := xml.MarshalIndent(profile, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling account data: %v", err)
	}
	fmt.Println("Fetched BYO Profile Details:", string(byoprofileXML))
}
