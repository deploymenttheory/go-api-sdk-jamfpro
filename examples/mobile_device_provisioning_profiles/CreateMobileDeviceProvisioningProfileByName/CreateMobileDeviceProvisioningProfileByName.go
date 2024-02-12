package main

import (
	"encoding/xml"
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

	// Instantiate the default logger and set the desired log level
	//logLevel := logger.LogLevelInfo // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

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

	profileName := "in-house app profile creation by name" // Replace with the actual profile name
	newProfile := &jamfpro.ResourceMobileDeviceProvisioningProfile{
		General: jamfpro.MobileDeviceProvisioningProfileSubsetGeneral{
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
