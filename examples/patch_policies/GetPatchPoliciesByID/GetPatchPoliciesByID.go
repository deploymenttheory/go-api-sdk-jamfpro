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

	// The ID of the patch policy you want to retrieve
	patchPolicyID := "1" // Replace with the actual ID you want to retrieve

	// Call the GetPatchPoliciesByID function
	patchPolicy, err := client.GetPatchPoliciesByID(patchPolicyID)
	if err != nil {
		log.Fatalf("Error fetching patch policy by ID: %v", err)
	}

	// Convert the response into pretty XML for printing
	output, err := xml.MarshalIndent(patchPolicy, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling patch policy to XML: %v", err)
	}

	// Print the pretty XML
	fmt.Printf("Patch Policy (ID: %d):\n%s\n", patchPolicyID, string(output))
}
