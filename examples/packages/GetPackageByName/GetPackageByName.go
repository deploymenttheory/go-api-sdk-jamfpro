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

	// Example ID to fetch
	packageName := "OpenJDK21U-jdk_aarch64_mac_hotspot_21.0.2_13.pkg"

	response, err := client.GetPackageByName(packageName)
	if err != nil {
		fmt.Println("Error fetching package by Name:", err)
		return
	}

	// Pretty print the created script details in XML
	packageXML, err := xml.MarshalIndent(response, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created script data: %v", err)
	}
	fmt.Println("Created Script Details:\n", string(packageXML))
}
