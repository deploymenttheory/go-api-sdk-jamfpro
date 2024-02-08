package main

import (
	"fmt"
	"log"
	"os"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
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
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
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
