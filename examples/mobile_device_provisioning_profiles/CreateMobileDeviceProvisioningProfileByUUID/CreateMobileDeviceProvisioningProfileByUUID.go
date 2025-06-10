package main

import (
	"encoding/xml"
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

	profileUUID := "116AF1E6-7EB5-4335-B598-276CDE5E015B" // Replace with the actual UUID
	newProfile := &jamfpro.ResourceMobileDeviceProvisioningProfile{
		General: jamfpro.MobileDeviceProvisioningProfileSubsetGeneral{
			Name:        "in-house app profile",
			DisplayName: "in-house app profile",
			UUID:        "116AF1E6-7EB5-4335-B598-276CDE5E015B",
		},
	}

	createdProfile, err := client.CreateMobileDeviceProvisioningProfileByUUID(profileUUID, newProfile)
	if err != nil {
		log.Fatalf("Error creating mobile device provisioning profile: %s\n", err)
	}

	createdProfileXML, err := xml.MarshalIndent(createdProfile, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created profile data: %v", err)
	}
	fmt.Println("Created Mobile Device Provisioning Profile:\n", string(createdProfileXML))
}
