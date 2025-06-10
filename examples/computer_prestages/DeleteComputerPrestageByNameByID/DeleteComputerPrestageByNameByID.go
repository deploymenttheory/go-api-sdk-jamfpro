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

	// The name of the computer prestage you want to delete
	prestageName := "YOUR_PRESTAGE_NAME_HERE"

	// Call DeleteComputerPrestageByName to delete the prestage by its name
	err = client.DeleteComputerPrestageByName(prestageName)
	if err != nil {
		log.Fatalf("Error deleting computer prestage by name: %v", err)
	}

	// Print a confirmation message
	fmt.Println("Computer prestage deleted successfully.")
}
