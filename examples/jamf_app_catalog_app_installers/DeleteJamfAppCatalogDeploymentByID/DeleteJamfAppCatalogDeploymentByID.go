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

	// Define the ID of the Jamf App Catalog Deployment you want to delete
	resourceID := "8"

	// Call the DeleteJamfAppCatalogAppInstallerDeploymentByID function
	err = client.DeleteJamfAppCatalogAppInstallerDeploymentByID(resourceID)
	if err != nil {
		log.Fatalf("Failed to delete Jamf App Catalog Deployment with ID %s: %v", resourceID, err)
	}

	fmt.Printf("Jamf App Catalog Deployment with ID %s deleted successfully\n", resourceID)
}
