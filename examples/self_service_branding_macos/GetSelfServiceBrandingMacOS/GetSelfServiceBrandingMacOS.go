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
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
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
