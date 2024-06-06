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

	// Example with both sort and filter parameters
	packages, err := client.GetPackages("id:asc", "packageName==tf-ghatest-package-suspiciouspackage")
	if err != nil {
		fmt.Println("Error fetching packages:", err)
		return
	}

	// Pretty print the created package details in XML
	packagesXML, err := xml.MarshalIndent(packages, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling created package data: %v", err)
	}
	fmt.Println("Package Details:\n", string(packagesXML))
}
