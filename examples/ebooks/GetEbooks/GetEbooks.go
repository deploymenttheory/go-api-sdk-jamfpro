package main

import (
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

	// Call the GetEbooks function
	ebooks, err := client.GetEbooks()
	if err != nil {
		log.Fatalf("Error fetching eBooks: %v", err)
	}

	// Output the fetched eBooks
	fmt.Printf("Fetched eBooks: %+v\n", ebooks)
}
