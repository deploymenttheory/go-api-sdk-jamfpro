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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define the advanced computer search details
	newSearch := &jamfpro.ResourceAdvancedComputerSearch{
		Name:   "jamf pro SDK - Advanced Search Name",
		ViewAs: "Standard Web Page",
		Criteria: jamfpro.SharedContainerCriteria{
			Size: 4,
			Criterion: []jamfpro.SharedSubsetCriteria{
				{
					Name:         "Building",
					Priority:     0,
					AndOr:        "and",
					SearchType:   "is",
					Value:        "square",
					OpeningParen: true,
					ClosingParen: false,
				},
				{
					Name:         "Model",
					Priority:     1,
					AndOr:        "and",
					SearchType:   "is",
					Value:        "macbook air",
					OpeningParen: false,
					ClosingParen: true,
				},
				{
					Name:         "Computer Name",
					Priority:     2,
					AndOr:        "or",
					SearchType:   "matches regex",
					Value:        "thing",
					OpeningParen: true,
					ClosingParen: false,
				},
				{
					Name:         "Licensed Software",
					Priority:     3,
					AndOr:        "and",
					SearchType:   "has",
					Value:        "office",
					OpeningParen: false,
					ClosingParen: true,
				},
			},
		},
		DisplayFields: []jamfpro.SharedAdvancedSearchContainerDisplayField{
			{
				DisplayField: []jamfpro.SharedAdvancedSearchSubsetDisplayField{
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
			},
		},
		Site: jamfpro.SharedResourceSite{
			ID:   -1,
			Name: "None",
		},
	}

	// Marshal the newSearch object into XML for logging
	newSearchJSON, err := xml.MarshalIndent(newSearch, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling new search to XML:", err)
		return
	}
	fmt.Printf("New Advanced Computer Search Request:\n%s\n", string(newSearchJSON))

	// Create the advanced computer search
	createdSearch, err := client.CreateAdvancedComputerSearch(newSearch)
	if err != nil {
		fmt.Println("Error creating advanced computer search:", err)
		return
	}

	// Print the created advanced computer search details
	createdSearchXML, err := xml.MarshalIndent(createdSearch, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling created search to XML:", err)
		return
	}
	fmt.Printf("Created Advanced Computer Search:\n%s\n", string(createdSearchXML))
}
