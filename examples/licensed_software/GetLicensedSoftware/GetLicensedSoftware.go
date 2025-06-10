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

	// Call the GetLicensedSoftware function
	licensedSoftware, err := client.GetLicensedSoftware()
	if err != nil {
		log.Fatalf("Error retrieving licensed software: %v", err)
	}

	// Output the result
	softwareXML, err := xml.MarshalIndent(licensedSoftware, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling licensed software data: %v", err)
	}
	fmt.Println("Licensed Software List:", string(softwareXML))
}
