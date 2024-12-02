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
			Name:                       "disk-encryption-sdk",
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
			Category: &jamfpro.SharedResourceCategory{
				ID:   -1,
				Name: "No category assigned",
			},
			DateTimeLimitations: &jamfpro.PolicySubsetGeneralDateTimeLimitations{
				ActivationDate:      "",
				ActivationDateEpoch: 0,
				ActivationDateUTC:   "",
				ExpirationDate:      "",
				ExpirationDateEpoch: 0,
				ExpirationDateUTC:   "",
				NoExecuteStart:      "",
				NoExecuteEnd:        "",
			},
			NetworkLimitations: &jamfpro.PolicySubsetGeneralNetworkLimitations{
				MinimumNetworkConnection: "No Minimum",
				AnyIPAddress:             jamfpro.TruePtr(),
				NetworkSegments:          "",
			},
			NetworkRequirements: "Any",
			Site: &jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           jamfpro.FalsePtr(),
			SelfServiceDisplayName:      "",
			InstallButtonText:           "Install",
			ReinstallButtonText:         "",
			SelfServiceDescription:      "",
			ForceUsersToViewDescription: jamfpro.FalsePtr(),
			//SelfServiceIcon:             jamfpro.PolicySelfServiceIcon{ID: -1, Filename: "", URI: ""},
			FeatureOnMainPage: jamfpro.FalsePtr(),
			SelfServiceCategories: []jamfpro.PolicySubsetSelfServiceCategory{
				{
					ID:        -1,
					Name:      "None",
					DisplayIn: jamfpro.FalsePtr(), // or jamfpro.TruePtr(), depending on your requirements
					FeatureIn: jamfpro.FalsePtr(), // or jamfpro.TruePtr(), depending on your requirements
				},
			},
		},
		AccountMaintenance: jamfpro.PolicySubsetAccountMaintenance{
			ManagementAccount: &jamfpro.PolicySubsetAccountMaintenanceManagementAccount{
				Action:                "rotate",
				ManagedPassword:       "",
				ManagedPasswordLength: 0,
			},
			OpenFirmwareEfiPassword: &jamfpro.PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword{
				OfMode:           "none",
				OfPassword:       "",
				OfPasswordSHA256: "",
			},
		},
		Maintenance: jamfpro.PolicySubsetMaintenance{
			Recon:                    jamfpro.FalsePtr(),
			ResetName:                jamfpro.FalsePtr(),
			InstallAllCachedPackages: jamfpro.FalsePtr(),
			Heal:                     jamfpro.FalsePtr(),
			Prebindings:              jamfpro.FalsePtr(),
			Permissions:              jamfpro.FalsePtr(),
			Byhost:                   jamfpro.FalsePtr(),
			SystemCache:              jamfpro.FalsePtr(),
			UserCache:                jamfpro.FalsePtr(),
			Verify:                   jamfpro.FalsePtr(),
		},
		FilesProcesses: jamfpro.PolicySubsetFilesProcesses{
			DeleteFile:           jamfpro.FalsePtr(),
			UpdateLocateDatabase: jamfpro.FalsePtr(),
			SpotlightSearch:      "",
			SearchForProcess:     "",
			KillProcess:          jamfpro.FalsePtr(),
			RunCommand:           "",
		},
		UserInteraction: jamfpro.PolicySubsetUserInteraction{
			MessageStart:          "",
			AllowUsersToDefer:     jamfpro.FalsePtr(),
			AllowDeferralUntilUtc: "",
			AllowDeferralMinutes:  0,
			MessageFinish:         "",
		},
		DiskEncryption: jamfpro.PolicySubsetDiskEncryption{
			Action:                        "apply",
			DiskEncryptionConfigurationID: 1,
			AuthRestart:                   jamfpro.TruePtr(),
			//RemediateKeyType:                       "",
			//RemediateDiskEncryptionConfigurationID: 0,
		},
		Reboot: jamfpro.PolicySubsetReboot{
			Message:                     "This computer will restart in 5 minutes. Please save anything you are working on and log out by choosing Log Out from the bottom of the Apple menu.",
			StartupDisk:                 "Current Startup Disk",
			SpecifyStartup:              "",
			NoUserLoggedIn:              "Do not restart",
			UserLoggedIn:                "Do not restart",
			MinutesUntilReboot:          5,
			StartRebootTimerImmediately: jamfpro.FalsePtr(),
			FileVault2Reboot:            jamfpro.FalsePtr(),
		},
	}

	policyXML, err := xml.MarshalIndent(updatedPolicy, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling policy data: %v", err)
	}
	fmt.Println("Policy Details to be Sent:\n", string(policyXML))

	policyName := "disk-encryption-sdk"

	// Update the policy
	response, err := client.UpdatePolicyByName(policyName, updatedPolicy)
	if err != nil {
		log.Fatalf("Error updating policy: %v", err)
	}

	// Print the ID of the updated policy
	fmt.Printf("Updated Policy ID: %d\n", response.ID)
}
