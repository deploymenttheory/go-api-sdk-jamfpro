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

	// Define the updated advanced user search details
	updatedAdvancedUserSearch := &jamfpro.ResourceAdvancedUserSearch{
		Name: "Advanced User Search Name by jamf pro sdk",
		Criteria: jamfpro.SharedContainerCriteria{
			Criterion: []jamfpro.SharedSubsetCriteria{
				{
					Name:         "Email Address",
					Priority:     0,
					AndOr:        "and",
					SearchType:   "like",
					Value:        "company.com",
					OpeningParen: false,
					ClosingParen: false,
				},
			},
		},
		DisplayFields: []jamfpro.DisplayField{
			{
				Name: "Activation Lock Manageable",
			},
			{
				Name: "Apple Silicon",
			},
			{
				Name: "Architecture Type",
			},
			{
				Name: "Available RAM Slots",
			},
		},

		Site: &jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	searchID := "17" // Replace 123 with the actual ID

	// Update by ID
	updatedByID, err := client.UpdateAdvancedUserSearchByID(searchID, updatedAdvancedUserSearch)
	if err != nil {
		log.Fatalf("Error updating advanced user search by ID: %v", err)
	}
	// Print updated search by ID
	updatedByIDXML, err := xml.MarshalIndent(updatedByID, "", "    ") // Indent with 4 spaces
	if err != nil {
		log.Fatalf("Error marshaling updated advanced user search by ID: %v", err)
	}
	fmt.Println("Updated Advanced User Search by ID:\n", string(updatedByIDXML))
}
