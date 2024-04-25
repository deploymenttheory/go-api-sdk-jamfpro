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

	// Define the name of the file extension you want to delete
	fileExtensionName := "qwerty3" // Replace with the name of the extension you want to delete

	// Call DeleteAllowedFileExtensionByNameByID function
	err = client.DeleteAllowedFileExtensionByName(fileExtensionName)
	if err != nil {
		log.Fatalf("Error deleting allowed file extension by name: %v", err)
	}

	// If the deletion was successful
	fmt.Println("Allowed file extension deleted successfully by name!")
}
