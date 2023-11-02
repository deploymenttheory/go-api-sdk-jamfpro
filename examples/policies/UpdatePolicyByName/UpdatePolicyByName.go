package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the policy name to be updated
	policyName := "test" // Replace with the actual policy name

	// Fetch the existing policy by name
	policy, err := client.GetPolicyByName(policyName)
	if err != nil {
		log.Fatalf("Error fetching policy: %v", err)
	}

	// Make changes to the policy
	policy.General.Enabled = true
	policy.SelfService.UseForSelfService = true
	policy.SelfService.SelfServiceDisplayName = "Install Firefox"

	// Update the policy
	updatedPolicy, err := client.UpdatePolicyByName(policyName, policy)
	if err != nil {
		log.Fatalf("Error updating policy: %v", err)
	}

	// Print the updated policy details
	fmt.Printf("Updated Policy: %+v\n", updatedPolicy)
}
