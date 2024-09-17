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
		Group: jamfpro.ResourcManagedSoftwareUpdatePlanObject{
			ObjectType: "COMPUTER_GROUP", // COMPUTER_GROUP / MOBILE_DEVICE_GROUP
			GroupId:    "55",
			// Do not set DeviceId for a group request
		},
		Config: jamfpro.ResourcManagedSoftwareUpdatePlanConfig{
			UpdateAction:    "DOWNLOAD_INSTALL_ALLOW_DEFERRAL", // DOWNLOAD_ONLY / DOWNLOAD_INSTALL / DOWNLOAD_INSTALL_ALLOW_DEFERRAL / DOWNLOAD_INSTALL_RESTART / DOWNLOAD_INSTALL_SCHEDULE / UNKNOWN
			VersionType:     "LATEST_MAJOR",                    // LATEST_MAJOR / LATEST_MINOR / LATEST_ANY / SPECIFIC_VERSION / UNKNOWN
			SpecificVersion: "NO_SPECIFIC_VERSION",             // NO_SPECIFIC_VERSION / 14.4.1 etc
			MaxDeferrals:    5,
			//ForceInstallLocalDateTime: "2023-12-25T21:09:31",
		},
	}

	// Call CreateManagedSoftwareUpdatePlanByGroupID function
	createdPlan, err := client.CreateManagedSoftwareUpdatePlanByGroupID(sampleUpdatePlan)
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
