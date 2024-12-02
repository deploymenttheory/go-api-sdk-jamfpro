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

	// Define a new policy with all required fields
	newPolicy := &jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			Name:                       "jamfpro-sdk-example-dockitems-policy-config",
			Enabled:                    jamfpro.FalsePtr(),
			TriggerCheckin:             jamfpro.FalsePtr(),
			TriggerEnrollmentComplete:  jamfpro.FalsePtr(),
			TriggerLogin:               jamfpro.FalsePtr(),
			TriggerLogout:              jamfpro.FalsePtr(),
			TriggerNetworkStateChanged: jamfpro.FalsePtr(),
			TriggerStartup:             jamfpro.FalsePtr(),
			TriggerOther:               "EVENT",
			Frequency:                  "Once per computer",
			RetryEvent:                 "none",
			RetryAttempts:              -1,
			NotifyOnEachFailedRetry:    jamfpro.FalsePtr(),
			LocationUserOnly:           jamfpro.FalsePtr(),
			TargetDrive:                "/",
			Offline:                    jamfpro.FalsePtr(),
		},
		DockItems: []jamfpro.PolicySubsetDockItem{
			{
				ID:     1832,
				Name:   "Safari",
				Action: "Add To End", // Add To Beginning / Add To End / Remove
			},
		},
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           jamfpro.TruePtr(),
			SelfServiceDisplayName:      "",
			InstallButtonText:           "Install",
			ReinstallButtonText:         "Reinstall",
			SelfServiceDescription:      "",
			ForceUsersToViewDescription: jamfpro.FalsePtr(),
			FeatureOnMainPage:           jamfpro.FalsePtr(),
		},
	}

	policyXML, err := xml.MarshalIndent(newPolicy, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling policy data: %v", err)
	}
	fmt.Println("Policy Details to be Sent:\n", string(policyXML))

	// Call CreatePolicy function
	createdPolicy, err := client.CreatePolicy(newPolicy)
	if err != nil {
		log.Fatalf("Error creating policy: %v", err)
	}

	// Pretty print the created policy details in XML
	policyXML, err = xml.MarshalIndent(createdPolicy, "", "    ") // Indent with 4 spaces and use '='
	if err != nil {
		log.Fatalf("Error marshaling policy details data: %v", err)
	}
	fmt.Println("Created Policy Details:\n", string(policyXML))
}
