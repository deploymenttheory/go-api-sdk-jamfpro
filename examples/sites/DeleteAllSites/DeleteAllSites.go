package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

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

	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Fetch all sites
	sites, err := client.GetSites()
	if err != nil {
		log.Fatalf("Error fetching sites: %v", err)
	}

	fmt.Println("Sites fetched. Starting deletion process:")

	// Iterate over each site and delete
	for _, site := range sites.Site {
		fmt.Printf("Deleting site ID: %d, Name: %s\n", site.ID, site.Name)

		err = client.DeleteSiteByID(site.ID)
		if err != nil {
			log.Printf("Error deleting site ID %d: %v\n", site.ID, err)
			continue // Move to the next site if there's an error
		}

		fmt.Printf("Site ID %d deleted successfully.\n", site.ID)
	}

	fmt.Println("Site deletion process completed.")
}
