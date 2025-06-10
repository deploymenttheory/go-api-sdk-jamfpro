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

	// Call GetConditionalAccessDeviceComplianceFeatureEnablemen function
	groups, err := client.GetConditionalAccessDeviceComplianceFeatureEnablement()
	if err != nil {
		log.Fatalf("Error fetching ConditionalAccessDeviceComplianceFeatureEnablement Preferences: %v", err)
	}

	// Pretty print the groups in JSON
	featureSettingsJSON, err := json.MarshalIndent(groups, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling ConditionalAccessDeviceComplianceFeatureEnablement Preferences data: %v", err)
	}
	fmt.Println("Fetched ConditionalAccessDeviceComplianceFeatureEnablement Preferences:\n", string(featureSettingsJSON))
}
