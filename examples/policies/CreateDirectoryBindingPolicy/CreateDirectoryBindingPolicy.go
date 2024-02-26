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
			Name:                       "jamfpro-sdk-example-DirectoryBinding-policy-config",
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
			Category: jamfpro.PolicyCategory{
				ID:        -1,
				Name:      "No category assigned",
				DisplayIn: false,
				FeatureIn: false,
			},
			DateTimeLimitations: jamfpro.PolicySubsetGeneralDateTimeLimitations{
				// Initialize as needed
			},
			NetworkLimitations: jamfpro.PolicySubsetGeneralNetworkLimitations{
				MinimumNetworkConnection: "No Minimum",
				AnyIPAddress:             true,
				NetworkSegments:          "",
			},
			NetworkRequirements: "Any",
			Site: jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		// Self Service
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           true,
			SelfServiceDisplayName:      "",
			InstallButtonText:           "Install",
			ReinstallButtonText:         "Reinstall",
			SelfServiceDescription:      "",
			ForceUsersToViewDescription: false,
			//SelfServiceIcon:             jamfpro.Icon{ID: -1, Filename: "", URI: ""},
			FeatureOnMainPage: false,
		},
		PackageConfiguration: jamfpro.PolicySubsetPackageConfiguration{
			Packages:          []jamfpro.PolicySubsetPackageConfigurationPackage{}, // Empty packages list
			DistributionPoint: "default",
		},
		AccountMaintenance: jamfpro.PolicySubsetAccountMaintenance{
			// Create directory binding
			DirectoryBindings: []jamfpro.PolicySubsetAccountMaintenanceDirectoryBindings{
				{
					ID:   2,
					Name: "New Binding",
				},
			},
			ManagementAccount: jamfpro.PolicySubsetAccountMaintenanceManagementAccount{
				Action:                "doNotChange",
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
