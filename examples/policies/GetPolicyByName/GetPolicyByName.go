package main

import (
	"encoding/xml"
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

	// Define a policy name for testing
	policyName := "test"

	// Call GetPolicyByName function
	policy, err := client.GetPolicyByName(policyName)
	if err != nil {
		log.Fatalf("Error fetching policy by name: %v", err)
	}

	// Pretty print the policy details in XML
	policyXML, err := xml.MarshalIndent(policy, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling policy details data: %v", err)
	}
	fmt.Println("Fetched Policy Details:\n", string(policyXML))
}
