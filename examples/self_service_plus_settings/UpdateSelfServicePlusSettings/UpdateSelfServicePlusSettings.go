package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create settings to update
	settings := jamfpro.ResourceSelfServicePlusSettings{
		Enabled: true,
	}

	// Update the settings
	err = client.UpdateSelfServicePlusSettings(settings)
	if err != nil {
		log.Fatalf("Error updating settings: %v", err)
	}

	fmt.Println("Successfully updated Self Service Plus settings")
}
