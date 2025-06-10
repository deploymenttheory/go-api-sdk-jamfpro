package main

import (
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

	// Create an instance of ResourceLocalAdminPasswordSettings with desired settings
	settingsToUpdate := &jamfpro.ResourceLocalAdminPasswordSettings{
		AutoDeployEnabled:        true,
		PasswordRotationTime:     3600,
		AutoRotateEnabled:        true,
		AutoRotateExpirationTime: 7776000,
	}

	// Call UpdateLocalAdminPasswordSettings function with the new settings
	err = client.UpdateLocalAdminPasswordSettings(settingsToUpdate)
	if err != nil {
		log.Fatalf("Error updating LAPS properties: %v", err)
	}

	fmt.Println("LAPS properties updated successfully.")
}
