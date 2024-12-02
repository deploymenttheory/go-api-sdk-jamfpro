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
			Name:                       "jamfpro-sdk-userinteration-policy-config",
			Enabled:                    jamfpro.FalsePtr(),
			TriggerOther:               "EVENT",
			TriggerCheckin:             jamfpro.FalsePtr(),
			TriggerEnrollmentComplete:  jamfpro.FalsePtr(),
			TriggerLogin:               jamfpro.FalsePtr(),
			TriggerLogout:              jamfpro.FalsePtr(),
			TriggerNetworkStateChanged: jamfpro.FalsePtr(),
			TriggerStartup:             jamfpro.FalsePtr(),
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
		},
		AccountMaintenance: jamfpro.PolicySubsetAccountMaintenance{
			ManagementAccount: &jamfpro.PolicySubsetAccountMaintenanceManagementAccount{
				Action:                "doNotChange",
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
			AllowUsersToDefer:     jamfpro.TruePtr(),
			AllowDeferralUntilUtc: "",
			AllowDeferralMinutes:  0,
			MessageFinish:         "",
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
