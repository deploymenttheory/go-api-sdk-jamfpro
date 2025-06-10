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

	// Call GetJamfConnectConfigProfiles function to fetch profiles
	// For more information on how to add parameters to this request, see docs/url_queries.md
	profiles, err := client.GetJamfConnectConfigProfiles(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching Jamf Connect config profiles: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(profiles, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Jamf Connect config profiles data: %v", err)
	}
	fmt.Println("Fetched Jamf Connect config profiles:\n", string(response))
}
