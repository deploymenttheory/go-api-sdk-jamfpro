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

	// Name of the distribution point to fetch
	distributionPointName := "New York Share" // Replace with the actual name

	// Call GetDistributionPointByName function
	distributionPoint, err := client.GetDistributionPointByName(distributionPointName)
	if err != nil {
		log.Fatalf("Error fetching distribution point: %v", err)
	}

	// Pretty print the distribution point in XML
	distributionPointXML, err := xml.MarshalIndent(distributionPoint, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling distribution point data: %v", err)
	}
	fmt.Println("Fetched Distribution Point:\n", string(distributionPointXML))
}
