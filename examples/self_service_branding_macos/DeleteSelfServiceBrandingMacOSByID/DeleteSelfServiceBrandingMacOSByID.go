package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json" // Update the path to your configuration file

	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	client, err := jamfpro.NewClient(config)
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
