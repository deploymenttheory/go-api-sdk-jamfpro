package main

import (
	"encoding/xml"
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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Define the createdBy parameter for testing
	createdBy := "jss" // Can be either 'casper' or 'jss'

	// Call GetPoliciesByType function
	policies, err := client.GetPoliciesByType(createdBy)
	if err != nil {
		log.Fatalf("Error fetching policies by type: %v", err)
	}

	// Pretty print the policies details in XML
	policiesXML, err := xml.MarshalIndent(policies, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling policies details data: %v", err)
	}
	fmt.Println("Fetched Policies Details:\n", string(policiesXML))
}
