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

	newRestrictedSoftware := &jamfpro.ResourceRestrictedSoftware{
		General: jamfpro.RestrictedSoftwareSubsetGeneral{
			Name:                  "Restrict High Sierra",
			ProcessName:           "Install macOS High Sierra.app",
			MatchExactProcessName: true,
			SendNotification:      true,
			KillProcess:           true,
			DeleteExecutable:      true,
			DisplayMessage:        "High Sierra is not yet supported, check Self Service after public release.",
			Site: jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: jamfpro.RestrictedSoftwareSubsetScope{
			AllComputers:   false,
			Computers:      []jamfpro.RestrictedSoftwareSubsetScopeComputer{},
			ComputerGroups: []jamfpro.RestrictedSoftwareSubsetScopeComputerGroup{},
			Buildings:      []jamfpro.RestrictedSoftwareSubsetScopeBuilding{},
			Departments:    []jamfpro.RestrictedSoftwareSubsetScopeDepartment{},
			Exclusions: jamfpro.RestrictedSoftwareSubsetScopeExclusions{
				Computers:      []jamfpro.RestrictedSoftwareSubsetScopeComputer{},
				ComputerGroups: []jamfpro.RestrictedSoftwareSubsetScopeComputerGroup{},
				Buildings:      []jamfpro.RestrictedSoftwareSubsetScopeBuilding{},
				Departments:    []jamfpro.RestrictedSoftwareSubsetScopeDepartment{},
				Users:          []jamfpro.RestrictedSoftwareSubsetScopeUser{},
			},
		},
	}

	createdRestrictedSoftware, err := client.CreateRestrictedSoftware(newRestrictedSoftware)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	createdRestrictedSoftwareXML, err := xml.MarshalIndent(createdRestrictedSoftware, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created restricted software data: %v", err)
	}
	fmt.Println("Created Restricted Software Details:\n", string(createdRestrictedSoftwareXML))
}
