package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Call GetMacOSConfigurationProfiles function
	profiles, err := client.GetMacOSConfigurationProfiles()
	if err != nil {
		log.Fatalf("Error fetching macOS Configuration Profiles: %v", err)
	}

	// Pretty print the profiles in XML
	profilesXML, err := xml.MarshalIndent(profiles, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling macOS Configuration Profiles data: %v", err)
	}
	fmt.Println("Fetched macOS Configuration Profiles:\n", string(profilesXML))
}
