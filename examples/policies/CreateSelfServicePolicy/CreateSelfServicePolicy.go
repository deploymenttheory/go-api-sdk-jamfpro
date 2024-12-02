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

	// Define a new policy with the fetched values
	newPolicy := &jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			Name:                       "example-self-service-policy-config",
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
			Category: &jamfpro.SharedResourceCategory{
				ID:   -1,
				Name: "No category assigned",
			},
			NetworkLimitations: &jamfpro.PolicySubsetGeneralNetworkLimitations{
				MinimumNetworkConnection: "No Minimum",
				AnyIPAddress:             false,
				NetworkSegments:          "",
			},
			NetworkRequirements: "Any",
			Site: &jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "NONE",
			},
		},
		Scope: jamfpro.PolicySubsetScope{
			AllComputers: false,
			AllJSSUsers:  false,
		},
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           true,
			SelfServiceDisplayName:      "the-name-we-after",
			InstallButtonText:           "Install",
			ReinstallButtonText:         "",
			SelfServiceDescription:      "",
			ForceUsersToViewDescription: false,
			SelfServiceIcon:             nil,
			FeatureOnMainPage:           true,
			SelfServiceCategories: []jamfpro.PolicySubsetSelfServiceCategory{
				{
					ID:        4133,
					Name:      "Productivity",
					DisplayIn: true,
					FeatureIn: true,
				},
			},
			Notification:        true,
			NotificationType:    "Self Service",
			NotificationSubject: "test",
			NotificationMessage: "",
		},
		Scripts: []jamfpro.PolicySubsetScript{},
		Printers: jamfpro.PolicySubsetPrinters{
			LeaveExistingDefault: false,
		},
		Maintenance: jamfpro.PolicySubsetMaintenance{
			Recon:                    false,
			ResetName:                false,
			InstallAllCachedPackages: false,
			Heal:                     false,
			Prebindings:              false,
			Permissions:              false,
			Byhost:                   false,
			SystemCache:              false,
			UserCache:                false,
			Verify:                   false,
		},
		FilesProcesses: jamfpro.PolicySubsetFilesProcesses{
			DeleteFile:           false,
			UpdateLocateDatabase: false,
			KillProcess:          false,
		},
		UserInteraction: jamfpro.PolicySubsetUserInteraction{
			AllowUsersToDefer:    false,
			AllowDeferralMinutes: 0,
		},
		DiskEncryption: jamfpro.PolicySubsetDiskEncryption{
			Action:                        "",
			DiskEncryptionConfigurationID: 0,
			AuthRestart:                   false,
			RemediateKeyType:              "Individual",
		},
		Reboot: jamfpro.PolicySubsetReboot{
			StartupDisk:                 "Current Startup Disk",
			MinutesUntilReboot:          0,
			StartRebootTimerImmediately: false,
			FileVault2Reboot:            false,
		},
	}

	// Marshal the policy to XML for display
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
	policyXML, err = xml.MarshalIndent(createdPolicy, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling created policy details: %v", err)
	}
	fmt.Println("Created Policy Details:\n", string(policyXML))
}
