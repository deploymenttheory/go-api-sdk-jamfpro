package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	configFilePath := "/Users/dafyddwatkins/localtesting/jamfpro/clientconfig.json"

	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create new policy matching the XML structure
	newPolicy := &jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			Name:    "minimum",
			Enabled: jamfpro.TruePtr(),
			//Trigger:                    "EVENT",
			TriggerCheckin:             jamfpro.FalsePtr(),
			TriggerEnrollmentComplete:  jamfpro.FalsePtr(),
			TriggerLogin:               jamfpro.FalsePtr(),
			TriggerNetworkStateChanged: jamfpro.FalsePtr(),
			TriggerStartup:             jamfpro.FalsePtr(),
			TriggerOther:               "",
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
			Site: &jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "NONE",
			},
		},
		Scope: jamfpro.PolicySubsetScope{
			AllComputers: jamfpro.FalsePtr(),
		},
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           jamfpro.FalsePtr(),
			InstallButtonText:           "Install",
			ReinstallButtonText:         "Reinstall",
			ForceUsersToViewDescription: jamfpro.FalsePtr(),
			FeatureOnMainPage:           jamfpro.FalsePtr(),
			Notification:                jamfpro.FalsePtr(),
			NotificationType:            "Self Service",
			NotificationSubject:         "minimum",
		},
		PackageConfiguration: jamfpro.PolicySubsetPackageConfiguration{
			DistributionPoint: "default",
		},
		Printers: jamfpro.PolicySubsetPrinters{
			LeaveExistingDefault: jamfpro.FalsePtr(),
		},
		AccountMaintenance: jamfpro.PolicySubsetAccountMaintenance{
			ManagementAccount: &jamfpro.PolicySubsetAccountMaintenanceManagementAccount{
				Action:                "doNotChange",
				ManagedPasswordLength: 0,
			},
			OpenFirmwareEfiPassword: &jamfpro.PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword{
				OfMode: "none",
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
			KillProcess:          jamfpro.FalsePtr(),
		},
		UserInteraction: jamfpro.PolicySubsetUserInteraction{
			AllowUsersToDefer:    jamfpro.FalsePtr(),
			AllowDeferralMinutes: 0,
		},
		DiskEncryption: jamfpro.PolicySubsetDiskEncryption{
			Action:                        "none",
			DiskEncryptionConfigurationID: 0,
		},
		Reboot: jamfpro.PolicySubsetReboot{
			Message:                     "This computer will restart in 5 minutes. Please save anything you are working on and log out by choosing Log Out from the bottom of the Apple menu.",
			StartupDisk:                 "Current Startup Disk",
			NoUserLoggedIn:              "Do not restart",
			UserLoggedIn:                "Do not restart",
			MinutesUntilReboot:          5,
			StartRebootTimerImmediately: jamfpro.FalsePtr(),
			FileVault2Reboot:            jamfpro.FalsePtr(),
		},
	}

	policyXML, err := xml.MarshalIndent(newPolicy, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling policy data: %v", err)
	}
	fmt.Println("Policy Details to be Sent:\n", string(policyXML))

	createdPolicy, err := client.CreatePolicy(newPolicy)
	if err != nil {
		log.Fatalf("Error creating policy: %v", err)
	}

	policyXML, err = xml.MarshalIndent(createdPolicy, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling policy details data: %v", err)
	}
	fmt.Println("Created Policy Details:\n", string(policyXML))
}
