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

	// Define the fuction parameters
	computerMACAddress := "00:1A:2B:3C:4D:5E"

	// Call the GetComputerHistoryByComputerID function
	computerHistory, err := client.GetComputerHistoryByComputerMACAddress(computerMACAddress)
	if err != nil {
		log.Fatalf("Error fetching computer inventory by computer MAC Address: %v", err)
	}

	// Pretty print the response
	prettyXML, err := xml.MarshalIndent(computerHistory, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty XML: %v", err)
	}
	fmt.Printf("%s\n", prettyXML)
}
