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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Example of calling GetVolumePurchaseLocations
	fmt.Println("Fetching all volume purchasing locations...")
	vplList, err := client.GetVolumePurchaseLocations("") // Pass nil or empty for no sort/filter
	if err != nil {
		fmt.Printf("Error fetching volume purchasing locations: %v\n", err)
		return
	}

	// Pretty print the JSON response for all volume purchasing locations
	jsonData, err := json.MarshalIndent(vplList, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal JSON: %v", err)
	}
	fmt.Printf("Volume Purchasing Locations: %s\n", jsonData)
}
