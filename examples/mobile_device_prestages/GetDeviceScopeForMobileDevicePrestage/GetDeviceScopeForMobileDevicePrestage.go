package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/neilmartin/GitHub/go-api-sdk-jamfpro/client_auth.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// The ID of the mobile device prestage to retrieve the device scope for
	prestageID := "4" // Replace with the actual ID

	// Fetch the device scope for the specified mobile device prestage
	deviceScope, err := client.GetDeviceScopeForMobileDevicePrestageByID(prestageID)
	if err != nil {
		log.Fatalf("Error fetching device scope for mobile device prestage: %v", err)
	}

	// Pretty print the computer prestage in JSON
	prestageJSON, err := json.MarshalIndent(deviceScope, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling mobile device prestage data: %v", err)
	}
	fmt.Println("Fetched mobile device prestage:\n", string(prestageJSON))
}
