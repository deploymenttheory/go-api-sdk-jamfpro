package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
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

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Specify the updated details for the building
	buildingUpdate := &jamfpro.ResponseBuilding{
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
	updatedBuilding, err := client.UpdateBuildingByNameByID(buildingName, buildingUpdate)
	if err != nil {
		log.Fatalf("Error updating building: %v", err)
	}

	// Output the result
	fmt.Printf("Updated Building: %+v\n", updatedBuilding)
}
