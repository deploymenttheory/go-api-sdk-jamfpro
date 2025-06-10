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

	computerName := "Steve Job's iMac"

	// Call the GetComputerByName method
	computer, err := client.GetComputerByName(computerName)
	if err != nil {
		log.Fatalf("Error fetching computer by name: %v", err)
	}

	// Pretty print the created department in JSON
	computerJSON, err := xml.MarshalIndent(computer, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created computer data: %v", err)
	}
	fmt.Println("Created Computer:\n", string(computerJSON))
}
