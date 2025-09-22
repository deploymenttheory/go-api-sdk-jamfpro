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

	// Replace with the actual branding ID you want to fetch
	brandingID := "8"
	branding, err := client.GetSelfServiceBrandingIOSByID(brandingID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching self-service branding for iOS with ID %s: %v\n", brandingID, err)
		os.Exit(1)
	}

	brandingJSON, err := json.MarshalIndent(branding, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling branding data: %v", err)
	}
	fmt.Println(string(brandingJSON))
}
