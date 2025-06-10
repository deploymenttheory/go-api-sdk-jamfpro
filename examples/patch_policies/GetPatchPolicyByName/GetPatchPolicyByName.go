package main

import (
	"encoding/json"
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

	// Specify the patch policy name to retrieve
	policyName := "Exceptions Group Ring 5" // Example name from previous logs

	// Get the patch policy by name
	policy, err := client.GetPatchPolicyByName(policyName)
	if err != nil {
		log.Fatalf("Error getting patch policy: %v", err)
	}

	// Pretty print the JSON response
	response, err := json.MarshalIndent(policy, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling patch policy data: %v", err)
	}
	fmt.Printf("Found patch policy with name '%s':\n%s\n", policyName, string(response))
}
