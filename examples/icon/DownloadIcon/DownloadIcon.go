package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file inside the main function
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		DebugMode:    true,
		Logger:       jamfpro.NewDefaultLogger(),
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instanceclient,
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Set the icon ID to download
	iconID := 1 // Replace with your actual icon ID

	// Set the path where the icon should be saved
	savePath := "/Users/dafyddwatkins/Downloads/saved-icon.png" // Replace with the actual path where you want to save the icon

	// Set the desired resolution and scale
	res := "original" // or "300" or "512"
	scale := "0"      // or other scale as a string

	// Call DownloadIcon with the new parameters
	err = client.DownloadIcon(iconID, savePath, res, scale)
	if err != nil {
		fmt.Printf("Error downloading icon: %s\n", err)
	} else {
		fmt.Println("Icon downloaded successfully!")
	}
}
