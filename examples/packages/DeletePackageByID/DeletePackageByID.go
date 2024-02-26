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

	packageID := 1

	err = client.DeletePackageByID(packageID)
	if err != nil {
		log.Fatalf("Error deleting network package by ID: %v", err)
	} else {
		fmt.Printf("Network package with ID %d successfully deleted.\n", packageID)
	}
}
