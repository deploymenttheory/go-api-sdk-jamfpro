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
	newBranding := jamfpro.ResourceSelfServiceBrandingDetail{
		ApplicationName:       "Self Service",
		BrandingName:          "Self Service",
		BrandingNameSecondary: "Self Service",
		//IconId:                1, // IconId should match the icon image ID uploaded seperately
		//BrandingHeaderImageId: 1, // BrandingHeaderImageId should match the icon image ID uploaded seperately
	}

	createdBranding, err := client.CreateSelfServiceBrandingMacOS(&newBranding)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating self-service branding for macOS: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Branding ID: %s\n", createdBranding.ID)
	fmt.Printf("Application Name: %s\n", createdBranding.ApplicationName)
	fmt.Printf("Branding Name: %s\n", createdBranding.BrandingName)
	fmt.Printf("Branding Name Secondary: %s\n", createdBranding.BrandingNameSecondary)
	fmt.Printf("Icon ID: %d\n", createdBranding.IconId)
	fmt.Printf("Branding Header Image ID: %d\n", createdBranding.BrandingHeaderImageId)
}
