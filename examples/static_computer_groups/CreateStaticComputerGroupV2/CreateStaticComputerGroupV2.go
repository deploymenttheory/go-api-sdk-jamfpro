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

	// Create new static computer group
	newGroup := jamfpro.ResourceStaticComputerGroupV2{
		Name:        "Static Test Group",
		Assignments: &[]string{"27"},
	}

	// Call function
	created, err := client.CreateStaticComputerGroupV2(newGroup)
	if err != nil {
		log.Fatalf("Error creating static computer group: %v", err)
	}

	// Pretty print the JSON
	response, err := json.MarshalIndent(created, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created group data: %v", err)
	}
	fmt.Println("Created Static Computer Group:\n", string(response))
}
