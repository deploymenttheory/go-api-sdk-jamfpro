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

	// Call GetManagedSoftwareUpdates function
	updatePlans, err := client.GetManagedSoftwareUpdates()
	if err != nil {
		log.Fatalf("Error fetching managed software updates: %v", err)
	}

	// Pretty print the managed software updates in json
	updateJSON, err := json.MarshalIndent(updatePlans, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling managed software updates data: %v", err)
	}
	fmt.Println("Fetched managed software updates:\n", string(updateJSON))
}
