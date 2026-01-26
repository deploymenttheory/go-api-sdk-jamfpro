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

	revocationEnabled := true
	outbound := true

	request := &jamfpro.ResourceAdcsSettingsV1{
		DisplayName:       "Example AD CS Settings",
		CAName:            "EXAMPLE-SUBCA02-CA",
		FQDN:              "example-subca02.example.com",
		RevocationEnabled: &revocationEnabled,
		APIClientID:       "c1bcec08-5f34-40fa-af52-9d3413ac916d",
		Outbound:          &outbound,
	}

	created, err := client.CreateAdcsSettingsV1(request)
	if err != nil {
		log.Fatalf("Error creating AD CS settings: %v", err)
	}

	output, err := json.MarshalIndent(created, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling creation response: %v", err)
	}

	fmt.Printf("Created AD CS settings: %s\n", string(output))
}
