package main

import (
	"encoding/json"
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

	// Define the ID of the computer inventory you want to retrieve
	computerInventoryID := "14"

	// Call the GetComputerInventoryByID function
	computerInventory, err := client.GetComputerInventoryByID(computerInventoryID)
	if err != nil {
		log.Fatalf("Error fetching computer inventory by ID: %v", err)
	}

	// Pretty print the response
	prettyJSON, err := json.MarshalIndent(computerInventory, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty JSON: %v", err)
	}
	fmt.Printf("%s\n", prettyJSON)
}
