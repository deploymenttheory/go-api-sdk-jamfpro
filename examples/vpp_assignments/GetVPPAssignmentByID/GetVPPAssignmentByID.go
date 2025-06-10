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

	// Specify the ID of the VPP assignment to retrieve
	vppAssignmentID := "1" // Replace with the actual ID

	// Call the GetVPPAssignmentByID function
	vppAssignment, err := client.GetVPPAssignmentByID(vppAssignmentID)
	if err != nil {
		log.Fatalf("Error retrieving VPP Assignment by ID: %v", err)
	}

	// Pretty print the VPP assignment details in XML
	vppAssignmentsXML, err := xml.MarshalIndent(vppAssignment, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error VPP assignment data: %v", err)
	}
	fmt.Println("VPP Assignment Details:\n", string(vppAssignmentsXML))
}
