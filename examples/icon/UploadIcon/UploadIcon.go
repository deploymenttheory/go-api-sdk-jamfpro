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

	// Path to the icon file you want to upload
	filePath := "/Users/dafyddwatkins/Downloads/mac-icon.png"

	// Call the UploadIcon function
	uploadResponse, err := client.UploadIcon(filePath)
	if err != nil {
		fmt.Printf("Error uploading icon: %s\n", err)
		return
	}

	// Print out the response from the server
	fmt.Printf("Icon uploaded successfully!\nURL: %s\nID: %d\n", uploadResponse.URL, uploadResponse.ID)
}
