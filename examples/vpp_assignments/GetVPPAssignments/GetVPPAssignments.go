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

	// Call the GetVPPAssignments function
	vppAssignments, err := client.GetVPPAssignments()
	if err != nil {
		log.Fatalf("Error retrieving VPP Assignments: %v", err)
	}

	// Pretty print the VPP assignment details in XML
	vppAssignmentsXML, err := xml.MarshalIndent(vppAssignments, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error VPP assignment data: %v", err)
	}
	fmt.Println("VPP Assignment Details:\n", string(vppAssignmentsXML))
}
