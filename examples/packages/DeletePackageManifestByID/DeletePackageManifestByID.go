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

	// Define the package ID to be deleted
	packageID := "250"

	err = client.DeletePackageManifestByID(packageID)
	if err != nil {
		log.Fatalf("Error deleting packages by IDs: %v", err)
	} else {
		fmt.Printf("Package manfest with ID %v successfully deleted.\n", packageID)
	}
}
