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

	// Example device ID
	deviceID := "1" // Replace with an actual device name

	// Get mobile device by ID
	deviceByID, err := client.GetMobileDeviceByID(deviceID)
	if err != nil {
		log.Fatalf("Error fetching mobile device by name: %v", err)
	}

	// Pretty print the network segments in XML
	mobileDeviceXML, err := xml.MarshalIndent(deviceByID, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling network segments data: %v", err)
	}
	fmt.Println("Network Segments:\n", string(mobileDeviceXML))
}
