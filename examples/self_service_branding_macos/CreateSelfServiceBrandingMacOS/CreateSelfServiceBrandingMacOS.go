package main

import (
	"fmt"
	"log"
	"os"

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

	newBranding := jamfpro.ResourceSelfServiceBrandingDetail{
		ApplicationName:       "Self Service", // Name to display for the application in Finder, Dock, and the menu bar (e.g., "My Company's Self Service")
		BrandingName:          "test",
		BrandingNameSecondary: "Self Service",
		HomeHeading:           "Self Service",
		HomeSubheading:        "Self Service",
		IconId:                5, // IconId should match the icon image ID uploaded seperately
		BrandingHeaderImageId: 6, // BrandingHeaderImageId should match the icon image ID uploaded seperately
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
