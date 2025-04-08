package main

import (
	"encoding/json"
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

	displayName := "idp-name"

	cloudIdp, err := client.GetCloudIdentityProviderConfigurationByName(displayName)
	if err != nil {
		log.Fatalf("Error fetching cloud identity provider: %v", err)
	}

	JSON, err := json.MarshalIndent(cloudIdp, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling cloud identity provider data: %v", err)
	}
	fmt.Printf("Cloud Identity Provider (Name: %s):\n%s\n", displayName, string(JSON))
}
