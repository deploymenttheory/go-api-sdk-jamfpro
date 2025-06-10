package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Fetch a mobile device application by ID
	appByID, err := client.GetMobileDeviceApplicationByID("5") // replace 123 with an actual ID
	if err != nil {
		fmt.Println("Error fetching by ID:", err)
	} else {
		fmt.Println("Fetched by ID:", appByID)
	}

	// Pretty print the details in XML
	applicationsXML, err := xml.MarshalIndent(appByID, "", "    ") // Indent with 4 spaces and use '='
	if err != nil {
		log.Fatalf("Error marshaling policy details data: %v", err)
	}
	fmt.Println("Created Policy Details:\n", string(applicationsXML))
}
