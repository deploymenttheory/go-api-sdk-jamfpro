package main

import (
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	adcsSettingsID := "4"
	revocationEnabled := false

	updatePayload := &jamfpro.ResourceAdcsSettingsV1{
		DisplayName:       "Example AD CS Settings (Updated)",
		RevocationEnabled: &revocationEnabled,
	}

	if err := client.UpdateAdcsSettingsByIDV1(adcsSettingsID, updatePayload); err != nil {
		log.Fatalf("Error updating AD CS settings %s: %v", adcsSettingsID, err)
	}

	log.Printf("Successfully updated AD CS settings %s", adcsSettingsID)
}
