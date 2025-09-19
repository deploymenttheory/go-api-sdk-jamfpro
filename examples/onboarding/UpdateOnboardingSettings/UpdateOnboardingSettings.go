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

	// Define the onboarding configuration you want to update
	onboardingRequest := jamfpro.ResourceUpdateOnboardingSettings{
		Enabled: true,
		OnboardingItems: []jamfpro.SubsetOnboardingItemRequest{
			{
				EntityID:              "4",
				SelfServiceEntityType: "OS_X_POLICY",
				Priority:              1,
			},
			{
				EntityID:              "1",
				SelfServiceEntityType: "OS_X_MAC_APP",
				Priority:              2,
			},
			{
				EntityID:              "7",
				SelfServiceEntityType: "OS_X_CONFIG_PROFILE",
				Priority:              3,
			},
		},
	}

	// Call UpdateOnboardingSettings function
	updatedSettings, err := client.UpdateOnboardingSettings(onboardingRequest)
	if err != nil {
		log.Fatalf("Error updating onboarding settings: %v", err)
	}

	// Pretty print the updated onboarding settings in JSON
	settingsJSON, err := json.MarshalIndent(updatedSettings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling onboarding settings data: %v", err)
	}
	fmt.Println("Updated Onboarding Settings:\n", string(settingsJSON))
}
