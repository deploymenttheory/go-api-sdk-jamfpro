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

	newBranding := jamfpro.SelfServiceBrandingDetail{
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
