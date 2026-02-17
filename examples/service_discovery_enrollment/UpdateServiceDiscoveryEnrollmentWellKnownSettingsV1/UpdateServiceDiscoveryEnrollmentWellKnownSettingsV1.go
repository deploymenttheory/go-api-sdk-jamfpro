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

	updatePayload := jamfpro.ResponseServiceDiscoveryEnrollmentWellKnownSettingsV1{
		WellKnownSettings: []jamfpro.ResourceServiceDiscoveryWellKnownSettingV1{
			{
				ServerUUID:     "9CB98B40026F4F0B8AAAB76636A7DD9D",
				EnrollmentType: "none",
			},
		},
	}

	if err := client.UpdateServiceDiscoveryEnrollmentWellKnownSettingsV1(updatePayload); err != nil {
		log.Fatalf("Error updating service discovery enrollment well-known settings: %v", err)
	}

	requestJSON, err := json.MarshalIndent(updatePayload, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated service discovery enrollment well-known settings: %v", err)
	}

	fmt.Println("Updated service discovery enrollment well-known settings:", string(requestJSON))
}
