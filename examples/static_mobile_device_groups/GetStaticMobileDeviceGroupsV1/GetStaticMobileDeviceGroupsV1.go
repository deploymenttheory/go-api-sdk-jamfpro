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
	params.Add("sort", "groupName:asc")

	// Call function
	groups, err := client.GetStaticMobileDeviceGroupsV1(params)
	if err != nil {
		log.Fatalf("Error fetching static mobile device groups v1: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(groups, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling groups data: %v", err)
	}
	fmt.Println("Fetched Static Mobile Device Groups v1:\n", string(response))
}
