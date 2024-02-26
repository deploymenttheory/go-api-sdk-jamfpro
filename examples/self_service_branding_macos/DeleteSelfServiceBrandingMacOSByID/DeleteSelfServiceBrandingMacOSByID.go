package main

import (
	"fmt"
	"log"
	"os"

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
	// Assume we have an ID for the branding we want to delete
	brandingID := "2" // Replace with your actual branding ID

	// Call the delete function with the ID
	err = client.DeleteSelfServiceBrandingMacOSByID(brandingID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error deleting self-service branding: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Self Service Branding deleted successfully")
}
