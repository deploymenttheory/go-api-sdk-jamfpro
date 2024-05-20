package main

import (
	"fmt"
	"log"
	"os"

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

	// Read the payload from a file
	payloadFilePath := "/Users/dafyddwatkins/localtesting/terraform/support_files/mobileconfigurationprofiles/mobile-wifi.mobileconfig"

	payload, err := readPayloadFromFile(payloadFilePath)
	if err != nil {
		log.Fatalf("Error reading payload from file: %v", err)
	}

	// Define a sample Mobile Device Configuration Profile
	profile := &jamfpro.ResourceMobileDeviceConfigurationProfile{
		General: jamfpro.MobileDeviceConfigurationProfileSubsetGeneral{
			Name:             "jamfpro-sdk-localtest-mobiledeviceconfigprofile-iOS-WiFi",
			Description:      "An example mobile device configuration profile.",
			Site:             jamfpro.SharedResourceSite{Name: "None"},
			Category:         jamfpro.SharedResourceCategory{Name: "No category assigned"},
			DeploymentMethod: "Install Automatically",
			Level:            "Device Level",
			RedeployOnUpdate: "Newly Assigned",
			Payloads:         payload,
		},
		Scope: jamfpro.MobileDeviceConfigurationProfileSubsetScope{
			AllMobileDevices:   true,
			AllJSSUsers:        false,
			MobileDevices:      []jamfpro.MobileDeviceConfigurationProfileSubsetMobileDevice{},
			MobileDeviceGroups: []jamfpro.MobileDeviceConfigurationProfileSubsetScopeEntity{},
			Buildings:          []jamfpro.MobileDeviceConfigurationProfileSubsetScopeEntity{},
			Departments:        []jamfpro.MobileDeviceConfigurationProfileSubsetScopeEntity{},
			JSSUsers:           []jamfpro.MobileDeviceConfigurationProfileSubsetScopeEntity{},
			JSSUserGroups:      []jamfpro.MobileDeviceConfigurationProfileSubsetScopeEntity{},
			Limitations:        jamfpro.MobileDeviceConfigurationProfileSubsetLimitation{},
			Exclusions:         jamfpro.MobileDeviceConfigurationProfileSubsetExclusion{},
		},
	}

	// Call CreateMobileDeviceConfigurationProfile function
	createdProfile, err := client.CreateMobileDeviceConfigurationProfile(profile)
	if err != nil {
		log.Fatalf("Error creating mobile device Configuration Profile: %v", err)
	}

	// Print the ID of the created profile
	fmt.Printf("Successfully created mobile device Configuration Profile with ID: %d\n", createdProfile.ID)

}

// Helper function to get a pointer from a string
func stringPtr(s string) *string {
	return &s
}

// readPayloadFromFile loads config profile for upload
func readPayloadFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
