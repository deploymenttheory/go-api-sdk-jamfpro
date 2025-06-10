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

	// Initialize the Jamf Pro client
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Specify the name to fetch
	sourceName := "Test External Source"

	// Get the patch external source by name
	source, err := client.GetPatchExternalSourceByName(sourceName)
	if err != nil {
		log.Fatalf("Error fetching patch external source: %v", err)
	}

	// Pretty print the source
	resp, err := json.MarshalIndent(source, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling source data: %v", err)
	}
	fmt.Printf("Patch External Source (Name: %s):\n%s\n", sourceName, string(resp))
}
