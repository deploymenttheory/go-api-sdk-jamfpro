package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for Jamf Pro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new Jamf Pro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Call the GetAdvancedMobileDeviceSearches function
	searches, err := client.GetAdvancedMobileDeviceSearches()
	if err != nil {
		fmt.Println("Error fetching advanced mobile device searches:", err)
		return
	}

	// Pretty print the results as XML
	searchesXML, err := xml.MarshalIndent(searches, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling searches to XML: %v", err)
	}
	fmt.Println("Advanced Mobile Device Searches XML:")
	fmt.Println(string(searchesXML))

}
