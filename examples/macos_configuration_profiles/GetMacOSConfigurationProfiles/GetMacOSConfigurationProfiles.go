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
