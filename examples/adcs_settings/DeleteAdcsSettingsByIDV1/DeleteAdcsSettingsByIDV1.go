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

	if err := client.DeleteAdcsSettingsByIDV1(adcsSettingsID); err != nil {
		log.Fatalf("Error deleting AD CS settings %s: %v", adcsSettingsID, err)
	}

	log.Printf("Successfully deleted AD CS settings %s", adcsSettingsID)
}
