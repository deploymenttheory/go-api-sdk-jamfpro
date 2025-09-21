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
	configFilePath := "/Users/Shared/GitHub/go-api-sdk-jamfpro/localtesting/clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Replace with the actual branding ID you want to update
	brandingID := "8"

	// Retrieve current to copy fields you don't want to change (optional)
	current, err := client.GetSelfServiceBrandingIOSByID(brandingID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error retrieving existing branding: %v\n", err)
		os.Exit(1)
	}

	// Modify desired fields
	// Helper to strip leading '#' if present
	stripHash := func(s string) string {
		if len(s) > 0 && s[0] == '#' {
			return s[1:]
		}
		return s
	}

	update := &jamfpro.ResourceSelfServiceBrandingIOSDetail{
		BrandingName:              current.BrandingName,
		IconId:                    current.IconId,
		HeaderBackgroundColorCode: current.HeaderBackgroundColorCode,
		MenuIconColorCode:         stripHash("#000000"), // new menu icon color
		BrandingNameColorCode:     current.BrandingNameColorCode,
		StatusBarTextColor:        current.StatusBarTextColor,
	}

	updated, err := client.UpdateSelfServiceBrandingIOSByID(brandingID, update)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating self-service branding: %v\n", err)
		os.Exit(1)
	}

	resultJSON, _ := json.MarshalIndent(updated, "", "  ")
	fmt.Println("Self Service Branding updated successfully:")
	fmt.Println(string(resultJSON))
}
