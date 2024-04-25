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

	// Sample data for creating a new computer group (replace with actual data as needed)
	newSmartGroup := &jamfpro.ResourceComputerGroup{
		Name:    "Operating System Version like 15",
		IsSmart: true,
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: jamfpro.SharedContainerCriteria{
			Size: 1, // Assuming there is only one criterion
			Criterion: []jamfpro.SharedSubsetCriteria{
				{
					Name:       "Operating System Version",
					Priority:   0,
					AndOr:      "and",
					SearchType: "like",
					Value:      "macOS 14",
				},
			},
		},
		// Include other fields if necessary
	}

	// Call CreateComputerGroup function
	createdGroup, err := client.CreateComputerGroup(newSmartGroup)
	if err != nil {
		log.Fatalf("Error creating Computer Group: %v", err)
	}

	// Pretty print the created group in XML
	groupXML, err := xml.MarshalIndent(createdGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling Computer Group data: %v", err)
	}
	fmt.Println("Created Computer Group:\n", string(groupXML))
}
