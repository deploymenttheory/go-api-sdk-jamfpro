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

	// Assume we have a name and the new details for the branding we want to update
	brandingName := "Self Service" // Branding is always called Self Service. there is only 1 as well

	newBrandingDetails := &jamfpro.ResourceSelfServiceBrandingDetail{
		// Set the new details for the branding
		ApplicationName:       "Updated Self Service",
		BrandingName:          "Updated Self Service",
		BrandingNameSecondary: "Updated Self Service Secondary",
		//IconId:                2, // New icon ID
		//BrandingHeaderImageId: 2, // New header image ID
	}

	// Call the update by name function
	updatedBranding, err := client.UpdateSelfServiceBrandingMacOSByName(brandingName, newBrandingDetails)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating self-service branding: %v\n", err)
		os.Exit(1)
	}

	// If there are no errors, print the updated branding information
	fmt.Printf("Updated Branding ID: %s, Application Name: %s, Branding Name: %s\n",
		updatedBranding.ID, updatedBranding.ApplicationName, updatedBranding.BrandingName)
}
