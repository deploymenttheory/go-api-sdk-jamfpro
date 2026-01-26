package main

import (
	"encoding/json"
	"fmt"
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

	settings, err := client.GetAdcsSettingsByIDV1(adcsSettingsID)
	if err != nil {
		log.Fatalf("Error fetching AD CS settings %s: %v", adcsSettingsID, err)
	}

	output, err := json.MarshalIndent(settings, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling AD CS settings: %v", err)
	}

	fmt.Printf("AD CS settings %s: %s\n", adcsSettingsID, string(output))
}
