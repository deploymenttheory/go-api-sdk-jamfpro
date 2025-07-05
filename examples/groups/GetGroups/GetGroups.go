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

	// Call GetGroups function
	// For more information on how to add parameters to this request, see docs/url_queries.md
	groups, err := client.GetGroups(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching groups: %v", err)
	}

	// Pretty print the groups in JSON
	groupsJSON, err := json.MarshalIndent(groups, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling groups data: %v", err)
	}
	fmt.Println("Fetched Groups:\n", string(groupsJSON))
}
