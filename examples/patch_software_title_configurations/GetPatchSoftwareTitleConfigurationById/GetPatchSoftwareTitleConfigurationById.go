package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Example display name
	patchSoftwareTitleConfigurationId := "13" // Replace with an actual device name

	// Get patch software title configuration by name
	response, err := client.GetPatchSoftwareTitleConfigurationById(patchSoftwareTitleConfigurationId)
	if err != nil {
		log.Fatalf("Error fetching patch software title configuration by name: %v", err)
	}

	// Pretty print the network segments in JSON
	mobileDeviceJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling network segments data: %v", err)
	}
	fmt.Println("Network Segments:\n", string(mobileDeviceJSON))
}
