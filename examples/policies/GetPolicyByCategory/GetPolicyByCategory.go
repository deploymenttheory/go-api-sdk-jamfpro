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

	// Define a policy category for testing
	policyCategory := "YourPolicyCategory"

	// Call GetPolicyByCategory function
	policies, err := client.GetPolicyByCategory(policyCategory)
	if err != nil {
		log.Fatalf("Error fetching policies by category: %v", err)
	}

	// Pretty print the policies in XML
	policiesXML, err := xml.MarshalIndent(policies, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling policies data: %v", err)
	}
	fmt.Println("Fetched Policies by Category:\n", string(policiesXML))
}
