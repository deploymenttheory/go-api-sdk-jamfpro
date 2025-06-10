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
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Assume we have an ID for the branding we want to update
	brandingID := "2" // Replace with your actual branding ID

	// Define the updated branding details
	updatedBranding := &jamfpro.ResourceSelfServiceBrandingDetail{
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
