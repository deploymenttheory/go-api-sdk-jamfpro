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
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Example usage of CreateWebhook
	newWebhook := &jamfpro.ResourceWebhook{
		Name:               "Sample",
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

	// Example usage of UpdateWebhookByName
	updatedWebhookByName, err := client.UpdateWebhookByName("Sample", newWebhook) // Replace with the desired webhook name
	if err != nil {
		fmt.Printf("Error updating webhook by Name: %v\n", err)
		return
	}
	fmt.Printf("Updated Webhook by Name: %+v\n", updatedWebhookByName)
}
