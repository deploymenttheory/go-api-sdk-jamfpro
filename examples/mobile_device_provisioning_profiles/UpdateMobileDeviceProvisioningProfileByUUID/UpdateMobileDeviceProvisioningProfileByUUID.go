package main

import (
	"encoding/xml"
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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	profileUDID := "116AF1E6-7EB5-4335-B598-276CDE5EXXXX" // Replace with the actual UDID

	updatedProfile := &jamfpro.ResourceMobileDeviceProvisioningProfile{
		General: jamfpro.MobileDeviceProvisioningProfileSubsetGeneral{
			Name:        "in-house app profile",
			DisplayName: "in-house app profile",
			UUID:        "116AF1E6-7EB5-4335-B598-276CDE5E015B",
		},
	}

	profile, err := client.UpdateMobileDeviceProvisioningProfileByUUID(profileUDID, updatedProfile)
	if err != nil {
		log.Fatalf("Error updating mobile device provisioning profile by UDID: %v", err)
	}

	profileXML, err := xml.MarshalIndent(profile, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated profile data: %v", err)
	}
	fmt.Println("Updated Mobile Device Provisioning Profile (ByUDID):\n", string(profileXML))
}
