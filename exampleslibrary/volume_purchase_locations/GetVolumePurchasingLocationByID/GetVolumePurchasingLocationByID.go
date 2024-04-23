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

	// Example ID to fetch a specific volume purchasing location
	specificID := "1" // Replace with a valid ID

	// Example of calling GetVolumePurchasingLocationByID
	fmt.Printf("Fetching volume purchasing location with ID %s...\n", specificID)
	vpl, err := client.GetVolumePurchasingLocationByID(specificID)
	if err != nil {
		fmt.Printf("Error fetching volume purchasing location by ID: %v\n", err)
		return
	}

	// Pretty print the JSON response
	jsonData, err := json.MarshalIndent(vpl, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Printf("Volume Purchasing Location [%s]: %s\n", specificID, jsonData)
}
