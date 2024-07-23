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

	// Define the department ID you want to retrieve
	eBooksID := "1" // Replace with the desired department ID

	// Call GetEbooksByID function
	ebook, err := client.GetEbookByID(eBooksID)
	if err != nil {
		log.Fatalf("Error fetching department by ID: %v", err)
	}

	// Pretty print the department in XML
	ebookXML, err := xml.MarshalIndent(ebook, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling ebooks data: %v", err)
	}
	fmt.Println("Fetched Ebook:\n", string(ebookXML))
}
