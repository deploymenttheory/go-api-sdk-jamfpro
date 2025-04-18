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
	configFilePath := "/Users/neilmartin/GitHub/go-api-sdk-jamfpro/client_auth.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define sorting parameters
	// For more information on how to add parameters to this request, see docs/url_queries.md
	params := url.Values{}
	params.Add("sort", "id")

	// Fetch mobile device prestages using the V3 API
	prestages, err := client.GetMobileDevicePrestages(params)
	if err != nil {
		log.Fatalf("Error fetching mobile device prestages: %v", err)
	}

	// Pretty print the mobile device prestage in JSON
	prestageJSON, err := json.MarshalIndent(prestages, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling mobile device prestage data: %v", err)
	}
	fmt.Println("Fetched mobile device prestage:\n", string(prestageJSON))
}
