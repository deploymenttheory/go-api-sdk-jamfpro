package main

import (
	"fmt"
	"log"

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

	newProfile := jamfpro.ResourceMobileDeviceConfigurationProfile{
		General: jamfpro.MobileDeviceConfigurationProfileSubsetGeneral{
			Name: "WiFi",
			Site: jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
			Category: jamfpro.SharedResourceCategory{
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

	createdProfile, err := client.CreateMobileDeviceConfigurationProfile(&newProfile)
	if err != nil {
		log.Fatalf("Error creating mobile device configuration profile: %v", err)
	}

	fmt.Printf("Created Profile: %+v\n", createdProfile)
}
