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

	// Define the path to the icon file
	iconPath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/examples/icon/UploadIcon/cat.png"

	// Upload the icon
	response, err := client.UploadIcon(iconPath)
	if err != nil {
		fmt.Println("Error uploading icon:", err)
		return
	}

	// Print the response
	fmt.Printf("Uploaded icon successfully. ID: %s, Href: %s\n", response.ID, response.Href)
}
