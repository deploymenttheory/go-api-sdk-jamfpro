package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json" // Update the path to your configuration file

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the Jamf Pro client
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new Jamf Pro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Assume we have an ID for the branding we want to update
	brandingID := "2" // Replace with your actual branding ID

	// Define the updated branding details
	updatedBranding := &jamfpro.SelfServiceBrandingDetail{
		ApplicationName:       "Updated App Name",
		BrandingName:          "Updated Branding Name",
		BrandingNameSecondary: "Updated Branding Secondary Name",
		//IconId:                1, // IconId should match the icon image ID uploaded seperately
		//BrandingHeaderImageId: 1, // BrandingHeaderImageId should match the icon image ID uploaded seperately
	}

	// Call the update function with the ID and updated details
	updatedBrandingResponse, err := client.UpdateSelfServiceBrandingMacOSByID(brandingID, updatedBranding)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating self-service branding: %v\n", err)
		os.Exit(1)
	}

	// Print the result in a pretty JSON format
	resultJSON, _ := json.MarshalIndent(updatedBrandingResponse, "", "  ")
	fmt.Println("Self Service Branding updated successfully:")
	fmt.Println(string(resultJSON))
}
