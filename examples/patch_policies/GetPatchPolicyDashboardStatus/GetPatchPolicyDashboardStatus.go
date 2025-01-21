package main

import (
	"encoding/json"
	"fmt"
	"log"

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

	// Specify the patch policy ID to check
	policyID := "5" // Example ID from your logs

	// Call GetPatchPolicyDashboardStatus function
	status, err := client.GetPatchPolicyDashboardStatus(policyID)
	if err != nil {
		log.Fatalf("Error checking patch policy dashboard status: %v", err)
	}

	// Pretty print the JSON response
	response, err := json.MarshalIndent(status, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling patch policy dashboard status data: %v", err)
	}
	fmt.Printf("Dashboard status for patch policy ID %s:\n%s\n", policyID, string(response))
}
