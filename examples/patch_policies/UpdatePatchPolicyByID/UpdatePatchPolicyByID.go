package main

import (
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

	// Create the patch policy update request
	policyUpdate := &jamfpro.ResourcePatchPolicyClassicAPI{
		General: jamfpro.ResourcePatchPolicyCreateRequestGeneral{
			Name:               "Google Chrome - 66.0.3359.117",
			Enabled:            true,
			TargetVersion:      "66.0.3359.117",
			ReleaseDate:        "1509048027663",
			IncrementalUpdates: false,
			Reboot:             true,
			MinimumOS:          "10.9",
			KillApps: jamfpro.ResourcePatchPolicyCreateRequestKillApps{
				KillApp: []jamfpro.ResourcePatchPolicyCreateRequestKillApp{
					{
						KillAppName:     "Google Chrome.app",
						KillAppBundleID: "com.google.Chrome",
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
		SoftwareTitleConfigurationID: "1",
	}

	policyID := "1"

	err = client.UpdatePatchPolicyByID(policyID, policyUpdate)
	if err != nil {
		log.Fatalf("Error updating patch policy: %v", err)
	}

	fmt.Printf("Successfully updated patch policy with ID: %s\n", policyID)
}
