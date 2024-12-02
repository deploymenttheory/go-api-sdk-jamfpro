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
	updatedPolicy := &jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			Name:                       "example-script-policy-config",
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
				ID: 4135,
			},
			// Category: &jamfpro.SharedResourceCategory{
			// 	ID:   -1,
			// 	Name: "No category assigned",
			// },
			NetworkLimitations: &jamfpro.PolicySubsetGeneralNetworkLimitations{
				MinimumNetworkConnection: "No Minimum",
				AnyIPAddress:             false,
				NetworkSegments:          "",
			},
			NetworkRequirements: "Any",
			Site: &jamfpro.SharedResourceSite{
				ID: 3855,
			},
			// Site: &jamfpro.SharedResourceSite{
			// 	ID:   -1,
			// 	Name: "NONE",
			// },
		},
		Scope: jamfpro.PolicySubsetScope{
			AllComputers: false,
			AllJSSUsers:  false,
		},
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           true,
			InstallButtonText:           "Install",
			ReinstallButtonText:         "REINSTALL",
			ForceUsersToViewDescription: false,
			FeatureOnMainPage:           false,
			Notification:                false,
		},
		Scripts: []jamfpro.PolicySubsetScript{
			{
				ID:          "14011",
				Priority:    "After",
				Parameter4:  "param_value_4",
				Parameter5:  "param_value_5",
				Parameter6:  "param_value_6",
				Parameter7:  "param_value_7",
				Parameter8:  "param_value_8",
				Parameter9:  "param_value_9",
				Parameter10: "param_value_10",
				Parameter11: "param_value_11",
			},
		},
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

	policyXML, err := xml.MarshalIndent(updatedPolicy, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling policy data: %v", err)
	}
	fmt.Println("Policy Details to be Sent:\n", string(policyXML))

	policyID := "1"

	// Update the policy
	response, err := client.UpdatePolicyByID(policyID, updatedPolicy)
	if err != nil {
		log.Fatalf("Error updating policy: %v", err)
	}

	// Print the ID of the updated policy
	fmt.Printf("Updated Policy ID: %d\n", response.ID)
}
