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

	// Branding name to update
	brandingName := "Example iOS Branding" // Replace with actual branding name

	// Prepare update payload (only fields you want to change)
	stripHash := func(s string) string {
		if len(s) > 0 && s[0] == '#' {
			return s[1:]
		}
		return s
	}

	update := &jamfpro.ResourceSelfServiceBrandingIOSDetail{
		BrandingName:      brandingName,
		MenuIconColorCode: stripHash("#001100"),
	}

	updated, err := client.UpdateSelfServiceBrandingIOSByName(brandingName, update)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error updating self-service branding: %v\n", err)
		os.Exit(1)
	}

	updatedJSON, err := json.MarshalIndent(updated, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated branding data: %v", err)
	}

	fmt.Println("Self Service Branding updated:")
	fmt.Println(string(updatedJSON))
}
