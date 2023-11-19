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

	// Name of the subscription to update
	subscriptionName := "Example Volume Purchasing Subscription"

	// Example update data
	updateData := jamfpro.VolumePurchasingSubscription{
		Enabled:  true,
		SiteId:   "-1",
		Name:     "Updated Volume Purchasing Subscription",
		Triggers: []string{"NO_MORE_LICENSES", "REMOVED_FROM_APP_STORE"},
		InternalRecipients: []struct {
			AccountId string `json:"accountId,omitempty"`
			Frequency string `json:"frequency,omitempty"`
		}{
			{Frequency: "DAILY", AccountId: "1"},
		},
		ExternalRecipients: []struct {
			Name  string `json:"name,omitempty"`
			Email string `json:"email,omitempty"`
		}{
			{Name: "Alice Examplewoman", Email: "alice@example.com"}, // New external recipient
		},
	}

	// Call the update function with the subscription name and update data
	updatedSubscription, err := client.UpdateVolumePurchasingSubscriptionByNameByID(subscriptionName, &updateData)
	if err != nil {
		fmt.Printf("Error updating volume purchasing subscription with name '%s': %s\n", subscriptionName, err)
		return
	}

	// If no error, print the details of the updated subscription
	fmt.Printf("Updated Subscription ID: %s\n", updatedSubscription.Id)
	fmt.Printf("Name: %s\n", updatedSubscription.Name)
	// ... print other fields as needed
}
