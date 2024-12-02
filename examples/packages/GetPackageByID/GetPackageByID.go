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

	// Example ID to fetch
	packageID := "17"

	response, err := client.GetPackageByID(packageID)
	if err != nil {
		fmt.Println("Error fetching package by ID:", err)
		return
	}

	// Pretty print the created script details in JSON
	packageJSON, err := json.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling package data: %v", err)
	}
	fmt.Println("Obtained package Details:\n", string(packageJSON))
}
