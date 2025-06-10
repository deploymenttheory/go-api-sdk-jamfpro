package main

import (
	"fmt"
	"log"

	"github.com/deploymenttheory/go-api-sdk-jamfpro/sdk/jamfpro"
)

func main() {
	// Define the path to the JSON configuration file
	configFilePath := "./clientconfig.json"

	// Initialize the Jamf Pro client with the HTTP client configuration
	client, err := jamfpro.BuildClientWithConfigFile(configFilePath)
	if err != nil {
		log.Fatalf("Failed to initialize Jamf Pro client: %v", err)
	}

	// Create the patch policy request
	policyRequest := &jamfpro.ResourcePatchPolicyClassicAPI{
		General: jamfpro.ResourcePatchPolicyCreateRequestGeneral{
			Name:               "Google Chrome - 66.0.3359.117",
			Enabled:            true,
			TargetVersion:      "12.2.1",
			ReleaseDate:        "2014-02-13T00:31:29Z",
			IncrementalUpdates: false,
			Reboot:             true,
			MinimumOS:          "10.9",
			KillApps: jamfpro.ResourcePatchPolicyCreateRequestKillApps{
				KillApp: []jamfpro.ResourcePatchPolicyCreateRequestKillApp{
					{
						KillAppName:     "Adobe After Effects CC",
						KillAppBundleID: "",
					},
				},
			},
			DistributionMethod: "selfservice",
			AllowDowngrade:     true,
			PatchUnknown:       true,
		},
		Scope: jamfpro.ResourcePatchPolicyCreateRequestScope{
			AllComputers: true,
		},
		UserInteraction: jamfpro.ResourcePatchPolicyCreateRequestUserInteraction{
			InstallButtonText:      "Update",
			SelfServiceDescription: "Latest update for Google Chrome",
			Notifications: jamfpro.ResourcePatchPolicyCreateRequestNotifications{
				Enabled: true,
				Type:    "Self Service",
				Subject: "Google Chrome Update Available",
				Message: "An update for Google Chrome is available within Self Service",
				Reminders: jamfpro.ResourcePatchPolicyCreateRequestReminders{
					Enabled:   true,
					Frequency: 1,
				},
			},
			Deadlines: jamfpro.ResourcePatchPolicyCreateRequestDeadlines{
				Enabled: true,
				Period:  7,
			},
			GracePeriod: jamfpro.ResourcePatchPolicyCreateRequestGracePeriod{
				Duration:            15,
				NotificationSubject: "Important",
				Message:             "$APP_NAMES will quit in $DELAY_MINUTES minutes so that $SOFTWARE_TITLE can be updated. Save anything you are working on and quit the app(s)",
			},
		},
		SoftwareTitleConfigurationID: "12",
	}

	patchSoftwareTitleConfigurationID := "12"

	// Create the patch policy
	err = client.CreatePatchPolicy(patchSoftwareTitleConfigurationID, policyRequest)
	if err != nil {
		log.Fatalf("Error creating patch policy: %v", err)
	}

	fmt.Println("Successfully created patch policy for Google Chrome")
}
