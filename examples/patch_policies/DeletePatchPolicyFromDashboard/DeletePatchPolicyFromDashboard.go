package main

import (
	"fmt"
	"log"

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

	// Specify the patch policy ID to remove from the dashboard
	policyID := "5"

	// Remove the patch policy from the dashboard
	err = client.DeletePatchPolicyFromDashboard(policyID)
	if err != nil {
		log.Fatalf("Error removing patch policy from dashboard: %v", err)
	}

	fmt.Printf("Successfully removed patch policy ID %s from the dashboard\n", policyID)
}
