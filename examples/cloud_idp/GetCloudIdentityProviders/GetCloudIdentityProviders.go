package main

import (
	"encoding/json"
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

	// Optional: Define sort filter (e.g., "id:asc" or "displayName:desc")
	params := url.Values{}

	// Call GetCloudIdentityProviders function
	cloudIdps, err := client.GetCloudIdentityProviders(params)
	if err != nil {
		log.Fatalf("Error fetching cloud identity providers: %v", err)
	}

	// Pretty print the cloud identity providers in JSON
	JSON, err := json.MarshalIndent(cloudIdps, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling cloud identity providers data: %v", err)
	}
	fmt.Println("Cloud Identity Providers:\n", string(JSON))
}
