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
	// Example ID to fetch
	printerName := "printer-name"

	printer, err := client.GetPrinterByName(printerName)
	if err != nil {
		fmt.Println("Error fetching printer by ID:", err)
		return
	}

	// Pretty print the created script details in XML
	createdPrintersXML, err := xml.MarshalIndent(printer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created script data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(createdPrintersXML))
}
