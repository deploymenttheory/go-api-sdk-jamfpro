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

	profileName := "WiFi Test" // Replace with the actual profile name you want to fetch

	// Call the function to get the profile by name
	profile, err := client.GetMacOSConfigurationProfileByNameByID(profileName)
	if err != nil {
		fmt.Printf("Error fetching profile: %v\n", err)
		return
	}

	// Pretty print the details in XML
	configurationProfileXML, err := xml.MarshalIndent(profile, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling profile data: %v", err)
	}
	fmt.Println("Fetched Profile Details:\n", string(configurationProfileXML))
}
