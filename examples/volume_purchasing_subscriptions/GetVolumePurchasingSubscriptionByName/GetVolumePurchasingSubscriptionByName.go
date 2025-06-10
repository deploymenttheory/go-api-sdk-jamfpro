package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Replace 'subscriptionName' with the actual name of the subscription you want to retrieve.
	subscriptionName := "Example"

	// Call the GetVolumePurchasingSubscriptionByName function with the subscription name.
	subscription, err := client.GetVolumePurchasingSubscriptionByName(subscriptionName)
	if err != nil {
		fmt.Printf("Error fetching volume purchasing subscription with name '%s': %s\n", subscriptionName, err)
		return
	}

	// If no error, print the details of the retrieved subscription.
	fmt.Printf("Subscription ID: %s\n", subscription.Id)
	fmt.Printf("Name: %s\n", subscription.Name)
	// ... print other fields as needed
}
