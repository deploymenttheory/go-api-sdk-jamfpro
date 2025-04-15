package main

import (
	"fmt"
	"log"
	"net/url"

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

	// Call the function with desired parameters
	subscriptions, err := client.GetVolumePurchasingSubscriptions(url.Values{})
	if err != nil {
		fmt.Printf("Error fetching volume purchasing subscriptions: %s\n", err)
		return
	}

	// Print out the retrieved subscriptions
	fmt.Printf("Total Subscriptions: %d\n", *subscriptions.TotalCount)
	for _, subscription := range subscriptions.Results {
		fmt.Printf("Subscription ID: %s, Name: %s\n", subscription.Id, subscription.Name)
	}
}
