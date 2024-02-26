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

	// Define the bundle ID and version of the mobile device application to delete
	bundleID := "com.example.bundleid" // Replace with actual bundle ID
	version := "1.0.0"                 // Replace with actual version

	// Perform the deletion
	err = client.DeleteMobileDeviceApplicationByBundleIDAndVersion(bundleID, version)
	if err != nil {
		fmt.Println("Error deleting mobile device application by bundle ID and version:", err)
	} else {
		fmt.Println("Successfully deleted mobile device application by bundle ID and version.")
	}
}
