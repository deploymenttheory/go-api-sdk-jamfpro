package main

import (
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

	// Specify the patch policy ID to add to the dashboard
	policyID := "5"

	// Add the patch policy to the dashboard
	err = client.AddPatchPolicyToDashboard(policyID)
	if err != nil {
		log.Fatalf("Error adding patch policy to dashboard: %v", err)
	}

	fmt.Printf("Successfully added patch policy ID %s to the dashboard\n", policyID)
}
