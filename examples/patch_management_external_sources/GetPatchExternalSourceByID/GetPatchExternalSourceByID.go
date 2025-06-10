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

	// Specify the ID to fetch
	sourceID := "3"

	// Get the patch external source by ID
	source, err := client.GetPatchExternalSourceByID(sourceID)
	if err != nil {
		log.Fatalf("Error fetching patch external source: %v", err)
	}

	// Pretty print the source
	resp, err := json.MarshalIndent(source, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling source data: %v", err)
	}
	fmt.Printf("Patch External Source (ID: %s):\n%s\n", sourceID, string(resp))
}
