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

	// Call the GetGSXConnectionInformation function
	gsxInfo, err := client.GetGSXConnectionInformation()
	if err != nil {
		log.Fatalf("Error retrieving GSX Connection Information: %v", err)
	}

	// Pretty print the gsxInfo details
	gsxInfoXML, err := xml.MarshalIndent(gsxInfo, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling gsx connection info data: %v", err)
	}
	fmt.Println("gsxInfo:", string(gsxInfoXML))
}
