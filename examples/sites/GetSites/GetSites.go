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

	// Fetch all sites
	sites, err := client.GetSites()
	if err != nil {
		log.Fatalf("Error fetching sites: %v", err)
	}

	fmt.Println("Sites:")
	for _, site := range sites.Site {
		fmt.Printf("ID: %d, Name: %s\n", site.ID, site.Name)
	}
}
