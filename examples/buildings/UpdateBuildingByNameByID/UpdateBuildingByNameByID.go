package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
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
	buildingName := "Apple Park" // Replace with the actual name
	updatedBuilding, err := client.UpdateBuildingByName(buildingName, buildingUpdate)
	if err != nil {
		log.Fatalf("Error updating building: %v", err)
	}

	// Output the result
	fmt.Printf("Updated Building: %+v\n", updatedBuilding)
}
