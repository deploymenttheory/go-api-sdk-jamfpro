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

	policies, err := client.GetPolicies()
	if err != nil {
		log.Fatalf("Error fetching policies: %v", err)
	}

	fmt.Println("Retrieved Policies:")
	for _, policy := range policies.Policy {
		fmt.Printf("ID: %d, Name: %s\n", policy.ID, policy.Name)
	}
}
