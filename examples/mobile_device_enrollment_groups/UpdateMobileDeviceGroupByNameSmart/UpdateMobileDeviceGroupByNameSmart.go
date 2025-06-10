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

	// Define the Name of the group you want to update
	groupName := "Smart-Group-Name" // Replace with the actual group Name

	// Define the updated group data
	updatedSmartGroup := &jamfpro.ResourceMobileDeviceGroup{
		Name:    "Sample Smart Group",
		IsSmart: true,
		Criteria: jamfpro.SharedContainerCriteria{
			Size: 3, // The number of criteria
			Criterion: []jamfpro.SharedSubsetCriteria{
				{
					Name:         "Last Inventory Update",
					Priority:     0,
					AndOr:        "AND",
					SearchType:   "more than x days ago",
					Value:        "7",
					OpeningParen: true,
				},
				{
					Name:         "Department",
					Priority:     1,
					AndOr:        "and",
					SearchType:   "is",
					Value:        "marketing",
					ClosingParen: true,
				},
				{
					Name:         "Building",
					Priority:     2,
					AndOr:        "or",
					SearchType:   "is",
					Value:        "london wall",
					OpeningParen: true,
					ClosingParen: true,
				},
			},
		},
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
		// other fields if necessary
	}

	// Call the UpdateMobileDeviceGroupByName function
	updatedGroup, err := client.UpdateMobileDeviceGroupByName(groupName, updatedSmartGroup)
	if err != nil {
		log.Fatalf("Error updating mobile device group: %s\n", err)
	}

	// Marshal and print the updated group
	updatedGroupXML, err := xml.MarshalIndent(updatedGroup, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated group data: %v", err)
	}
	fmt.Println("Updated Mobile Device Group:\n", string(updatedGroupXML))
}
