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

	// Define the group name you want to retrieve
	groupName := "All Managed Clients" // Replace with the actual group name

	// Call GetComputerGroupByJamfProName function
	computerGroup, err := client.GetComputerGroupByJamfProName(groupName)
	if err != nil {
		log.Fatalf("Error fetching computer group by name '%s': %v", groupName, err)
	}

	// Pretty print the computer group in JSON
	groupJSON, err := json.MarshalIndent(computerGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling computer group data: %v", err)
	}
	fmt.Printf("Fetched Computer Group (Name: %s):\n%s\n", groupName, string(groupJSON))
}
