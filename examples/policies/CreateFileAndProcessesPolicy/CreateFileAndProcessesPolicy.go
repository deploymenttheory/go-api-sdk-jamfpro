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
			Name:                       "jamfpro-sdk-example-FileAndProcesses-policy-config",
			Enabled:                    false,
			TriggerOther:               "EVENT",
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
			Category: &jamfpro.SharedResourceCategory{
				ID:   -1,
				Name: "No category assigned",
			},
			DateTimeLimitations: &jamfpro.PolicySubsetGeneralDateTimeLimitations{
				// Initialize as needed
			},
			NetworkLimitations: &jamfpro.PolicySubsetGeneralNetworkLimitations{
				MinimumNetworkConnection: "No Minimum",
				AnyIPAddress:             true,
				NetworkSegments:          "",
			},
			OverrideDefaultSettings: &jamfpro.PolicySubsetGeneralOverrideSettings{
				TargetDrive:       "/",
				DistributionPoint: "default",
				ForceAfpSmb:       false,
				SUS:               "default",
			},
			NetworkRequirements: "Any",
			Site: &jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "None",
			},
		},
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           true,
			SelfServiceDisplayName:      "",
			InstallButtonText:           "Install",
			ReinstallButtonText:         "",
			SelfServiceDescription:      "",
			ForceUsersToViewDescription: false,
			//SelfServiceIcon:             jamfpro.PolicySelfServiceIcon{ID: -1, Filename: "", URI: ""},
			FeatureOnMainPage: false,
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
		// File and processes policy
		FilesProcesses: jamfpro.PolicySubsetFilesProcesses{
			SearchByPath:         "thing",
			DeleteFile:           true,
			LocateFile:           "thing",
			UpdateLocateDatabase: true,
			SpotlightSearch:      "thing",
			SearchForProcess:     "thing",
			KillProcess:          true,
			RunCommand:           "thing",
		},
		UserInteraction: jamfpro.PolicySubsetUserInteraction{
			MessageStart:          "thing",
			AllowUserToDefer:      true,
			AllowDeferralUntilUtc: "",
			AllowDeferralMinutes:  1440,
			MessageFinish:         "thing",
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
