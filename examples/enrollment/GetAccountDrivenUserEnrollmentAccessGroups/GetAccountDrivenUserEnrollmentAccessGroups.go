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

	// Call GetAccountDrivenUserEnrollmentAccessGroups function
	// For more information on how to add parameters to this request, see docs/url_queries.md
	ADUEAccessGroups, err := client.GetAccountDrivenUserEnrollmentAccessGroups(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching ADUE Access Groups: %v", err)
	}

	// Pretty print the scripts in XML
	JSON, err := json.MarshalIndent(ADUEAccessGroups, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling scripts data: %v", err)
	}
	fmt.Println("Fetched Scripts:\n", string(JSON))
}
