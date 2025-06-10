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

	// Example software title ID and sort parameters
	// For more information on how to add parameters to this request, see docs/url_queries.md
	titleID := "14" // Example ID from previous logs
	params := url.Values{}
	params.Add("sort", "absoluteOrderId:asc")

	// Get the patch software title definitions
	definitions, err := client.GetPatchSoftwareTitleDefinitions(titleID, params)
	if err != nil {
		log.Fatalf("Error getting patch software title definitions: %v", err)
	}

	// Pretty print the JSON response
	response, err := json.MarshalIndent(definitions, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling patch software title definitions data: %v", err)
	}
	fmt.Printf("Found patch software title definitions for ID %s:\n%s\n", titleID, string(response))
}
