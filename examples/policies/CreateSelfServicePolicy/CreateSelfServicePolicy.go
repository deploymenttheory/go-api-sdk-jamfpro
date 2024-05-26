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
		// General
		General: jamfpro.PolicySubsetGeneral{
			Name:                       "jamfpro-sdk-example-selfservice-policy-config",
			Enabled:                    false,
			TriggerCheckin:             false,
			TriggerEnrollmentComplete:  false,
			TriggerLogin:               false,
			TriggerLogout:              false,
			TriggerNetworkStateChanged: false,
			TriggerStartup:             false,
			TriggerOther:               "EVENT",
			Frequency:                  "Once per computer",
			RetryEvent:                 "none",
			RetryAttempts:              -1,
			NotifyOnEachFailedRetry:    false,
			LocationUserOnly:           false,
			TargetDrive:                "/",
			Offline:                    false,
		},
		// Self Service policy settings
		SelfService: &jamfpro.PolicySubsetSelfService{
			UseForSelfService:           true,
			SelfServiceDisplayName:      "some text here",
			InstallButtonText:           "Install",
			ReinstallButtonText:         "Reinstall",
			SelfServiceDescription:      "some text here",
			ForceUsersToViewDescription: true,
			SelfServiceIcon: &jamfpro.SharedResourceSelfServiceIcon{
				ID:       3,
				Filename: "mac-icon.png",
				URI:      "https://euw2.ics.services.jamfcloud.com/icon/hash_f6d371a96ce011c4c297bcc09641bdce76e90c53e79bb212f8ca3024cbb53034",
			},
			FeatureOnMainPage: true,
			SelfServiceCategories: &[]jamfpro.PolicySubsetSelfServiceCategory{
				{
					ID:        6,
					Name:      "Productivity",
					DisplayIn: true,
					FeatureIn: true,
				},
			},
			Notification:        true,
			NotificationType:    "Self Service and Notification Center",
			NotificationSubject: "thing",
			NotificationMessage: "thing",
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
