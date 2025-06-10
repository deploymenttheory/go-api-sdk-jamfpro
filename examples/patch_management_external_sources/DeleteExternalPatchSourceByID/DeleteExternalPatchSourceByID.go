package main

import (
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

	// Specify the ID to delete
	sourceID := "3"

	// Delete the patch external source
	err = client.DeleteExternalPatchSourceByID(sourceID)
	if err != nil {
		log.Fatalf("Error deleting patch external source: %v", err)
	}

	fmt.Printf("Successfully deleted patch external source with ID: %s\n", sourceID)
}
