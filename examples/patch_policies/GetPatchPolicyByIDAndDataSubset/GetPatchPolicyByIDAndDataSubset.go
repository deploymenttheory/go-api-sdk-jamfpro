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

	patchPolicyID := 1  // Example ID
	subset := "general" // Replace with the desired subset name

	patchPolicy, err := client.GetPatchPolicyByIDAndDataSubset(patchPolicyID, subset)
	if err != nil {
		log.Fatalf("Error fetching patch policy by ID and subset: %v", err)
	}

	output, err := xml.MarshalIndent(patchPolicy, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling patch policy to XML: %v", err)
	}

	fmt.Printf("Patch Policy Subset (ID: %d, Subset: %s):\n%s\n", patchPolicyID, subset, string(output))
}
