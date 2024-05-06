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

	// Define the computer name of the computer history you want to retrieve
	computerName := "APIGWFYFGH"

	// Call the GetComputerHistoryByComputerName function
	computerHistory, err := client.GetComputerHistoryByComputerName(computerName)
	if err != nil {
		log.Fatalf("Error fetching computer history by Name: %v", err)
	}

	// Pretty print the response
	prettyXML, err := xml.MarshalIndent(computerHistory, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty XML: %v", err)
	}
	fmt.Printf("%s\n", prettyXML)
}
