package main

import (
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
	logLevel := http_client.LogLevelDebug

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

	// Example usage of CreateWebhook
	newWebhook := &jamfpro.ResponseWebhook{
		Name:               "Webhook",
		Enabled:            true,
		URL:                "https://server.com",
		ContentType:        "application/json",
		Event:              "SmartGroupComputerMembershipChange",
		ConnectionTimeout:  5,
		ReadTimeout:        2,
		AuthenticationType: "BASIC",
		Username:           "Sample User",
		Password:           "SamplePassword",
		SmartGroupID:       1,
	}

	createdWebhook, err := client.CreateWebhook(newWebhook)
	if err != nil {
		fmt.Printf("Error creating webhook: %v\n", err)
		return
	}
	fmt.Printf("Created Webhook: %+v\n", createdWebhook)
}
