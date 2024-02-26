package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Name of the subscription to update
	subscriptionName := "Example Volume Purchasing Subscription"

	// Example update data
	updateData := jamfpro.ResourceVolumePurchasingSubscription{
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

	// Call the update function with the subscription name and update data
	updatedSubscription, err := client.UpdateVolumePurchasingSubscriptionByName(subscriptionName, &updateData)
	if err != nil {
		fmt.Printf("Error updating volume purchasing subscription with name '%s': %s\n", subscriptionName, err)
		return
	}

	// If no error, print the details of the updated subscription
	fmt.Printf("Updated Subscription ID: %s\n", updatedSubscription.Id)
	fmt.Printf("Name: %s\n", updatedSubscription.Name)
	// ... print other fields as needed
}
