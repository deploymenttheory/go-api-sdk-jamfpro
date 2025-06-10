package main

import (
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

	// Set up the device pin request
	devicePin := jamfpro.RequestEraseDeviceComputer{
		Pin: nil, // or the six-character : Pin: &"123456" for Find My.
	}

	// Call Function to erase computer with ID "1"
	err = client.EraseComputerByID("1", devicePin)
	if err != nil {
		log.Fatalf("Error erasing computer: %v", err)
	}

	fmt.Printf("Successfully initiated erase for computer with ID: %s\n", "1")
}
