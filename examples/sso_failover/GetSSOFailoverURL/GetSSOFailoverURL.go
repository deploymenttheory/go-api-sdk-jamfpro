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

	// Fetch SSO failover settings
	failoverSettings, err := client.GetSSOFailoverSettings()
	if err != nil {
		log.Fatalf("Error fetching SSO failover settings: %v", err)
	}

	fmt.Printf("Failover URL: %s\n", failoverSettings.FailoverURL)
	fmt.Printf("Generation Time: %d\n", failoverSettings.GenerationTime)
}
