package main

import (
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

	softwareTitleConfigID := 1 // Replace with actual Software Title Configuration ID

	newPatchPolicy := &jamfpro.ResourcePatchPolicies{
		General: jamfpro.PatchPoliciesSubsetGeneral{
			Name:               "Google Chrome - 66.0.3359.117",
			Enabled:            true,
			TargetVersion:      "66.0.3359.117",
			ReleaseDate:        "1509048027663", // Adjust the format as needed
			IncrementalUpdates: false,
			Reboot:             true,
			MinimumOS:          "10.9",
			KillApps: []jamfpro.PatchPoliciesContainerGeneralKillApps{
				{
					KillApp: jamfpro.PatchPoliciesSubsetGeneralKillApp{
						KillAppName:     "Google Chrome.app",
						KillAppBundleID: "com.google.Chrome",
					},
				},
			},
			DistributionMethod: "selfservice",
			AllowDowngrade:     true,
			PatchUnknown:       true,
		},
		Scope: jamfpro.PatchPoliciesSubsetScope{
			AllComputers: true,
			// Include other fields as necessary
		},
		UserInteraction: jamfpro.PatchPoliciesSubsetUserInteraction{
			InstallButtonText:      "Update",
			SelfServiceDescription: "Latest update for Google Chrome",
			Notifications: jamfpro.PatchPoliciesSubsetUserInteractionNotifications{
				NotificationEnabled: true,
				NotificationType:    "Self Service",
				NotificationSubject: "Google Chrome Update Available",
				NotificationMessage: "An update for Google Chrome is available within Self Service",
				Reminders: jamfpro.PatchPoliciesSubsetUserInteractionNotificationsReminders{
					NotificationRemindersEnabled:  true,
					NotificationReminderFrequency: 1,
				},
			},
			Deadlines: jamfpro.PatchPoliciesSubsetUserInteractionDeadlines{
				DeadlineEnabled: true,
				DeadlinePeriod:  7,
			},
			GracePeriod: jamfpro.PatchPoliciesSubsetUserInteractionGracePeriod{
				GracePeriodDuration:       15,
				NotificationCenterSubject: "Important",
				Message:                   "$APP_NAMES will quit in $DELAY_MINUTES minutes so that $SOFTWARE_TITLE can be updated. Save anything you are working on and quit the app(s)",
			},
		},
		SoftwareTitleConfigurationID: softwareTitleConfigID, // ID applied here
	}

	updatedPatchPolicy, err := client.UpdatePatchPolicy(newPatchPolicy, softwareTitleConfigID)
	if err != nil {
		log.Fatalf("Failed to update patch policy: %v", err)
	}

	fmt.Printf("Updated Patch Policy: %+v\n", updatedPatchPolicy)
}
