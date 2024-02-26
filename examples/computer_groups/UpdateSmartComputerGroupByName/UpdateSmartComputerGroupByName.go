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

	// Define the computer group name and details for update
	groupName := "Group Name"
	// Define the computer group details for update
	groupUpdate := &jamfpro.ResourceComputerGroup{
		Name:    "Group Name",
		IsSmart: true,
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		Criteria: jamfpro.SharedContainerCriteria{
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

	// Call UpdateComputerGroupByName function
	updatedGroup, err := client.UpdateComputerGroupByName(groupName, groupUpdate)
	if err != nil {
		log.Fatalf("Error updating Computer Group by Name: %v", err)
	}

	// Pretty print the updated computer group in XML
	updatedGroupXML, err := xml.MarshalIndent(updatedGroup, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated Computer Group data: %v", err)
	}
	fmt.Println("Updated Computer Group:\n", string(updatedGroupXML))
}
