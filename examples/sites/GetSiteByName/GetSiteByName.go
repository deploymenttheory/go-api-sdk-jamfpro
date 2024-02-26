package main

import (
	"fmt"
	"log"

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

	siteName := "Site Name" // Replace with the actual site name

	// Fetch the site by name
	site, err := client.GetSiteByName(siteName)
	if err != nil {
		log.Fatalf("Error fetching site by name: %v", err)
	}

	fmt.Printf("Site ID: %d, Name: %s\n", site.ID, site.Name)
}
