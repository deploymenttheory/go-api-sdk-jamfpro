package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
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
	logLevel := http_client.LogLevelDebug

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

	// Specify the ID of the VPP assignment to update
	vppAssignmentID := 1 // Replace with the actual ID

	// Define updated VPP assignment details
	updatedVPPAssignment := &jamfpro.ResponseVPPAssignment{
		General: jamfpro.VPPAssignmentGeneral{
			Name:              "Sample Assignment",
			VPPAdminAccountID: 1,
		},
		IOSApps: []jamfpro.VPPApp{
			{AdamID: 767319014, Name: "Angry Birds Epic RPG"},
			{AdamID: 923394341, Name: "Alien Blue for iPad - reddit official client"},
		},
		MacApps: []jamfpro.VPPApp{}, // Empty as per the example
		EBooks: []jamfpro.VPPApp{
			{AdamID: 1058120411, Name: "Transforming Healthcare"},
		},
		Scope: jamfpro.VPPAssignmentScope{
			AllJSSUsers:   false,
			JSSUsers:      []jamfpro.VPPUser{},      // Empty as per the example
			JSSUserGroups: []jamfpro.VPPUserGroup{}, // Empty as per the example
			Limitations: jamfpro.VPPLimitations{
				UserGroups: []jamfpro.VPPUserGroup{}, // Empty as per the example
			},
			Exclusions: jamfpro.VPPExclusions{
				JSSUsers:      []jamfpro.VPPUser{},      // Empty as per the example
				UserGroups:    []jamfpro.VPPUserGroup{}, // Empty as per the example
				JSSUserGroups: []jamfpro.VPPUserGroup{}, // Empty as per the example
			},
		},
	}

	// Call the UpdateVPPAssignmentByID function
	err = client.UpdateVPPAssignmentByID(vppAssignmentID, updatedVPPAssignment)
	if err != nil {
		log.Fatalf("Error updating VPP Assignment by ID: %v", err)
	}

	fmt.Println("VPP Assignment updated successfully.")
}
