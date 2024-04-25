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

	// Define the name of the ebook you want to retrieve
	ebookName := "iPhone User Guide for iOS 10.3" // Replace with the desired ebook name

	// Call GetEbooksByName function
	ebook, err := client.GetEbookByName(ebookName)
	if err != nil {
		log.Fatalf("Error fetching ebook by name: %v", err)
	}

	// Pretty print the ebook in XML
	ebookXML, err := xml.MarshalIndent(ebook, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling ebook data: %v", err)
	}
	fmt.Println("Fetched Ebook:\n", string(ebookXML))
}
