package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

// Define the original site name and updated name as constants
const originalSiteName = "OriginalSiteName" // Replace "OriginalSiteName" with the current name of the site you want to update
const updatedSiteName = "UpdatedSiteName"   // Replace "UpdatedSiteName" with the new name for the site

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the updated site data
	updatedSite := &jamfpro.SharedResourceSite{
		Name: updatedSiteName,
	}

	// Update the site by name
	updated, err := client.UpdateSiteByName(originalSiteName, updatedSite)
	if err != nil {
		log.Fatalf("Error updating site by name: %v", err)
	}

	fmt.Printf("Site updated successfully. ID: %d, Name: %s\n", updated.ID, updated.Name)
}
