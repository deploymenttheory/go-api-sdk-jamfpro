package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client"
	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "/Users/dafyddwatkins/localtesting/clientauth.json"

	// Load the client OAuth credentials from the configuration file
	authConfig, err := jamfpro.LoadAuthConfig(configFilePath)
	if err != nil {
		log.Fatalf("Failed to load client OAuth configuration: %v", err)
	}

	// Instantiate the default logger and set the desired log level
	logLevel := http_client.LogLevelWarning // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

	// Configuration for the jamfpro
	config := http_client.Config{
		InstanceName: authConfig.InstanceName,
		Auth: http_client.AuthConfig{
			ClientID:     authConfig.ClientID,
			ClientSecret: authConfig.ClientSecret,
		},
		LogLevel: logLevel,
	}

	// Create a new jamfpro client instance
	client, err := jamfpro.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}

	// Define new VPP assignment details
	newVPPAssignment := &jamfpro.ResourceVPPAssignment{
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

	// Call the CreateVPPAssignment function
	err = client.CreateVPPAssignment(newVPPAssignment)
	if err != nil {
		log.Fatalf("Error creating VPP Assignment: %v", err)
	}

	fmt.Println("VPP Assignment created successfully.")
}
