package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define a sample plan for testing
	sampleUpdatePlan := &jamfpro.ResourceManagedSoftwareUpdatePlan{
		Devices: []jamfpro.ManagedSoftwareUpdatePlanDevice{{
			ObjectType: "COMPUTER",
			DeviceId:   "1",
		}},
		Config: jamfpro.ManagedSoftwareUpdatePlanConfig{
			UpdateAction:              "DOWNLOAD_INSTALL",
			VersionType:               "LATEST_MINOR",
			SpecificVersion:           "12.6.1",
			MaxDeferrals:              5,
			ForceInstallLocalDateTime: "2023-12-25T21:09:31",
		},
	}

	// Call CreateManagedSoftwareUpdatePlan function
	createdPlan, err := client.CreateManagedSoftwareUpdatePlan(sampleUpdatePlan)
	if err != nil {
		log.Fatalf("Error creating managed software update plan: %v", err)
	}

	// Pretty print the created managed software update plan details in JSON
	createdPlanJSON, err := json.MarshalIndent(createdPlan, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created managed software update plan data: %v", err)
	}
	fmt.Println("Created managed software update plan Details:\n", string(createdPlanJSON))
}
