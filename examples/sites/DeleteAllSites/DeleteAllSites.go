package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

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

	fmt.Println("Sites fetched. Starting deletion process:")

	// Iterate over each site and delete
	for _, site := range sites.Site {
		fmt.Printf("Deleting site ID: %d, Name: %s\n", site.ID, site.Name)

		err = client.DeleteSiteByID(strconv.Itoa(site.ID))
		if err != nil {
			log.Printf("Error deleting site ID %d: %v\n", site.ID, err)
			continue // Move to the next site if there's an error
		}

		fmt.Printf("Site ID %d deleted successfully.\n", site.ID)
	}

	fmt.Println("Site deletion process completed.")
}
