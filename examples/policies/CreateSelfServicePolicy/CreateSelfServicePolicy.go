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
			Name:                       "sample-self-service-policy",
			Enabled:                    jamfpro.TruePtr(),
			TriggerCheckin:             jamfpro.FalsePtr(),
			TriggerEnrollmentComplete:  jamfpro.FalsePtr(),
			TriggerLogin:               jamfpro.FalsePtr(),
			TriggerLogout:              jamfpro.FalsePtr(),
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
			OverrideDefaultSettings: &jamfpro.PolicySubsetGeneralOverrideSettings{
				TargetDrive:       "/",
				DistributionPoint: "default",
				ForceAfpSmb:       jamfpro.FalsePtr(),
				SUS:               "default",
				NetbootServer:     "",
			},
			NetworkRequirements: "Any",
			Site: &jamfpro.SharedResourceSite{
				ID:   -1,
				Name: "NONE",
			},
		},
		Scope: jamfpro.PolicySubsetScope{
			AllComputers: jamfpro.FalsePtr(),
			AllJSSUsers:  jamfpro.FalsePtr(),
			Limitations:  &jamfpro.PolicySubsetScopeLimitations{},
			Exclusions:   &jamfpro.PolicySubsetScopeExclusions{},
		},
		SelfService: jamfpro.PolicySubsetSelfService{
			UseForSelfService:           jamfpro.TruePtr(),
			SelfServiceDisplayName:      "the-name-we-after",
			InstallButtonText:           "Install",
			ReinstallButtonText:         "",
			SelfServiceDescription:      "",
			ForceUsersToViewDescription: jamfpro.FalsePtr(),
			SelfServiceIcon:             nil,
			FeatureOnMainPage:           jamfpro.TruePtr(),
			SelfServiceCategories: []jamfpro.PolicySubsetSelfServiceCategory{
				{
					ID:        4133,
					Name:      "Productivity",
					DisplayIn: jamfpro.TruePtr(),
					FeatureIn: jamfpro.TruePtr(),
				},
			},
			Notification:        jamfpro.TruePtr(),
			NotificationType:    "Self Service",
			NotificationSubject: "test",
			NotificationMessage: "",
		},
		PackageConfiguration: jamfpro.PolicySubsetPackageConfiguration{
			Packages:          []jamfpro.PolicySubsetPackageConfigurationPackage{},
			DistributionPoint: "default",
		},
		Scripts: []jamfpro.PolicySubsetScript{
			{
				ID:       "14009",
				Name:     "Sample Script",
				Priority: "Before",
			},
		},
		Printers: jamfpro.PolicySubsetPrinters{
			LeaveExistingDefault: jamfpro.FalsePtr(),
			Printer:              []jamfpro.PolicySubsetPrinter{},
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
