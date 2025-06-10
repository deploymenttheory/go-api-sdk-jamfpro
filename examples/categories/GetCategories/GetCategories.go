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

	// Define the sort and filter query parameters
	// For more information on how to add parameters to this request, see docs/url_queries.md

	// Call the GetCategories function
	categories, err := client.GetCategories(url.Values{}) // Will return all results by default
	if err != nil {
		fmt.Printf("Error fetching categories: %v\n", err)
		return
	}

	// Print the fetched categories
	fmt.Println("Fetched Categories:")
	for _, category := range categories.Results {
		fmt.Printf("ID: %s, Name: %s, Priority: %d\n", category.Id, category.Name, category.Priority)
	}
}
