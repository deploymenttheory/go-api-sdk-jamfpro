package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug

	// Configuration for the jamfpro client
	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the ebook to be created
	newEbook := jamfpro.ResourceEbooks{
		General: jamfpro.EbookSubsetGeneral{
			Name:            "iPhone User Guide for iOS 10.3",
			Author:          "Apple Inc.",
			Version:         "1",
			Free:            true,
			URL:             "https://itunes.apple.com/us/book/iphone-user-guide-for-ios-10-3/id1134772174?mt=11&amp;uo=4",
			DeploymentType:  "Install Automatically/Prompt Users to Install",
			FileType:        "PDF",
			DeployAsManaged: true,
			Category:        jamfpro.SharedResourceCategory{ID: -1, Name: "Unknown"},
			Site:            jamfpro.SharedResourceSite{ID: -1, Name: "None"},
		},
		// Add Scope and SelfService if needed
	}

	// Call CreateEbook function
	createdEbook, err := client.CreateEbook(newEbook)
	if err != nil {
		log.Fatalf("Error creating ebook: %v", err)
	}

	// Pretty print the created ebook in XML
	ebookXML, err := xml.MarshalIndent(createdEbook, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling ebook data: %v", err)
	}
	fmt.Println("Created Ebook:\n", string(ebookXML))
}
