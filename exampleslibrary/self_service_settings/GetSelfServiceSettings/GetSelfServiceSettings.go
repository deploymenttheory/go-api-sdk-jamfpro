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

	// Call GetSelfServiceSettings
	checkinSettings, err := client.GetSelfServiceSettings()
	if err != nil {
		fmt.Printf("Error fetching self service settings: %s\n", err)
		return
	}

	// Pretty print the created attribute in JSON
	selfServiceSettingsJSON, err := json.MarshalIndent(checkinSettings, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling self service settings data: %v", err)
	}
	fmt.Println("self service settings:\n", string(selfServiceSettingsJSON))
}
