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

	// Define the package IDs to be deleted
	packageIDs := []string{"267", "264"}

	err = client.DeleteMultiplePackagesByID(packageIDs)
	if err != nil {
		log.Fatalf("Error deleting multiple packages by IDs: %v", err)
	} else {
		fmt.Printf("Packages with IDs %v successfully deleted.\n", packageIDs)
	}
}
