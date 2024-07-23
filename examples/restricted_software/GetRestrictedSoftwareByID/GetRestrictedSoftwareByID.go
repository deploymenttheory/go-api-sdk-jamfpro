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

	restrictedSoftwareID := "1" // Replace with actual ID

	restrictedSoftware, err := client.GetRestrictedSoftwareByID(restrictedSoftwareID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Pretty print the restricted software details in XML
	restrictedSoftwareXML, err := xml.MarshalIndent(restrictedSoftware, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling restricted software data: %v", err)
	}
	fmt.Println("Restricted Software Details:\n", string(restrictedSoftwareXML))
}
