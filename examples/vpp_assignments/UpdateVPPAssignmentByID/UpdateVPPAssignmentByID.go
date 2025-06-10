package main

import (
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

	// Specify the ID of the VPP assignment to update
	vppAssignmentID := "1" // Replace with the actual ID

	// Define updated VPP assignment details
	updatedVPPAssignment := &jamfpro.ResourceVPPAssignment{
		General: jamfpro.VPPAssignmentSubsetGeneral{
			Name:              "Sample Assignment",
			VPPAdminAccountID: 1,
		},
		IOSApps: []jamfpro.VPPSubsetVPPApp{
			{AdamID: 767319014, Name: "Angry Birds Epic RPG"},
			{AdamID: 923394341, Name: "Alien Blue for iPad - reddit official client"},
		},
		MacApps: []jamfpro.VPPSubsetVPPApp{}, // Empty as per the example
		EBooks: []jamfpro.VPPSubsetVPPApp{
			{AdamID: 1058120411, Name: "Transforming Healthcare"},
		},
		Scope: jamfpro.VPPAssignmentSubsetScope{
			AllJSSUsers:   false,
			JSSUsers:      []jamfpro.VPPSubsetVPPUser{},      // Empty as per the example
			JSSUserGroups: []jamfpro.VPPSubsetVPPUserGroup{}, // Empty as per the example
			Limitations: jamfpro.VPPAssignmentSubsetScopeLimitations{
				UserGroups: []jamfpro.VPPSubsetVPPUserGroup{}, // Empty as per the example
			},
			Exclusions: jamfpro.VPPAssignmentSubsetScopeExclusions{
				JSSUsers:      []jamfpro.VPPSubsetVPPUser{},      // Empty as per the example
				UserGroups:    []jamfpro.VPPSubsetVPPUserGroup{}, // Empty as per the example
				JSSUserGroups: []jamfpro.VPPSubsetVPPUserGroup{}, // Empty as per the example
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
