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

	// Sort filter
	sortFilter := "sort=id:desc"

	// Call the GetSelfServiceBrandingMacOS function and handle any errors
	branding, err := client.GetSelfServiceBrandingMacOS(sortFilter)
	if err != nil {
		// If there's an error, log it to stderr and exit with a non-zero status code
		fmt.Fprintf(os.Stderr, "Error fetching self-service branding for macOS: %v\n", err)
		os.Exit(1)
	}

	// If there are no errors, print the retrieved branding information
	fmt.Printf("Total Count: %d\n", branding.TotalCount)
	for _, detail := range branding.Results {
		fmt.Printf("ID: %s, Application Name: %s, Branding Name: %s\n", detail.ID, detail.ApplicationName, detail.BrandingName)
	}
}
