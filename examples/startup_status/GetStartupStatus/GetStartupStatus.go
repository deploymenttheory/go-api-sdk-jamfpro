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

	// Call GetStartupStatus function
	startupStatus, err := client.GetStartupStatus()
	if err != nil {
		log.Fatalf("Error fetching startup status: %v", err)
	}

	// Pretty print the startup status in JSON
	JSON, err := json.MarshalIndent(startupStatus, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling startup status data: %v", err)
	}
	fmt.Println("Jamf Pro Startup Status:\n", string(JSON))
}
