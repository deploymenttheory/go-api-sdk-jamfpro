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

	// Create a new PatchSoftwareTitleConfiguration
	newConfig := &jamfpro.ResourcePatchSoftwareTitleConfiguration{
		DisplayName:        "CCleaner",
		CategoryID:         "-1",
		SiteID:             "-1",
		UiNotifications:    true,
		EmailNotifications: true,
		SoftwareTitleID:    "10",
		ExtensionAttributes: []jamfpro.PatchSoftwareTitleConfigurationSubsetExtensionAttribute{
			{
				Accepted: true,
				EaID:     "CCleaner-ea",
			},
		},
	}

	// Call the CreatePatchSoftwareTitleConfiguration method
	response, err := client.CreatePatchSoftwareTitleConfiguration(*newConfig)
	if err != nil {
		fmt.Println("Error creating Patch Software Title Configuration:", err)
		return
	}

	// Print the response
	fmt.Printf("Created Patch Software Title Configuration with ID: %s, Href: %s\n", response.ID, response.Href)
}
