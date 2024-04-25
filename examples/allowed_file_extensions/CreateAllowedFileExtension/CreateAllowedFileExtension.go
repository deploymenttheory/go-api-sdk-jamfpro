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

	// Define the allowed file extension you want to create
	newExtension := &jamfpro.ResourceAllowedFileExtension{
		Extension: "qwerty3", // Replace with the desired extension name
	}

	response, err := client.CreateAllowedFileExtension(newExtension)
	if err != nil {
		log.Fatalf("Error creating allowed file extension: %v", err)
	}

	fmt.Printf("Successfully created allowed file extension: %s with ID: %d\n", response.Extension, response.ID)

}
