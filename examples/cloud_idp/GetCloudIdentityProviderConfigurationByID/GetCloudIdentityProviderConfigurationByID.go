package main

import (
	"encoding/json"
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

	// Define the ID of the cloud identity provider to retrieve
	cloudIdpID := "1001"

	// Call GetCloudIdentityProviderByID function
	cloudIdp, err := client.GetCloudIdentityProviderConfigurationByID(cloudIdpID)
	if err != nil {
		log.Fatalf("Error fetching cloud identity provider: %v", err)
	}

	// Pretty print the cloud identity provider in JSON
	JSON, err := json.MarshalIndent(cloudIdp, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling cloud identity provider data: %v", err)
	}
	fmt.Printf("Cloud Identity Provider (ID: %s):\n%s\n", cloudIdpID, string(JSON))
}
