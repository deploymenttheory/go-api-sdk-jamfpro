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

	// Update the SSO failover URL
	updatedFailoverSettings, err := client.UpdateFailoverUrl()
	if err != nil {
		log.Fatalf("Error updating SSO failover URL: %v", err)
	}

	fmt.Println("Updated SSO Failover URL:", updatedFailoverSettings.FailoverURL)
	fmt.Println("New Generation Time:", updatedFailoverSettings.GenerationTime)
}
