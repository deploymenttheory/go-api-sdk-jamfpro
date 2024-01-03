package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug

	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
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
