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

	// Create a new jamfpro client instance
	client, err := jamfpro.BuildClient(config)
	if err != nil {
		log.Fatalf("Failed to create Jamf Pro client: %v", err)
	}
	// Define a new policy with all required fields
	newPolicy := &jamfpro.ResourcePolicy{
		General: jamfpro.PolicySubsetGeneral{
			Name:                       "jamfpro-sdk-scope-policy-config",
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
		},
		Scope: jamfpro.PolicySubsetScope{
			AllComputers: false,
			AllJSSUsers:  false,
			Computers: []jamfpro.PolicyDataSubsetComputer{
				{
					ID:   18,
					Name: "APIGWFYFGH",
				},
				{
					ID:   16,
					Name: "APIGWFYKTF",
				},
			}, // Empty as per XML
			ComputerGroups: []jamfpro.PolicyDataSubsetComputerGroup{
				{
					ID:   40,
					Name: "test-smart-computer-group-01",
				},
				{
					ID:   41,
					Name: "test-static-computer-group-01",
				},
				// Additional computer groups can be added here if needed
			},
			JSSUsers: []jamfpro.PolicyDataSubsetJSSUser{
				// {
				// 	ID:   12,
				// 	Name: "Barry White",
				// },
				// {
				// 	ID:   39,
				// 	Name: "Michael Jackson",
				// },
				// Additional jss users can be added here if needed
			},
			JSSUserGroups: []jamfpro.PolicyDataSubsetJSSUserGroup{
				{
					ID:   3,
					Name: "test-static-user-group-01",
				},
				{
					ID:   4,
					Name: "test-smart-user-group-01",
				},
			}, // Empty as per XML
			Buildings: []jamfpro.PolicyDataSubsetBuilding{
				{
					ID:   1320,
					Name: "Apple Park 2",
				},
				{
					ID:   1321,
					Name: "Apple Infinite Loop",
				},
				// Additional buildings can be added here if needed
			},
			Departments: []jamfpro.PolicyDataSubsetDepartment{
				{
					ID:   23517,
					Name: "Marketing",
				},
				{
					ID:   23518,
					Name: "Modern workpalce",
				},
				// Additional departments can be added here if needed
			},
			LimitToUsers: jamfpro.PolicyLimitToUsers{
				UserGroups: []string{}, // Empty as per XML
				// Additional user groups can be added here if needed
			},
			Limitations: jamfpro.PolicySubsetScopeLimitations{
				// supports jamf pro and directory services user objects
				Users: []jamfpro.PolicyDataSubsetUser{
					{
						ID:   12,
						Name: "Barry White",
					},
					{
						ID:   39,
						Name: "Michael Jackson",
					},
					{
						ID:   4,
						Name: "Jane Smith",
					},
					{
						ID:   3,
						Name: "John Doe",
					},
					// Additional users can be added here if needed
				},
				UserGroups: []jamfpro.PolicyDataSubsetUserGroup{
					// Directory Service User Groups
					// {
					// 	ID:   3,
					// 	Name: "test-static-user-group-01",
					// },
					// {
					// 	ID:   4,
					// 	Name: "test-smart-user-group-01",
					// },
					// Additional computer groups can be added here if needed
				},
				NetworkSegments: []jamfpro.PolicyDataSubsetNetworkSegment{
					{
						ID:   2,
						Name: "NY Office",
						UID:  "",
					},
					{
						ID:   3,
						Name: "London Office",
						UID:  "",
					},
					// Additional network segment can be added here if needed
				},
				IBeacons: []jamfpro.PolicyDataSubsetIBeacon{
					{
						ID:   3,
						Name: "test-ibeacon-01",
					},
					{
						ID:   4,
						Name: "test-ibeacon-02",
					},
					// Additional iBeacons can be added here if needed
				},
			},
			Exclusions: jamfpro.PolicySubsetScopeExclusions{
				Computers: []jamfpro.PolicyDataSubsetComputer{
					{
						ID:   18,
						Name: "APIGWFYFGH",
					},
					{
						ID:   16,
						Name: "APIGWFYKTF",
					},
				},
				ComputerGroups: []jamfpro.PolicyDataSubsetComputerGroup{
					{
						ID:   40,
						Name: "test-smart-computer-group-01",
					},
					{
						ID:   41,
						Name: "test-static-computer-group-01",
					},
				},
				Users: []jamfpro.PolicyDataSubsetUser{
					{
						ID:   12,
						Name: "Barry White",
					},
					{
						ID:   39,
						Name: "Michael Jackson",
					},
				},
				UserGroups: []jamfpro.PolicyDataSubsetUserGroup{
					// appears to not be working
				},
				Buildings: []jamfpro.PolicyDataSubsetBuilding{
					{
						ID:   1320,
						Name: "Apple Park 2",
					},
					{
						ID:   1321,
						Name: "Apple Infinite Loop",
					},
				},
				Departments: []jamfpro.PolicyDataSubsetDepartment{
					{
						ID:   23517,
						Name: "Marketing",
					},
					{
						ID:   23518,
						Name: "Modern workpalce",
					},
				},
				NetworkSegments: []jamfpro.PolicyDataSubsetNetworkSegment{
					{
						ID:   2,
						Name: "NY Office",
						UID:  "",
					},
					{
						ID:   3,
						Name: "London Office",
						UID:  "",
					},
				},
				JSSUsers: []jamfpro.PolicyDataSubsetJSSUser{
					{
						ID:   4,
						Name: "Jane Smith",
					},
					{
						ID:   3,
						Name: "John Doe",
					},
				}, // Empty as per XML
				JSSUserGroups: []jamfpro.PolicyDataSubsetJSSUserGroup{
					{
						ID:   3,
						Name: "test-static-user-group-01",
					},
					{
						ID:   4,
						Name: "test-smart-user-group-01",
					},
				}, // Empty as per XML
				IBeacons: []jamfpro.PolicyDataSubsetIBeacon{
					{
						ID:   3,
						Name: "test-ibeacon-01",
					},
					{
						ID:   4,
						Name: "test-ibeacon-02",
					},
				},
			},
		},
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           false,
			SelfServiceDisplayName:      "",
			InstallButtonText:           "Install",
			ReinstallButtonText:         "",
			SelfServiceDescription:      "",
			ForceUsersToViewDescription: false,
			//SelfServiceIcon:             jamfpro.Icon{ID: -1, Filename: "", URI: ""},
			FeatureOnMainPage: false,
		},
		AccountMaintenance: jamfpro.PolicySubsetAccountMaintenance{
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
		// User interation policy settings
		UserInteraction: jamfpro.PolicySubsetUserInteraction{
			MessageStart:          "",
			AllowUserToDefer:      true,
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
