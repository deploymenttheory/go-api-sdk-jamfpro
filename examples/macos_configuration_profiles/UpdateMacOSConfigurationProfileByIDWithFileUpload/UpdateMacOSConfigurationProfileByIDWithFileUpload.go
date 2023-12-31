package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	maxConcurrentRequestsAllowed = 5
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
)

func readPayloadFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	authConfig, err := http_client.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Load payload from file
	payloads, err := readPayloadFromFile("/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/support_files/accessibility-chara-nosub-test.mobileconfig")
	if err != nil {
		log.Fatalf("Failed to read payload: %v", err)
	}

	// General profile data
	// Define the macOS Configuration Profile as per the given XML structure
	profile := jamfpro.ResourceMacOSConfigurationProfiles{
		General: jamfpro.MacOSConfigurationProfilesSubsetGeneral{
			Name:               "WiFi Test",
			Description:        "",
			Site:               jamfpro.SharedResourceSite{ID: -1, Name: "None"},                     // Optional, the Create fuction will set default values if no site is set
			Category:           jamfpro.SharedResourceCategory{ID: -1, Name: "No category assigned"}, // Optional, the Create fuction will set default values if no category is set
			DistributionMethod: "Install Automatically",
			UserRemovable:      false,
			Level:              "computer",
			RedeployOnUpdate:   "Newly Assigned",
			Payloads:           payloads,
		},
		Scope: jamfpro.MacOSConfigurationProfilesSubsetScope{
			AllComputers: false,
			AllJSSUsers:  false,
		},
		SelfService: jamfpro.MacOSConfigurationProfilesSubsetSelfService{
			InstallButtonText:           "Install",
			SelfServiceDescription:      "null",
			ForceUsersToViewDescription: false,
			// Add other fields as per the XML example
		},
	}

	// Set the config profile ID you want to update
	id := 153 // Replace with the actual ID of the profile you want to update

	// Call the UpdateMacOSConfigurationProfileByID function
	updatedProfileID, err := client.UpdateMacOSConfigurationProfileByID(id, &profile)
	if err != nil {
		log.Fatalf("Failed to update macOS Configuration Profile: %v", err)
	}

	fmt.Printf("Profile updated successfully. Updated Profile ID: %d\n", updatedProfileID)
}
