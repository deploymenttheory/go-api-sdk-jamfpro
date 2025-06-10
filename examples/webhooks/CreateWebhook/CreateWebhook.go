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
		Name:               "go-api-sdk-Webhook",
		Enabled:            true,
		URL:                "https://server.com",
		ContentType:        "application/json",
		Event:              "SmartGroupComputerMembershipChange",
		ConnectionTimeout:  5,
		ReadTimeout:        2,
		AuthenticationType: "HEADER",
		// AuthenticationHeaders: `{
		// 	"headers": {
		// 		"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		// 		"Accept-Encoding": "gzip, deflate, br",
		// 		"Accept-Language": "en-US,en;q=0.5",
		// 		"Connection": "keep-alive",
		// 		"Host": "www.example.com",
		// 		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0"
		// 	}
		// }`,
		Username:     "Sample User",
		Password:     "SamplePassword",
		SmartGroupID: 1,
	}

	createdWebhook, err := client.CreateWebhook(newWebhook)
	if err != nil {
		fmt.Printf("Error creating webhook: %v\n", err)
		return
	}
	fmt.Printf("Created Webhook: %+v\n", createdWebhook)
}
