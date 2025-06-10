package main

import (
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

	// Define a new profile
	newProfile := jamfpro.ResourceMobileDeviceConfigurationProfile{
		General: jamfpro.MobileDeviceConfigurationProfileSubsetGeneral{
			Name: "WiFi",
			Site: &jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
			Category: &jamfpro.SharedResourceCategory{
				ID:   -1,
				Name: "No category assigned",
			},
			DeploymentMethod: "Install Automatically",
			Payloads:         "<plist version=\"1\"><dict>...</dict></plist>", // Replace with actual XML payload
		},
		Scope: jamfpro.MobileDeviceConfigurationProfileSubsetScope{
			AllMobileDevices: false,
			AllJSSUsers:      false,
		},
		SelfService: jamfpro.MobileDeviceConfigurationProfileSubsetSelfService{
			// Fill in self service details if needed
		},
	}

	// Update a profile by Name (assuming name is known)
	updatedProfileByName, err := client.UpdateMobileDeviceConfigurationProfileByName("WiFi", &newProfile)
	if err != nil {
		fmt.Println("Error updating profile by Name:", err)
	} else {
		fmt.Println("Updated Profile by Name:", updatedProfileByName)
	}
}
