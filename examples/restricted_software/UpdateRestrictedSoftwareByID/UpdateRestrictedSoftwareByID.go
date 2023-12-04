package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/GitHub/deploymenttheory/go-api-sdk-jamfpro/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadClientAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logger := http_client.NewDefaultLogger()
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := jamfpro.Config{
		InstanceName:       authConfig.InstanceName,
		OverrideBaseDomain: authConfig.OverrideBaseDomain,
		LogLevel:           logLevel,
		Logger:             logger,
		ClientID:           authConfig.ClientID,
		ClientSecret:       authConfig.ClientSecret,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	restrictedSoftwareID := 1 // Replace with actual ID

	updatedRestrictedSoftware := &jamfpro.ResponseRestrictedSoftware{
		General: jamfpro.RestrictedSoftwareDataSubsetGeneral{
			Name:                  "Restrict High Sierra 2",
			ProcessName:           "Install macOS High Sierra.app",
			MatchExactProcessName: true,
			SendNotification:      true,
			KillProcess:           true,
			DeleteExecutable:      true,
			DisplayMessage:        "High Sierra is not yet supported, check Self Service after public release.",
			Site: jamfpro.RestrictedSoftwareDataSubsetSite{
				ID:   -1,
				Name: "None",
			},
		},
		Scope: jamfpro.RestrictedSoftwareDataSubsetScope{
			AllComputers:   false,
			Computers:      []jamfpro.RestrictedSoftwareDataSubsetScopeComputer{},
			ComputerGroups: []jamfpro.RestrictedSoftwareDataSubsetScopeComputerGroup{},
			Buildings:      []jamfpro.RestrictedSoftwareDataSubsetScopeBuilding{},
			Departments:    []jamfpro.RestrictedSoftwareDataSubsetScopeDepartment{},
			Exclusions: jamfpro.RestrictedSoftwareExclusions{
				Computers:      []jamfpro.RestrictedSoftwareDataSubsetScopeComputer{},
				ComputerGroups: []jamfpro.RestrictedSoftwareDataSubsetScopeComputerGroup{},
				Buildings:      []jamfpro.RestrictedSoftwareDataSubsetScopeBuilding{},
				Departments:    []jamfpro.RestrictedSoftwareDataSubsetScopeDepartment{},
				Users:          []jamfpro.RestrictedSoftwareDataSubsetScopeUser{},
			},
		},
	}

	err = client.UpdateRestrictedSoftwareByID(restrictedSoftwareID, updatedRestrictedSoftware)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Restricted software updated successfully.")
}
