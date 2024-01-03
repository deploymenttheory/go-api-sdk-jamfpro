package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug

	config := jamfpro.Config{
		InstanceName: authConfig.InstanceName,
		LogLevel:     logLevel,
		Logger:       logger,
		ClientID:     authConfig.ClientID,
		ClientSecret: authConfig.ClientSecret,
	}

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
