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

	// Example patch policy ID to delete
	policyID := "1"

	// Delete the patch policy
	err = client.DeletePatchPolicyByID(policyID)
	if err != nil {
		log.Fatalf("Error deleting patch policy: %v", err)
	}

	fmt.Printf("Successfully deleted patch policy with ID: %s\n", policyID)
}
