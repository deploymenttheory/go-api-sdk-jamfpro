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

	// ID of the distribution point to fetch
	distributionPointID := 5 // Replace with actual ID

	// Call GetDistributionPointByID function
	distributionPoint, err := client.GetDistributionPointByID(distributionPointID)
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
