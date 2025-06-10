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

	// Initialize the Jamf Pro client
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Get all patch external sources
	sources, err := client.GetPatchExternalSources()
	if err != nil {
		log.Fatalf("Error fetching patch external sources: %v", err)
	}

	// Pretty print the sources
	resp, err := json.MarshalIndent(sources, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling sources data: %v", err)
	}
	fmt.Printf("Patch External Sources:\n%s\n", string(resp))
}
