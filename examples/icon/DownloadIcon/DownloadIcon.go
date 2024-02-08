package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-http-client/httpclient"
	"github.com/deploymenttheory/go-api-http-client/logger"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := logger.LogLevelWarn // LogLevelNone / LogLevelDebug / LogLevelInfo / LogLevelError

	// Configuration for the jamfpro
	config := httpclient.Config{
		InstanceName: authConfig.InstanceName,
		Auth: httpclient.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Set the icon ID to download
	iconID := 2 // Replace with your actual icon ID

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
