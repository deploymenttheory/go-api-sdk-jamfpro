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

	// Call GetVPPAccounts
	vppAccounts, err := client.GetVPPAccounts()
	if err != nil {
		log.Fatalf("Error fetching VPP accounts: %v", err)
	}

	// Pretty print the details in XML
	vppAccountXML, err := xml.MarshalIndent(vppAccounts, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling VPP account data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(vppAccountXML))
}
