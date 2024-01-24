package main

import (
	"encoding/xml"
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/http_client" // Import http_client for logging
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
	logLevel := http_client.LogLevelDebug // LogLevelNone // LogLevelWarning // LogLevelInfo  // LogLevelDebug

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

	// Define a new policy with all required fields
	updatedPolicy := &jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			Name:                       "disk-encryption-sdk",
			Enabled:                    false,
			Trigger:                    "EVENT",
			TriggerCheckin:             false,
			TriggerEnrollmentComplete:  false,
			TriggerLogin:               false,
			TriggerLogout:              false,
			TriggerNetworkStateChanged: false,
			TriggerStartup:             false,
			Frequency:                  "Once per computer",
			RetryEvent:                 "none",
			RetryAttempts:              -1,
			NotifyOnEachFailedRetry:    false,
			LocationUserOnly:           false,
			TargetDrive:                "/",
			Offline:                    false,
		},
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           false,
			SelfServiceDisplayName:      "",
			InstallButtonText:           "Install",
			ReinstallButtonText:         "",
			SelfServiceDescription:      "",
			ForceUsersToViewDescription: false,
			//SelfServiceIcon:             jamfpro.PolicySelfServiceIcon{ID: -1, Filename: "", URI: ""},
			FeatureOnMainPage: false,
			SelfServiceCategories: []jamfpro.PolicySubsetSelfServiceCategory{
				{
					Category: jamfpro.PolicyCategory{
						//ID:        "-1",
						//Name:      "None",
						DisplayIn: false, // or true, depending on your requirements
						FeatureIn: false, // or true, depending on your requirements
					},
				},
			},
		},
		AccountMaintenance: jamfpro.PolicySubsetAccountMaintenance{
			ManagementAccount: jamfpro.PolicySubsetAccountMaintenanceManagementAccount{
				Action:                "rotate",
				ManagedPassword:       "",
				ManagedPasswordLength: 0,
			},
			OpenFirmwareEfiPassword: jamfpro.PolicySubsetAccountMaintenanceOpenFirmwareEfiPassword{
				OfMode:           "none",
				OfPassword:       "",
				OfPasswordSHA256: "",
			},
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
			SpotlightSearch:      "",
			SearchForProcess:     "",
			KillProcess:          false,
			RunCommand:           "",
		},
		UserInteraction: jamfpro.PolicySubsetUserInteraction{
			MessageStart:          "",
			AllowUserToDefer:      false,
			AllowDeferralUntilUtc: "",
			AllowDeferralMinutes:  0,
			MessageFinish:         "",
		},
		DiskEncryption: jamfpro.PolicySubsetDiskEncryption{
			Action:                        "apply",
			DiskEncryptionConfigurationID: 1,
			AuthRestart:                   true,
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
			StartRebootTimerImmediately: false,
			FileVault2Reboot:            false,
		},
	}

	policyXML, err := xml.MarshalIndent(updatedPolicy, "", "    ")
	if err != nil {
		log.Fatalf("Error marshaling policy data: %v", err)
	}
	fmt.Println("Policy Details to be Sent:\n", string(policyXML))

	policyName := "policy-creation-test"

	// Update the policy
	response, err := client.UpdatePolicyByName(policyName, updatedPolicy)
	if err != nil {
		log.Fatalf("Error updating policy: %v", err)
	}

	// Print the ID of the updated policy
	fmt.Printf("Updated Policy ID: %d\n", response.ID)

}
