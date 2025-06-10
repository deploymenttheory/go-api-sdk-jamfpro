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
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Call GetPatchPolicies function to fetch policies
	// For more information on how to add parameters to this request, see docs/url_queries.md
	policies, err := client.GetPatchPolicies(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching patch policies: %v", err)
	}

	// Pretty print the JSON response
	response, err := json.MarshalIndent(policies, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling patch policies data: %v", err)
	}
	fmt.Println("Fetched patch policies:\n", string(response))
}
