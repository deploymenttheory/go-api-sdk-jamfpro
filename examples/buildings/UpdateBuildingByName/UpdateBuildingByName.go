package main

import (
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

	// Specify the updated details for the building
	buildingUpdate := &jamfpro.ResourceBuilding{
		Name:           "Updated Building Name",
		StreetAddress1: "Updated Address 1",
		StreetAddress2: "Updated Address 2",
		City:           "Updated City",
		StateProvince:  "Updated State",
		ZipPostalCode:  "Updated Zip Code",
		Country:        "Updated Country",
	}

	// Update building by name
	buildingName := "Updated Building Name" // Replace with the actual name
	updatedBuilding, err := client.UpdateBuildingByName(buildingName, buildingUpdate)
	if err != nil {
		log.Fatalf("Error updating building: %v", err)
	}

	// Output the result
	fmt.Printf("Updated Building: %+v\n", updatedBuilding)
}