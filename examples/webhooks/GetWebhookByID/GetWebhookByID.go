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

	// Example usage of GetWebhookByID
	webhookByID, err := client.GetWebhookByID(1) // Replace with the desired webhook ID
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Pretty print the user groups details in XML
	webhooksXML, err := xml.MarshalIndent(webhookByID, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling webhook data: %v", err)
	}
	fmt.Println("User Groups Details:\n", string(webhooksXML))
}
