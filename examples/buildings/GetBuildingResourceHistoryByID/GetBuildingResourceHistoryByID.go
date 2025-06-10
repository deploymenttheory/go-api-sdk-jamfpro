package main

import (
	"fmt"
	"log"
	"net/url"

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

	// Example: Fetch the resource history of a building by ID
	// For more information on how to add parameters to this request, see docs/url_queries.md
	buildingID := "" // Replace with a real building ID
	history, err := client.GetBuildingResourceHistoryByID(buildingID, url.Values{})
	if err != nil {
		log.Fatalf("Error fetching building resource history: %v", err)
	}

	// Print the resource history
	fmt.Printf("Resource History for Building ID %s:\n", buildingID)
	for _, record := range history.Results {
		fmt.Printf("ID: %d, Username: %s, Date: %s, Note: %s, Details: %s\n",
			record.ID, record.Username, record.Date, record.Note, record.Details)
	}
}
