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

	iconId := 5
	// API expects 6 hex characters (RRGGBB) without a leading '#'. Allow examples to include '#', but strip it before sending.
	stripHash := func(s string) string {
		if len(s) > 0 && s[0] == '#' {
			return s[1:]
		}
		return s
	}

	newBranding := jamfpro.ResourceSelfServiceBrandingIOSDetail{
		BrandingName:              "iOS Self Service",
		IconId:                    &iconId,              // IconId should match the icon image ID uploaded separately
		HeaderBackgroundColorCode: stripHash("#0066CC"), // Background color for the header
		MenuIconColorCode:         stripHash("#FFFFFF"), // Color for menu icons
		BrandingNameColorCode:     stripHash("#FFFFFF"), // Color for the branding name text
		StatusBarTextColor:        "light",              // Status bar text color: "light" or "dark"
	}

	createdBranding, err := client.CreateSelfServiceBrandingIOS(&newBranding)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating self-service branding for iOS: %v\n", err)
		os.Exit(1)
	}
	// The SDK now returns a minimal create response containing id and href.
	// Print the minimal JSON response.
	out, _ := json.MarshalIndent(createdBranding, "", "  ")
	fmt.Println(string(out))
}
