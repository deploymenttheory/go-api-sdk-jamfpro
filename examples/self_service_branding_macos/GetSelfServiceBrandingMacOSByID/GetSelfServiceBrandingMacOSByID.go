package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json" // Update the path to your configuration file

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the ID of the self-service branding to fetch
	brandingID := "1" // Replace with the actual ID

	// Call the GetSelfServiceBrandingMacOSByID function and handle any errors
	brandingDetail, err := client.GetSelfServiceBrandingMacOSByID(brandingID)
	if err != nil {
		// If there's an error, log it to stderr and exit with a non-zero status code
		fmt.Fprintf(os.Stderr, "Error fetching self-service branding for macOS with ID %s: %v\n", brandingID, err)
		os.Exit(1)
	}

	// If there are no errors, print the retrieved branding information
	fmt.Printf("Branding ID: %s\n", brandingDetail.ID)
	fmt.Printf("Application Name: %s\n", brandingDetail.ApplicationName)
	fmt.Printf("Branding Name: %s\n", brandingDetail.BrandingName)
	fmt.Printf("Branding Name Secondary: %s\n", brandingDetail.BrandingNameSecondary)
	fmt.Printf("Icon ID: %d\n", brandingDetail.IconId)
	fmt.Printf("Branding Header Image ID: %d\n", brandingDetail.BrandingHeaderImageId)
}
