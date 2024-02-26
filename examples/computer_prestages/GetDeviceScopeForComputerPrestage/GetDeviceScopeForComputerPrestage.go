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

	// The ID of the computer prestage to retrieve the device scope for
	prestageID := "123" // Replace with the actual ID

	// Fetch the device scope for the specified computer prestage
	deviceScope, err := client.GetDeviceScopeForComputerPrestageByID(prestageID)
	if err != nil {
		log.Fatalf("Error fetching device scope for computer prestage: %v", err)
	}

	// Pretty print the computer prestage in JSON
	prestageJSON, err := json.MarshalIndent(deviceScope, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling computer prestage data: %v", err)
	}
	fmt.Println("Fetched computer prestage:\n", string(prestageJSON))
}
