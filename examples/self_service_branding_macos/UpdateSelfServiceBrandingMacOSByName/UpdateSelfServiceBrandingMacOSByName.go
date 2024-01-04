package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
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
