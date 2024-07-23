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

	// Example patch policy ID
	patchPolicyID := "1" // Replace with an actual patch policy ID

	// Call DeletePatchPolicy function
	err = client.DeletePatchPolicyByID(patchPolicyID)
	if err != nil {
		log.Fatalf("Error deleting patch policy: %v", err)
	}

	fmt.Printf("Patch policy with ID %d deleted successfully.\n", patchPolicyID)
}
