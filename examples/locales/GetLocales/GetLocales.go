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

	// Call function to get log flushing settings
	logFlushingSettings, err := client.GetLocales()
	if err != nil {
		log.Fatalf("Error fetching locales settings: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(logFlushingSettings, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling locales settings data: %v", err)
	}
	fmt.Println("Fetched locales settings:\n", string(response))
}
