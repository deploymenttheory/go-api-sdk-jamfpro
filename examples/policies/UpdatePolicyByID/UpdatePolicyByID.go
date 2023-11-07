package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
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

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the policy ID to be updated
	policyID := 6 // Replace with the actual policy ID

	// Fetch the existing policy by ID
	policy, err := client.GetPolicyByID(policyID)
	if err != nil {
		log.Fatalf("Error fetching policy: %v", err)
	}

	// Make changes to the policy
	policy.General.Name = "Updated Policy Name"
	policy.General.Enabled = true

	// Update the policy
	updatedPolicy, err := client.UpdatePolicyByID(policyID, policy)
	if err != nil {
		log.Fatalf("Error updating policy: %v", err)
	}

	// Print the updated policy details
	fmt.Printf("Updated Policy: %+v\n", updatedPolicy)
}
