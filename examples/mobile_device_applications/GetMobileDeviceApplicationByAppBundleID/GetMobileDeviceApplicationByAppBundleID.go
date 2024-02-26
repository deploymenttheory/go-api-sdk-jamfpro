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

	// Replace "YourBundleID" with an actual bundle ID
	bundleID := "YourBundleID"

	appByBundleID, err := client.GetMobileDeviceApplicationByAppBundleID(bundleID)
	if err != nil {
		fmt.Println("Error fetching application by Bundle ID:", err)
	} else {
		fmt.Println("Fetched Application by Bundle ID:", appByBundleID)
	}
}
