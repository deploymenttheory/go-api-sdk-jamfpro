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

	ebookID := "1" // Replace with the actual eBook ID

	err = client.DeleteEbookByID(ebookID)
	if err != nil {
		log.Fatalf("Error deleting ebook by ID: %v", err)
	}

	fmt.Println("Ebook successfully deleted by ID")
}
