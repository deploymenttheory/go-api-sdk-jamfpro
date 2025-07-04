package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Define the group ID you want to retrieve
	groupID := "ff5a53d5-e585-4f76-96b8-aaf2d2945b94" // Replace with the actual group ID

	// Call GetGroupByID function
	group, err := client.GetGroupByID(groupID)
	if err != nil {
		log.Fatalf("Error fetching group by ID %s: %v", groupID, err)
	}

	// Pretty print the group in JSON
	groupJSON, err := json.MarshalIndent(group, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling group data: %v", err)
	}
	fmt.Printf("Fetched Group (ID: %s):\n%s\n", groupID, string(groupJSON))
}
