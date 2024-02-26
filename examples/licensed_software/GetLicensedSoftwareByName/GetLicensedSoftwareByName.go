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

	// Name of the licensed software to fetch
	softwareName := "Adobe Creative Suite" // Replace with actual Name

	licensedSoftware, err := client.GetLicensedSoftwareByName(softwareName)
	if err != nil {
		log.Fatalf("Failed to get licensed software by ID: %v", err)
	}

	// Pretty print the created software details
	createdSoftwareXML, err := xml.MarshalIndent(licensedSoftware, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created software data: %v", err)
	}
	fmt.Println("Created Licensed Software:", string(createdSoftwareXML))
}
