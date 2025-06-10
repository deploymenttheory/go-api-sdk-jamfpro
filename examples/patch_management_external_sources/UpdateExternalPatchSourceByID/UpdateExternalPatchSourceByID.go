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

	// Updated source details
	updatedSource := &jamfpro.ResourcePatchExternalSource{
		Name:       "Updated External Source",
		HostName:   "patch.example.com",
		Port:       443,
		SSLEnabled: true,
	}

	// Specify the ID to update
	sourceID := "3"

	// Update the patch external source
	resultSource, err := client.UpdateExternalPatchSourceByID(sourceID, updatedSource)
	if err != nil {
		log.Fatalf("Error updating patch external source: %v", err)
	}

	// Pretty print the updated source
	resp, err := json.MarshalIndent(resultSource, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated source data: %v", err)
	}
	fmt.Printf("Updated Patch External Source:\n%s\n", string(resp))
}
