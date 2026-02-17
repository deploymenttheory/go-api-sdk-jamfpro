package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Path to the JSON configuration file for the SDK client
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	wellKnownSettings, err := client.GetServiceDiscoveryEnrollmentWellKnownSettingsV1()
	if err != nil {
		log.Fatalf("Error fetching service discovery enrollment well-known settings: %v", err)
	}

	responseJSON, err := json.MarshalIndent(wellKnownSettings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling service discovery enrollment well-known settings: %v", err)
	}

	fmt.Println("Service discovery enrollment well-known settings:", string(responseJSON))
}
