package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
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
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
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

	// Define the ebook to be created
	ebookID := 1 // Replace with the actual eBook ID
	ebookToUpdate := jamfpro.ResourceEbooks{
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

	updatedEbook, err := client.UpdateEbookByID(ebookID, ebookToUpdate)
	if err != nil {
		log.Fatalf("Error updating ebook by ID: %v", err)
	}

	ebookXML, err := xml.MarshalIndent(updatedEbook, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated ebook data: %v", err)
	}
	fmt.Println("Updated Ebook by ID:\n", string(ebookXML))
}
