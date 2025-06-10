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

	newGroup := &jamfpro.ResourceMobileDeviceGroup{
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

	createdGroup, err := client.CreateMobileDeviceGroup(newGroup)
	if err != nil {
		log.Fatalf("Error creating mobile device group: %s\n", err)
	}

	createdGroupXML, err := xml.MarshalIndent(createdGroup, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created group data: %v", err)
	}
	fmt.Println("Created Mobile Device Group:\n", string(createdGroupXML))
}
