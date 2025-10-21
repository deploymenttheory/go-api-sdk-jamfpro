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
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Set sorting filter (optional)
	// For more information on how to add parameters to this request, see docs/url_queries.md

	params := url.Values{}
	params.Add("sort", "name:asc")

	// Call function
	groups, err := client.GetStaticComputerGroupsV2(params)
	if err != nil {
		log.Fatalf("Error fetching static computer groups v2: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(groups, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling groups data: %v", err)
	}
	fmt.Println("Fetched Static Computer Groups V2:\n", string(response))
}
