package main

import (
	"encoding/xml"
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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	ebookName := "iPhone User Guide for iOS 10.3" // Replace with the actual eBook name
	ebookToUpdate := jamfpro.ResourceEbooks{
		General: jamfpro.EbookSubsetGeneral{
			Name:            "iPhone User Guide for iOS 16",
			Author:          "Apple Inc.",
			Version:         "1",
			Free:            true,
			URL:             "https://books.apple.com/gb/book/iphone-user-guide/id6443146864",
			DeploymentType:  "Install Automatically/Prompt Users to Install",
			FileType:        "PDF",
			DeployAsManaged: true,
			Category:        jamfpro.SharedResourceCategory{ID: -1, Name: "Unknown"},
			Site:            jamfpro.SharedResourceSite{ID: -1, Name: "None"},
		},
		// Add Scope and SelfService if needed
	}

	updatedEbook, err := client.UpdateEbookByName(ebookName, ebookToUpdate)
	if err != nil {
		log.Fatalf("Error updating ebook by Name: %v", err)
	}

	ebookXML, err := xml.MarshalIndent(updatedEbook, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated ebook data: %v", err)
	}
	fmt.Println("Updated Ebook by Name:\n", string(ebookXML))
}
