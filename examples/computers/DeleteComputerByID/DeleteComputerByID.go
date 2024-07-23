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

	computerID := "1" // Replace with actual computer ID

	err = client.DeleteComputerByID(computerID)
	if err != nil {
		log.Fatalf("Error deleting computer by ID: %v", err)
	}

	fmt.Printf("Successfully deleted computer with ID: %d\n", computerID)
}
