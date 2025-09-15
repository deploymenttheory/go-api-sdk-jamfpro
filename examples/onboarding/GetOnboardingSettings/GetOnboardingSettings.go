package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Call GetOnboardingSettings function
	onboardingSettings, err := client.GetOnboardingSettings()
	if err != nil {
		log.Fatalf("Error fetching onboarding settings: %v", err)
	}

	// Pretty print the onboarding settings in JSON
	settingsJSON, err := json.MarshalIndent(onboardingSettings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling onboarding settings data: %v", err)
	}
	fmt.Println("Current Onboarding Settings:\n", string(settingsJSON))
}
