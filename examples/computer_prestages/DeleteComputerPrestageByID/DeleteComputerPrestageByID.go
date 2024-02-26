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

	// The ID of the computer prestage you want to delete
	prestageID := "1"

	// Call DeleteComputerPrestageByID to delete the prestage
	err = client.DeleteComputerPrestageByID(prestageID)
	if err != nil {
		log.Fatalf("Error deleting computer prestage: %v", err)
	}

	// Print a confirmation message
	fmt.Println("Computer prestage deleted successfully.")
}
