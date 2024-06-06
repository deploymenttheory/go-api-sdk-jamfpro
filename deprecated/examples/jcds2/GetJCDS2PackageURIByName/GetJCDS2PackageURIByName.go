package main

import (
	"encoding/json"
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

	// Let's assume you want to get the uploaded package with name "Firefox 122.0.dmg"
	packageName := "powershell-7.4.1-osx-arm64.pkg"

	configuration, err := client.GetJCDS2PackageURIByName(packageName)
	if err != nil {
		log.Fatalf("Error fetching JCDS 2.0 package by name: %v", err)
	}

	// Print the configuration in a pretty JSON format
	configJSON, err := json.MarshalIndent(configuration, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling configuration data: %v", err)
	}
	fmt.Printf("Fetched JCDS 2.0 package by Name:\n%s\n", configJSON)
}
