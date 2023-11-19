package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelInfo // Adjust log level as needed

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	profileName := "in-house app profile creation by name" // Replace with the actual profile name
	newProfile := &jamfpro.ResponseMobileDeviceProvisioningProfile{
		General: jamfpro.MobileDeviceProvisioningProfileGeneral{
			Name:        "in-house app profile creation by name",
			DisplayName: "in-house app profile",
			UUID:        "116AF1E6-7EB5-4335-B598-276CDE5E015B",
		},
	}

	createdProfile, err := client.CreateMobileDeviceProvisioningProfileByName(profileName, newProfile)
	if err != nil {
		log.Fatalf("Error creating mobile device provisioning profile: %s\n", err)
	}

	createdProfileXML, err := xml.MarshalIndent(createdProfile, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created profile data: %v", err)
	}
	fmt.Println("Created Mobile Device Provisioning Profile:\n", string(createdProfileXML))
}
