package main

import (
	"encoding/json"
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

	// For more information on how to add parameters to this request, see docs/url_queries.md

	// Call function
	bookmarksList, err := client.GetBookmarks(url.Values{})
	if err != nil {
		log.Fatalf("Error fetching accounts: %v", err)
	}

	// Pretty print the accounts details
	BookmarksJSON, err := json.MarshalIndent(bookmarksList, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling bookmarks data: %v", err)
	}
	fmt.Println("Fetched Bookmarks List:", string(BookmarksJSON))
}
