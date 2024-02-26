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

	// Fetching a single mobile device group by name
	groupName := "GroupName" // Replace with the actual group name

	mobileGroup, err := client.GetMobileDeviceGroupByName(groupName)
	if err != nil {
		log.Fatalf("Error fetching mobile device group by name: %s\n", err)
	}

	// Pretty print the extension attribute details in XML
	applicationsXML, err := xml.MarshalIndent(mobileGroup, "", "    ") // Indent with 4 spaces and use '='
	if err != nil {
		log.Fatalf("Error marshaling policy details data: %v", err)
	}
	fmt.Println("Created Policy Details:\n", string(applicationsXML))
}
