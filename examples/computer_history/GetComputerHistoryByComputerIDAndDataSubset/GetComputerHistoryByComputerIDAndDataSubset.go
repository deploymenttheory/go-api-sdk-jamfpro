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

	// Define the fuction parameters
	computerID := "1"
	DataSubset := "General" // General / ComputerUsageLogs / Audits / PolicyLogs / CasperRemoteLogs / ScreenSharingLogs / CasperImagingLogs / Commands / UserLocation / MacAppStoreApplications

	// Call the GetComputerHistoryByComputerIDAndDataSubset function
	computerHistory, err := client.GetComputerHistoryByComputerIDAndDataSubset(computerID, DataSubset)
	if err != nil {
		log.Fatalf("Error fetching computer inventory by ID: %v", err)
	}

	// Pretty print the response
	prettyXML, err := xml.MarshalIndent(computerHistory, "", "    ")
	if err != nil {
		log.Fatalf("Failed to generate pretty XML: %v", err)
	}
	fmt.Printf("%s\n", prettyXML)
}
