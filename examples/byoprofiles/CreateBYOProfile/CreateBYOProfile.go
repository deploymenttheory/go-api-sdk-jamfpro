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
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"
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
	// Create a BYOProfile structure to send
	newProfile := jamfpro.ResourceBYOProfile{
		General: jamfpro.BYOProfileSubsetGeneral{
			Name:        "Personal Device Profile with jamf pro sdk",
			Site:        jamfpro.SharedResourceSite{ID: -1, Name: "None"},
			Enabled:     true,
			Description: "Used for Android or iOS BYO device enrollments",
		},
	}

	// Convert the profile to XML to see the output (optional, for debug purposes)
	xmlData, err := xml.MarshalIndent(newProfile, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling XML: %v", err)
	}
	fmt.Printf("XML Request: %s\n", xmlData)

	// Now call the create function
	createdProfile, err := client.CreateBYOProfile(&newProfile)
	if err != nil {
		log.Fatalf("Error creating BYO Profile: %v", err)
	}
	fmt.Printf("Created BYO Profile: %+v\n", createdProfile)
}
