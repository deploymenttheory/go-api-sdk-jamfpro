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

	// Retrieve computer inventory collection settings
	settings, err := client.GetComputerInventoryCollectionSettings()
	if err != nil {
		log.Fatalf("Error fetching Computer Inventory Collection Settings: %s", err)
	}

	// Convert the settings to pretty-printed JSON
	settingsJSON, err := json.MarshalIndent(settings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling Computer Inventory Collection Settings to JSON: %s", err)
	}

	// Print the pretty-printed JSON
	fmt.Println("Computer Inventory Collection Settings:")
	fmt.Println(string(settingsJSON))
}
