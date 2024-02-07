package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
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
