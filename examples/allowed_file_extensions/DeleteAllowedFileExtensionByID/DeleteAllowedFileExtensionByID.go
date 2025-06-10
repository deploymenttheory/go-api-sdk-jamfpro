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

	// Define the ID of the file extension you want to delete
	fileExtensionID := "147" // Replace with the ID of the extension you want to delete

	// Call DeleteAllowedFileExtensionByID function
	err = client.DeleteAllowedFileExtensionByID(fileExtensionID)
	if err != nil {
		log.Fatalf("Error deleting allowed file extension by ID: %v", err)
	}

	// If the deletion was successful
	fmt.Println("Allowed file extension deleted successfully!")
}
