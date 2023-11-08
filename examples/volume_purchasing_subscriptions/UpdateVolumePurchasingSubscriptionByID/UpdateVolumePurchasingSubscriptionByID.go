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

	subscriptionID := "1" // Replace with the ID of the subscription you want to update

	// Example subscription with updated data
	updateSubscription := jamfpro.VolumePurchasingSubscription{
		Enabled:  true,
		SiteId:   "-1",
		Name:     "Example Volume Purchasing Subscription Updated",
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
			{Name: "Bob Exampleman Updated", Email: "bob.updated@example.com"},
		},
	}

	// Call the UpdateVolumePurchasingSubscriptionByID function with the subscription ID and updated data.
	updatedSubscription, err := client.UpdateVolumePurchasingSubscriptionByID(subscriptionID, &updateSubscription)
	if err != nil {
		fmt.Printf("Error updating volume purchasing subscription with ID '%s': %s\n", subscriptionID, err)
		return
	}

	// If no error, print the details of the updated subscription.
	fmt.Printf("Updated Subscription ID: %s\n", updatedSubscription.Id)
	fmt.Printf("Name: %s\n", updatedSubscription.Name)
	// ... print other fields as needed
}
