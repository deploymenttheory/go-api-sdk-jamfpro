package main

import (
	"encoding/json"
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

	// The ID of the patch policy you want to retrieve
	patchPolicyID := "1" // Replace with the actual ID you want to retrieve

	// Call the GetPatchPoliciesByID function
	response, err := client.GetPatchPoliciesByID(patchPolicyID)
	if err != nil {
		log.Fatalf("Error fetching patch policy by ID: %v", err)
	}

	// Pretty print the created script details in JSON
	packageJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling package data: %v", err)
	}
	fmt.Println("Obtained patch policy Details:\n", string(packageJSON))
}
