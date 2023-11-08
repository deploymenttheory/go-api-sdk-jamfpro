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
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Replace '123' with the actual ID of the subscription you want to retrieve.
	subscriptionID := "1"

	// Call the GetVolumePurchasingSubscriptionByID function with the subscription ID.
	subscription, err := client.GetVolumePurchasingSubscriptionByID(subscriptionID)
	if err != nil {
		fmt.Printf("Error fetching volume purchasing subscription with ID %s: %s\n", subscriptionID, err)
		return
	}

	// If no error, print the details of the retrieved subscription.
	fmt.Printf("Subscription ID: %s\n", subscription.Id)
	fmt.Printf("Name: %s\n", subscription.Name)
	fmt.Printf("Enabled: %t\n", subscription.Enabled)
	fmt.Printf("Triggers: %v\n", subscription.Triggers)
	fmt.Printf("Location IDs: %v\n", subscription.LocationIds)
	fmt.Printf("Internal Recipients: %+v\n", subscription.InternalRecipients)
	fmt.Printf("External Recipients: %+v\n", subscription.ExternalRecipients)
	fmt.Printf("Site ID: %s\n", subscription.SiteId)
}
