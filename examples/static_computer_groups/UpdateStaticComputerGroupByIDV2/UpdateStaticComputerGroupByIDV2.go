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

	// Define group ID to update
	groupID := "1202"

	// Create update data
	updateGroup := jamfpro.ResourceStaticComputerGroupV2{
		Name:        "Updated Static Group",
		Description: "Updated description",
	}

	// Call function
	updated, err := client.UpdateStaticComputerGroupByIDV2(groupID, updateGroup)
	if err != nil {
		log.Fatalf("Error updating static computer group: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(updated, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated group data: %v", err)
	}
	fmt.Println("Updated Static Computer Group:\n", string(response))
}
