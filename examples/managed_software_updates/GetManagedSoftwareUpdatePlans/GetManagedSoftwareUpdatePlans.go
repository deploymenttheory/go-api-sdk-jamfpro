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

	// Call GetManagedSoftwareUpdatePlans function
	updatePlans, err := client.GetManagedSoftwareUpdatePlans(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching managed software update plans: %v", err)
	}

	// Pretty print the managed software update plans in json
	// For more information on how to add parameters to this request, see docs/url_queries.md
	updatePlansJSON, err := json.MarshalIndent(updatePlans, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling managed software update plans data: %v", err)
	}
	fmt.Println("Fetched managed software update plans:\n", string(updatePlansJSON))
}
