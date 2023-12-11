package main

import (
	"fmt"
	"log"
	"time"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

const (
	concurrentRequests           = 10 // Number of simultaneous requests.
	maxConcurrentRequestsAllowed = 5  // Maximum allowed concurrent requests.
	defaultTokenLifespan         = 30 * time.Minute
	defaultBufferPeriod          = 5 * time.Minute
	apiIntegrationID             = 82
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/joseph/github/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelInfo // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Create an instance of the ApiIntegration with updated values
	integrationUpdate := &jamfpro.ResourceApiIntegration{
		AuthorizationScopes:        []string{"sdktest"}, // your api role names
		DisplayName:                "My API Integration - new name",
		Enabled:                    true,
		AccessTokenLifetimeSeconds: 1,
	}

	// Update the API Integration using its ID
	updatedIntegration, err := client.UpdateApiIntegrationByID(apiIntegrationID, integrationUpdate)
	if err != nil {
		fmt.Println("Error updating API Integration:", err)
		return
	}

	// Print the response
	fmt.Printf("Updated API Integration with ID: %d and Display Name: %s\n", apiIntegrationID, updatedIntegration.DisplayName)
}
