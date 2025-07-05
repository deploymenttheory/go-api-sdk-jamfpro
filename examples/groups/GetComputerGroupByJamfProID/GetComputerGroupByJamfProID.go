package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Path to the Jamf Pro client configuration file
	configPath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client
	jamfClient, err := jamfpro.BuildClientWithConfigFile(configPath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// The Jamf Pro Computer Group ID to retrieve
	computerGroupID := "1" // Replace with the actual Computer Group Jamf Pro ID

	// Fetch the computer group by Jamf Pro ID
	computerGroup, err := jamfClient.GetComputerGroupByJamfProID(computerGroupID)
	if err != nil {
		log.Fatalf("Error fetching computer group by Jamf Pro ID '%s': %v", computerGroupID, err)
	}

	// Pretty print the computer group in JSON
	groupJSON, err := json.MarshalIndent(computerGroup, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling computer group data: %v", err)
	}
	fmt.Printf("Fetched Computer Group (Jamf Pro ID: %s):\n%s\n", computerGroupID, string(groupJSON))
}
