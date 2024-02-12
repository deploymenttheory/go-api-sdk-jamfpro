package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientconfig.json"
	// Load the client OAuth credentials from the configuration file
	loadedConfig, err := jamfpro.LoadClientConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the HTTP client
	config := httpclient.ClientConfig{
		Auth: httpclient.AuthConfig{
			ClientID:     loadedConfig.Auth.ClientID,
			ClientSecret: loadedConfig.Auth.ClientSecret,
		},
		Environment: httpclient.EnvironmentConfig{
			APIType:      loadedConfig.Environment.APIType,
			InstanceName: loadedConfig.Environment.InstanceName,
		},
		ClientOptions: httpclient.ClientOptions{
			LogLevel:          loadedConfig.ClientOptions.LogLevel,
			HideSensitiveData: loadedConfig.ClientOptions.HideSensitiveData,
			LogOutputFormat:   loadedConfig.ClientOptions.LogOutputFormat,
		},
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
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
