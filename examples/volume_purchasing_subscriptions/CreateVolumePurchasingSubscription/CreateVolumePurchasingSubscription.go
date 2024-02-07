package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
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
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
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

	// Define a new subscription to create.
	newSubscription := jamfpro.ResourceVolumePurchasingSubscription{
		Enabled:  true,
		SiteId:   "-1",
		Name:     "Example Volume Purchasing Subscription",
		Triggers: []string{"NO_MORE_LICENSES", "REMOVED_FROM_APP_STORE"},
		InternalRecipients: []jamfpro.VolumePurchasingSubscriptionSubsetInternalRecipients{
			{Frequency: "DAILY", AccountId: "1"},
		},
		ExternalRecipients: []jamfpro.VolumePurchasingSubscriptionSubsetExternalRecipients{
			{Name: "Bob Exampleman", Email: "bob@example.com"},
		},
	}

	// Call the CreateVolumePurchasingSubscription function with the new subscription data.
	createdSubscription, err := client.CreateVolumePurchasingSubscription(&newSubscription)
	if err != nil {
		fmt.Printf("Error creating volume purchasing subscription: %s\n", err)
		return
	}

	// If no error, print the details of the created subscription.
	fmt.Printf("Created Subscription ID: %s\n", createdSubscription.Id)
	fmt.Printf("Name: %s\n", createdSubscription.Name)
	// ... print other fields as needed
}
