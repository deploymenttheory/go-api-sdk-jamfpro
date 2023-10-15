package main

import (
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

	// Define the allowed file extension you want to create
	newExtension := &jamfpro.AllowedFileExtensionDetail{
		Extension: "qwerty3", // Replace with the desired extension name
	}

	response, err := client.CreateAllowedFileExtension(newExtension)
	if err != nil {
		log.Fatalf("Error creating allowed file extension: %v", err)
	}

	fmt.Printf("Successfully created allowed file extension: %s with ID: %d\n", response.Extension, response.ID)

}
