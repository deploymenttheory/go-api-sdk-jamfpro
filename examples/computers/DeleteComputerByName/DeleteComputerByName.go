package main

import (
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

	computerName := "admin’s MacBook Pro" // Replace with actual computer name

	err = client.DeleteComputerByName(computerName)
	if err != nil {
		log.Fatalf("Error deleting computer by name: %v", err)
	}

	fmt.Printf("Successfully deleted computer with name: %s\n", computerName)
}
