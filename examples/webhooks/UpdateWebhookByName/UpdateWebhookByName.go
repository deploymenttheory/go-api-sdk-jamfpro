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
