package main

import (
	"encoding/json"
	"fmt"
	"log"

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

	// update bookmark
	bookmark := &jamfpro.ResourceBookmark{
		SiteID:           "-1",
		Priority:         5,
		DisplayInBrowser: jamfpro.BoolPtr(false),
		Name:             "Example Bookmark",
		Description:      "This is an example bookmark",
		ScopeDescription: "", //"No scope defined"
		URL:              "https://example.com",
		IconID:           "38",
	}

	id := "1"

	// Call the Bookmark function
	response, err := client.UpdateBookmarkByID(id, bookmark)
	if err != nil {
		log.Fatalf("Failed to create bookmark: %v", err)
	}

	// Pretty print the response
	prettyResponse, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("Failed to marshal response: %v", err)
	}

	fmt.Printf("Successfully created bookmark. Response:\n%s\n", string(prettyResponse))
}
