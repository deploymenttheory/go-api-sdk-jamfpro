package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"strconv"

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

	restrictedSoftwareName := "Restrict High Sierra" // Replace with actual ID

	updatedRestrictedSoftware := &jamfpro.ResourceRestrictedSoftware{
		General: jamfpro.RestrictedSoftwareSubsetGeneral{
			Name:                  "Restrict High Sierra",
			ProcessName:           "Install macOS High Sierra.app",
			MatchExactProcessName: true,
			SendNotification:      true,
			KillProcess:           true,
			DeleteExecutable:      true,
			DisplayMessage:        "High Sierra is not yet supported, check Self Service after public release.",
			Site: &jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: jamfpro.RestrictedSoftwareSubsetScope{
			AllComputers:   false,
			Computers:      []jamfpro.RestrictedSoftwareSubsetScopeEntity{},
			ComputerGroups: []jamfpro.RestrictedSoftwareSubsetScopeEntity{},
			Buildings:      []jamfpro.RestrictedSoftwareSubsetScopeEntity{},
			Departments:    []jamfpro.RestrictedSoftwareSubsetScopeEntity{},
			Exclusions: jamfpro.RestrictedSoftwareSubsetScopeExclusions{
				Computers:      []jamfpro.RestrictedSoftwareSubsetScopeEntity{},
				ComputerGroups: []jamfpro.RestrictedSoftwareSubsetScopeEntity{},
				Buildings:      []jamfpro.RestrictedSoftwareSubsetScopeEntity{},
				Departments:    []jamfpro.RestrictedSoftwareSubsetScopeEntity{},
				Users:          []jamfpro.RestrictedSoftwareSubsetScopeEntity{},
			},
		},
	}

	response, err := client.UpdateRestrictedSoftwareByName(restrictedSoftwareName, updatedRestrictedSoftware)
	if err != nil {
		fmt.Println("Error updating restricted software:", err)
		return
	}

	fmt.Printf("restricted software updated successfully, ID: %d\n", response.ID)

	// Fetch the full details of the updated restricted software
	updatedPrinterDetails, err := client.GetRestrictedSoftwareByID(strconv.Itoa(response.ID))
	if err != nil {
		fmt.Println("Error fetching updated restricted software details:", err)
		return
	}

	// Marshal the updated restricted software details to XML for display
	softwareXML, err := xml.MarshalIndent(updatedPrinterDetails, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling updated restricted software to XML: %v", err)
	}

	fmt.Printf("Updated restricted software Details:\n%s\n", softwareXML)
}
