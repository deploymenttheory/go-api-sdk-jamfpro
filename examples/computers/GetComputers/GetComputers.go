package main

import (
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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Call the GetComputers method
	computers, err := client.GetComputers()
	if err != nil {
		log.Fatalf("Error fetching computers: %v", err)
	}

	// Print out the fetched computers
	fmt.Println("Fetched Computers:")
	for _, computer := range computers.Results {
		fmt.Printf("ID: %d, Name: %s\n", computer.ID, computer.Name)
	}
}
