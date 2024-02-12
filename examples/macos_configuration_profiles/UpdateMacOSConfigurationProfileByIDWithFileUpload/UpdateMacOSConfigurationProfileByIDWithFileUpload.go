package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
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
	profile := jamfpro.ResourceMacOSConfigurationProfile{
		General: jamfpro.MacOSConfigurationProfileSubsetGeneral{
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
		Scope: jamfpro.MacOSConfigurationProfileSubsetScope{
			AllComputers: false,
			AllJSSUsers:  false,
		},
		SelfService: jamfpro.MacOSConfigurationProfileSubsetSelfService{
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

func readPayloadFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
