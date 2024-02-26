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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Define the ID of the self-service branding to fetch
	brandingID := "1" // Replace with the actual ID

	// Call the GetSelfServiceBrandingMacOSByID function and handle any errors
	brandingDetail, err := client.GetSelfServiceBrandingMacOSByID(brandingID)
	if err != nil {
		// If there's an error, log it to stderr and exit with a non-zero status code
		fmt.Fprintf(os.Stderr, "Error fetching self-service branding for macOS with ID %s: %v\n", brandingID, err)
		os.Exit(1)
	}

	// If there are no errors, print the retrieved branding information
	fmt.Printf("Branding ID: %s\n", brandingDetail.ID)
	fmt.Printf("Application Name: %s\n", brandingDetail.ApplicationName)
	fmt.Printf("Branding Name: %s\n", brandingDetail.BrandingName)
	fmt.Printf("Branding Name Secondary: %s\n", brandingDetail.BrandingNameSecondary)
	fmt.Printf("Icon ID: %d\n", brandingDetail.IconId)
	fmt.Printf("Branding Header Image ID: %d\n", brandingDetail.BrandingHeaderImageId)
}
