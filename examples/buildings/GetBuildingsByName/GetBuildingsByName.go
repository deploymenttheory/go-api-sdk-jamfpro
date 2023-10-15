package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	concurrentRequests           = 10 // Number of simultaneous requests.
	maxConcurrentRequestsAllowed = 5  // Maximum allowed concurrent requests.
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:             authConfig.InstanceName,
		DebugMode:                true,
		Logger:                   jamfpro.NewDefaultLogger(),
		MaxConcurrentRequests:    maxConcurrentRequestsAllowed,
		TokenLifespan:            defaultTokenLifespan,
		TokenRefreshBufferPeriod: defaultBufferPeriod,
		ClientID:                 authConfig.ClientID,
		ClientSecret:             authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the variable for the building name
	buildingName := "bld-terraform_test-0.0.1-test" // Change this value as needed

	// Call GetBuildingByName function
	building, err := client.GetBuildingByName(buildingName)
	if err != nil {
		log.Fatalf("Error fetching building by Name: %v", err)
	}

	// Pretty print the building details using XML marshaling
	buildingXML, err := xml.MarshalIndent(building, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling building data: %v", err)
	}
	fmt.Println("Fetched Building Details:", string(buildingXML))
}
