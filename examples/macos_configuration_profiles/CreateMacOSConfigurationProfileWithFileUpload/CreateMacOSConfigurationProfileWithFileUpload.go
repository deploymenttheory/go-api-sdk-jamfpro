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
	payloadFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/accessibility-chara-nosub-test.mobileconfig"
	//payloadFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/accessibility-chara-sub-test.mobileconfig"
	payload, err := readPayloadFromFile(payloadFilePath)
	if err != nil {
		log.Fatalf("Error reading payload from file: %v", err)
	}

	// Define a sample macOS Configuration Profile
	profile := &jamfpro.ResourceMacOSConfigurationProfile{
		General: jamfpro.MacOSConfigurationProfileSubsetGeneral{
			Name:               "accessibility-formatted-test",
			Site:               jamfpro.SharedResourceSite{Name: "None"},
			Category:           jamfpro.SharedResourceCategory{Name: "No category assigned"},
			DistributionMethod: "Install Automatically",
			Level:              "computer",
			RedeployOnUpdate:   "Newly Assigned",
			Payloads:           payload,
		},
		Scope: jamfpro.MacOSConfigurationProfileSubsetScope{},
		SelfService: jamfpro.MacOSConfigurationProfileSubsetSelfService{
			InstallButtonText: "Install",
		},
	}

	// Call CreateMacOSConfigurationProfile function
	createdProfile, err := client.CreateMacOSConfigurationProfile(profile)
	if err != nil {
		log.Fatalf("Error creating macOS Configuration Profile: %v", err)
	}

	// Print the ID of the created profile
	fmt.Printf("Successfully created macOS Configuration Profile with ID: %d\n", createdProfile)

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
